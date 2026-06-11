package com.runner.assinador;

import io.javalin.Javalin;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.Timeout;

import java.net.ServerSocket;
import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.fail;

/**
 * Verifica o comportamento do auto-shutdown por inatividade (critério E2.4).
 *
 * <p>Prova que:
 * <ol>
 *   <li>O timer reinicia a cada requisição recebida.</li>
 *   <li>O servidor encerra automaticamente após o timeout de inatividade.</li>
 * </ol>
 */
class InactivityTimerTest {

    private static final HttpClient CLIENT = HttpClient.newHttpClient();

    /**
     * Cenário: timer reinicia a cada requisição.
     *
     * <p>Sequência:
     * <ul>
     *   <li>t=0s  — servidor inicia com timeout=2s</li>
     *   <li>t=1s  — requisição feita (timer reseta para t=1s)</li>
     *   <li>t=2s  — servidor ainda ATIVO (idle=1s &lt; 2s desde último reset)</li>
     *   <li>t=5s  — servidor INATIVO (idle=4s &gt; 2s; watchdog com período=1s já disparou)</li>
     * </ul>
     */
    @Test
    @Timeout(15)
    void timerResetsOnRequestAndServerShutsDownAfterInactivity() throws Exception {
        int port = findFreePort();
        Javalin server = App.startServer(port, 2);

        try {
            assertServerUp(port);

            // Faz uma requisição em t=1s — reseta o timer de inatividade
            Thread.sleep(1000);
            sendHealthRequest(port);

            // Em t=2s, servidor ainda deve estar ativo (idle=1s < 2s desde o reset)
            Thread.sleep(1000);
            assertServerUp(port);

            // Aguarda até t=5s sem mais requisições — watchdog (período=1s) dispara quando idle>2s
            Thread.sleep(3000);
            assertServerDown(port);
        } finally {
            try {
                server.stop();
            } catch (Exception ignored) {
                // Servidor pode já ter encerrado via auto-shutdown
            }
        }
    }

    /**
     * Cenário: sem requisições, servidor encerra após o timeout.
     */
    @Test
    @Timeout(10)
    void serverShutsDownAfterInactivityWithNoRequests() throws Exception {
        int port = findFreePort();
        Javalin server = App.startServer(port, 2);

        try {
            assertServerUp(port);

            // Aguarda sem fazer requisições — watchdog dispara quando idle > 2s
            // Primeiro disparo em t=2s, mas idle == 2s ainda não é > 2s.
            // Segundo disparo em t=3s, idle=3s > 2s → shutdown.
            Thread.sleep(5000);
            assertServerDown(port);
        } finally {
            try {
                server.stop();
            } catch (Exception ignored) {
                // Servidor pode já ter encerrado via auto-shutdown
            }
        }
    }

    private void sendHealthRequest(int port) throws Exception {
        HttpRequest req = HttpRequest.newBuilder()
                .uri(URI.create("http://localhost:" + port + "/health"))
                .GET()
                .build();
        CLIENT.send(req, HttpResponse.BodyHandlers.ofString());
    }

    private void assertServerUp(int port) throws Exception {
        HttpRequest req = HttpRequest.newBuilder()
                .uri(URI.create("http://localhost:" + port + "/health"))
                .GET()
                .build();
        HttpResponse<String> resp = CLIENT.send(req, HttpResponse.BodyHandlers.ofString());
        assertEquals(200, resp.statusCode(), "Servidor deveria estar ativo na porta " + port);
    }

    private void assertServerDown(int port) {
        try {
            HttpRequest req = HttpRequest.newBuilder()
                    .uri(URI.create("http://localhost:" + port + "/health"))
                    .GET()
                    .build();
            CLIENT.send(req, HttpResponse.BodyHandlers.ofString());
            fail("Esperava servidor encerrado na porta " + port + ", mas recebeu resposta");
        } catch (Exception e) {
            // Esperado: conexão recusada — servidor encerrou por inatividade
        }
    }

    private int findFreePort() throws Exception {
        try (ServerSocket s = new ServerSocket(0)) {
            return s.getLocalPort();
        }
    }
}
