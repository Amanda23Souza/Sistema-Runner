package command

import (
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"time"
)

// StartCmd implementa o comando "start" — inicia o servidor assinador em background.
type StartCmd struct {
	out     io.Writer
	errOut  io.Writer
	port    int
	jarPath string
	verbose bool
}

// NewStartCmd cria uma nova instância do comando start.
func NewStartCmd() *StartCmd {
	return &StartCmd{
		out:    os.Stdout,
		errOut: os.Stderr,
	}
}

// Help retorna a descrição de uso do comando start.
func (c *StartCmd) Help() string {
	return `Usage: assinatura start [OPTIONS]

Inicia o servidor assinador (assinador.jar) em background.

Idempotente: se já houver uma instância saudável na porta indicada,
o comando reutiliza a instância existente sem iniciar outra.

Options:
  --port PORT       Porta em que o servidor irá escutar (padrão: 8080)
  --jar  PATH       Caminho para o assinador.jar (padrão: busca automática)
  --verbose         Habilita saída detalhada
  --help            Exibe esta mensagem de ajuda

Examples:
  assinatura start
  assinatura start --port 9090
  assinatura start --jar /caminho/assinador.jar --port 8080`
}

// Run executa o comando start.
func (c *StartCmd) Run(args []string) error {
	fs := flag.NewFlagSet("start", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.IntVar(&c.port, "port", 8080, "Porta do servidor")
	fs.StringVar(&c.jarPath, "jar", "", "Caminho para o assinador.jar")
	fs.BoolVar(&c.verbose, "verbose", false, "Habilita saída detalhada")

	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			fmt.Fprintln(c.out, c.Help())
			return nil
		}
	}

	if err := fs.Parse(args); err != nil {
		fmt.Fprintf(c.errOut, "Erro do usuário: parâmetro inválido.\n")
		return &UserError{msg: "sintaxe de comando inválida"}
	}

	// Idempotência: verifica se já há instância saudável na porta
	healthURL := fmt.Sprintf("http://localhost:%d/health", c.port)
	if resp, err := httpGet(healthURL); err == nil {
		if status, ok := resp["status"].(string); ok && status == "UP" {
			fmt.Fprintf(c.out, "Servidor assinador já está em execução e saudável na porta %d.\n", c.port)
			slog.Info("instância existente reutilizada", "porta", c.port)
			return nil
		}
	}

	// Localiza o JAR
	jarPath, err := c.findJar()
	if err != nil {
		fmt.Fprintf(c.errOut, "Erro do sistema: %v\n", err)
		fmt.Fprintf(c.errOut, "Como resolver: use --jar /caminho/para/assinador.jar\n")
		slog.Error("JAR não encontrado", "erro", err)
		return fmt.Errorf("assinador.jar não encontrado: %w", err)
	}

	if c.verbose {
		slog.Info("iniciando servidor", "jar", jarPath, "porta", c.port)
	}

	// Verifica se JVM está disponível
	if _, err := exec.LookPath("java"); err != nil {
		fmt.Fprintf(c.errOut, "Erro do sistema: 'java' não encontrado no PATH.\n")
		fmt.Fprintf(c.errOut, "Como resolver: instale o JDK 21+ e certifique-se de que 'java' está no PATH.\n")
		slog.Error("JVM não encontrada no PATH")
		return fmt.Errorf("JVM (java) não encontrada no PATH")
	}

	// Verifica versão do Java (mínimo 21)
	if err := checkJavaVersion(c.errOut); err != nil {
		fmt.Fprintf(c.errOut, "Aviso: %v\n", err)
	}

	// Verifica se a porta está ocupada por outro processo (não nosso JAR)
	if isPortOccupied(c.port) {
		fmt.Fprintf(c.errOut, "Erro do sistema: a porta %d está ocupada por outro processo.\n", c.port)
		fmt.Fprintf(c.errOut, "Como resolver: escolha outra porta com --port <numero> ou encerre o processo que está usando a porta %d.\n", c.port)
		slog.Error("porta ocupada por outro processo", "porta", c.port)
		return fmt.Errorf("porta %d ocupada por outro processo", c.port)
	}

	// Inicia o servidor em background
	cmd := exec.Command("java", "-jar", jarPath, strconv.Itoa(c.port))
	cmd.Stdout = nil // Desacopla stdout do processo filho
	cmd.Stderr = nil

	if err := cmd.Start(); err != nil {
		fmt.Fprintf(c.errOut, "Erro do sistema: falha ao iniciar o servidor assinador.\n")
		fmt.Fprintf(c.errOut, "Causa: %v\n", err)
		slog.Error("falha ao iniciar servidor", "jar", jarPath, "erro", err)
		return fmt.Errorf("falha ao iniciar servidor: %w", err)
	}

	// Salva o PID para uso pelos comandos stop/status
	if err := savePID(c.port, cmd.Process.Pid); err != nil {
		slog.Warn("não foi possível salvar PID", "erro", err)
	}

	// Aguarda o servidor ficar pronto (health check real)
	fmt.Fprintf(c.out, "Aguardando o servidor assinador ficar pronto na porta %d...\n", c.port)
	if err := waitForReady(healthURL, 30*time.Second); err != nil {
		fmt.Fprintf(c.errOut, "Erro do sistema: servidor iniciado mas não ficou pronto a tempo.\n")
		fmt.Fprintf(c.errOut, "Causa: %v\n", err)
		fmt.Fprintf(c.errOut, "Verifique os logs do servidor para mais detalhes.\n")
		slog.Error("servidor não ficou pronto", "url", healthURL, "erro", err)
		return fmt.Errorf("servidor não ficou pronto: %w", err)
	}

	fmt.Fprintf(c.out, "Servidor assinador iniciado com sucesso na porta %d (PID: %d).\n", c.port, cmd.Process.Pid)
	slog.Info("servidor iniciado", "porta", c.port, "pid", cmd.Process.Pid)
	return nil
}

