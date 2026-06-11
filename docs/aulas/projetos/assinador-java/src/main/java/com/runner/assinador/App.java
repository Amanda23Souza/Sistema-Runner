package com.runner.assinador;

import io.javalin.Javalin;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.Map;
import java.util.concurrent.Executors;
import java.util.concurrent.ScheduledExecutorService;
import java.util.concurrent.ScheduledFuture;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.atomic.AtomicLong;
import java.util.concurrent.atomic.AtomicReference;

/**
 * Ponto de entrada do servidor assinador HTTP.
 *
 * <p>Suporta:
 * <ul>
 *   <li>GET  /health   — health check (status UP quando pronto para receber requisições)</li>
 *   <li>POST /sign     — assina um conteúdo</li>
 *   <li>POST /validate — valida uma assinatura</li>
 *   <li>POST /shutdown — encerra o servidor de forma controlada</li>
 * </ul>
 *
 * <p>Auto-shutdown: encerra automaticamente após {@code --inactivity-timeout} segundos
 * sem receber requisições (padrão: 300s). O timer é reiniciado a cada requisição.
 */
public class App {

    private static final Logger LOG = LoggerFactory.getLogger(App.class);
    private static final int DEFAULT_PORT = 8080;
    private static final int DEFAULT_INACTIVITY_TIMEOUT_SECONDS = 300; // 5 minutos

    public static void main(String[] args) {
        int port = DEFAULT_PORT;
        int inactivityTimeout = DEFAULT_INACTIVITY_TIMEOUT_SECONDS;

        // Parse de argumentos: <port> [--inactivity-timeout <segundos>]
        for (int i = 0; i < args.length; i++) {
            if (i == 0 && !args[0].startsWith("--")) {
                try {
                    port = Integer.parseInt(args[0]);
                } catch (NumberFormatException e) {
                    System.err.println("Porta inválida: '" + args[0] + "'. Usando porta padrão " + DEFAULT_PORT + ".");
                    System.err.println("Como resolver: passe um número de porta válido (ex: 8080).");
                }
            } else if ("--inactivity-timeout".equals(args[i]) && i + 1 < args.length) {
                try {
                    inactivityTimeout = Integer.parseInt(args[++i]);
                } catch (NumberFormatException e) {
                    System.err.println("Timeout de inatividade inválido: '" + args[i] + "'. Usando padrão " + DEFAULT_INACTIVITY_TIMEOUT_SECONDS + "s.");
                }
            }
        }

        startServer(port, inactivityTimeout);
    }

    static Javalin startServer(int port, int inactivityTimeoutSeconds) {
        SignatureService signatureService = new FakeSignatureService();
        SignatureController controller = new SignatureController(signatureService);

        // Rastreia o timestamp da última requisição para auto-shutdown
        AtomicLong lastActivity = new AtomicLong(System.currentTimeMillis());
        AtomicReference<Javalin> appRef = new AtomicReference<>();

        ScheduledExecutorService scheduler = Executors.newSingleThreadScheduledExecutor(r -> {
            Thread t = new Thread(r, "auto-shutdown-watchdog");
            t.setDaemon(true);
            return t;
        });

        Javalin app = Javalin.create(config -> {
            config.showJavalinBanner = false;
        });

        appRef.set(app);

        // Middleware: atualiza timestamp de última atividade a cada requisição
        app.before(ctx -> lastActivity.set(System.currentTimeMillis()));

        // Health check — distingue "processo subiu" de "pronto para receber requisições"
        app.get("/health", ctx -> ctx.status(200).json(Map.of(
                "status", "UP",
                "port", port,
                "inactivityTimeoutSeconds", inactivityTimeoutSeconds
        )));

        // Endpoints de assinatura
        app.post("/sign", controller::sign);
        app.post("/validate", controller::validate);

        // Endpoint de shutdown controlado
        app.post("/shutdown", ctx -> {
            LOG.info("Shutdown solicitado via endpoint /shutdown");
            ctx.status(200).json(Map.of("message", "Servidor encerrando..."));
            // Encerra em background para que a resposta seja enviada
            new Thread(() -> {
                try {
                    Thread.sleep(200);
                } catch (InterruptedException ignored) {
                    Thread.currentThread().interrupt();
                }
                appRef.get().stop();
                scheduler.shutdownNow();
            }, "shutdown-thread").start();
        });

        app.start(port);
        LOG.info("Assinador iniciado na porta {} (timeout de inatividade: {}s)", port, inactivityTimeoutSeconds);

        // Agenda verificação periódica de inatividade
        final int timeoutSeconds = inactivityTimeoutSeconds;
        ScheduledFuture<?> watchdogTask = scheduler.scheduleAtFixedRate(() -> {
            long idleMs = System.currentTimeMillis() - lastActivity.get();
            if (idleMs > timeoutSeconds * 1000L) {
                LOG.info("Auto-shutdown por inatividade: {}s sem requisições. Encerrando servidor.", idleMs / 1000);
                appRef.get().stop();
                scheduler.shutdownNow();
            }
        }, timeoutSeconds, 1, TimeUnit.SECONDS);

        // Garante limpeza ao encerrar via sinal do SO
        Runtime.getRuntime().addShutdownHook(new Thread(() -> {
            watchdogTask.cancel(true);
            scheduler.shutdownNow();
            LOG.info("Servidor assinador encerrado.");
        }, "shutdown-hook"));

        return app;
    }
}