// findJar localiza o assinador.jar nas localizações conhecidas.
func (c *StartCmd) findJar() (string, error) {
	if c.jarPath != "" {
		if _, err := os.Stat(c.jarPath); err == nil {
			return c.jarPath, nil
		}
		return "", fmt.Errorf("JAR não encontrado em %q", c.jarPath)
	}

	// Caminhos de busca automática
	candidates := []string{
		"docs/aulas/projetos/assinador-java/target/assinador-java-1.0.0-SNAPSHOT-jar-with-dependencies.jar",
		"assinador.jar",
		filepath.Join(os.Getenv("HOME"), ".assinador", "assinador.jar"),
	}

	for _, p := range candidates {
		if _, err := os.Stat(p); err == nil {
			return p, nil
		}
	}

	return "", fmt.Errorf("assinador.jar não encontrado nos caminhos padrão; use --jar para especificar o caminho")
}

// waitForReady aguarda até que o endpoint de health retorne status UP ou o timeout expire.
func waitForReady(healthURL string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	for time.Now().Before(deadline) {
		resp, err := httpGet(healthURL)
		if err == nil {
			if status, ok := resp["status"].(string); ok && status == "UP" {
				return nil
			}
		}
		time.Sleep(500 * time.Millisecond)
	}
	return fmt.Errorf("timeout de %s expirou aguardando servidor ficar pronto", timeout)
}

// isPortOccupied verifica se a porta está sendo usada por algum processo.
func isPortOccupied(port int) bool {
	// Tenta conectar rapidamente; se conectar, alguém está na porta
	// mas NÃO é o nosso servidor (já verificamos health check acima)
	_, err := httpGet(fmt.Sprintf("http://localhost:%d/health", port))
	// Se retornou erro de "not found" ou "status 404", a porta está ocupada por outro processo
	if err != nil && strings.Contains(err.Error(), "status HTTP 404") {
		return true
	}
	// Erro de conexão recusada = porta livre
	return false
}

// checkJavaVersion verifica se a versão do Java é >= 21.
func checkJavaVersion(errOut io.Writer) error {
	out, err := exec.Command("java", "--version").Output()
	if err != nil {
		// Tenta com -version (Java antigo)
		out, err = exec.Command("java", "-version").Output()
		if err != nil {
			return fmt.Errorf("não foi possível verificar versão do Java: %v", err)
		}
	}

	output := string(out)
	slog.Info("versão do Java detectada", "saída", strings.TrimSpace(output))

	// Verifica se a versão é >= 21
	if strings.Contains(output, "version \"1.") || strings.Contains(output, "version \"8") ||
		strings.Contains(output, "version \"11") || strings.Contains(output, "version \"17") {
		fmt.Fprintf(errOut, "Aviso: versão do Java detectada pode ser inferior a 21. O assinador requer Java 21+.\n")
		return fmt.Errorf("versão do Java pode ser inferior ao mínimo requerido (21)")
	}

	return nil
}

// pidFilePath retorna o caminho do arquivo PID para uma porta.
func pidFilePath(port int) string {
	return filepath.Join(os.TempDir(), fmt.Sprintf("assinador-%d.pid", port))
}

// savePID salva o PID do processo servidor em arquivo temporário.
func savePID(port, pid int) error {
	return os.WriteFile(pidFilePath(port), []byte(strconv.Itoa(pid)), 0600)
}

// loadPID carrega o PID salvo para uma porta.
func loadPID(port int) (int, error) {
	data, err := os.ReadFile(pidFilePath(port))
	if err != nil {
		return 0, err
	}
	pid, err := strconv.Atoi(strings.TrimSpace(string(data)))
	if err != nil {
		return 0, fmt.Errorf("arquivo PID corrompido: %w", err)
	}
	return pid, nil
}

// removePID remove o arquivo PID de uma porta.
func removePID(port int) {
	_ = os.Remove(pidFilePath(port))
}

// killProcess encerra um processo pelo PID de forma cross-platform.
func killProcess(pid int) error {
	proc, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("processo PID %d não encontrado: %w", pid, err)
	}
	if runtime.GOOS == "windows" {
		return exec.Command("taskkill", "/F", "/PID", strconv.Itoa(pid)).Run()
	}
	return proc.Signal(os.Interrupt) // SIGINT para shutdown graceful
}
