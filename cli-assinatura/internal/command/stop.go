package command

import (
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
)

// StopCmd implementa o comando "stop" — encerra o servidor assinador.
type StopCmd struct {
	out    io.Writer
	errOut io.Writer
	port   int
}

// NewStopCmd cria uma nova instância do comando stop.
func NewStopCmd() *StopCmd {
	return &StopCmd{
		out:    os.Stdout,
		errOut: os.Stderr,
	}
}

// Help retorna a descrição de uso do comando stop.
func (c *StopCmd) Help() string {
	return `Usage: assinatura stop [OPTIONS]

Encerra o servidor assinador em execução na porta indicada.

Options:
  --port PORT   Porta do servidor a ser encerrado (padrão: 8080)
  --help        Exibe esta mensagem de ajuda

Examples:
  assinatura stop
  assinatura stop --port 9090`
}

// Run executa o comando stop.
func (c *StopCmd) Run(args []string) error {
	fs := flag.NewFlagSet("stop", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.IntVar(&c.port, "port", 8080, "Porta do servidor")

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

	// Verifica se o servidor está rodando via health check
	healthURL := fmt.Sprintf("http://localhost:%d/health", c.port)
	if _, err := httpGet(healthURL); err != nil {
		fmt.Fprintf(c.out, "Nenhum servidor assinador encontrado na porta %d.\n", c.port)
		slog.Info("nenhum servidor ativo na porta", "porta", c.port)
		return nil
	}

	// Carrega o PID salvo
	pid, err := loadPID(c.port)
	if err != nil {
		fmt.Fprintf(c.errOut, "Erro do sistema: não foi possível encontrar o PID do servidor na porta %d.\n", c.port)
		fmt.Fprintf(c.errOut, "Causa: %v\n", err)
		fmt.Fprintf(c.errOut, "Como resolver: encerre manualmente o processo na porta %d.\n", c.port)
		slog.Error("arquivo PID não encontrado", "porta", c.port, "erro", err)
		return fmt.Errorf("PID não encontrado para porta %d: %w", c.port, err)
	}

	// Encerra o processo
	if err := killProcess(pid); err != nil {
		fmt.Fprintf(c.errOut, "Erro do sistema: falha ao encerrar o processo PID %d.\n", pid)
		fmt.Fprintf(c.errOut, "Causa: %v\n", err)
		slog.Error("falha ao encerrar processo", "pid", pid, "erro", err)
		return fmt.Errorf("falha ao encerrar servidor (PID %d): %w", pid, err)
	}

	removePID(c.port)
	fmt.Fprintf(c.out, "Servidor assinador na porta %d encerrado com sucesso (PID: %d).\n", c.port, pid)
	slog.Info("servidor encerrado", "porta", c.port, "pid", pid)
	return nil
}

// StatusCmd implementa o comando "status" — verifica o estado do servidor assinador.
type StatusCmd struct {
	out        io.Writer
	errOut     io.Writer
	port       int
	formatJSON bool
}

// NewStatusCmd cria uma nova instância do comando status.
func NewStatusCmd() *StatusCmd {
	return &StatusCmd{
		out:    os.Stdout,
		errOut: os.Stderr,
	}
}

// Help retorna a descrição de uso do comando status.
func (c *StatusCmd) Help() string {
	return `Usage: assinatura status [OPTIONS]

Verifica o estado do servidor assinador.
Distingue entre "processo subiu" e "pronto para receber requisições".

Options:
  --port PORT   Porta do servidor (padrão: 8080)
  --json        Saída em formato JSON estruturado
  --help        Exibe esta mensagem de ajuda

Examples:
  assinatura status
  assinatura status --port 9090 --json`
}

// Run executa o comando status.
func (c *StatusCmd) Run(args []string) error {
	fs := flag.NewFlagSet("status", flag.ContinueOnError)
	fs.SetOutput(io.Discard)
	fs.IntVar(&c.port, "port", 8080, "Porta do servidor")
	fs.BoolVar(&c.formatJSON, "json", false, "Saída em formato JSON")

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

	healthURL := fmt.Sprintf("http://localhost:%d/health", c.port)
	pid, pidErr := loadPID(c.port)

	resp, err := httpGet(healthURL)
	if err != nil {
		fmt.Fprintf(c.out, "Status: INATIVO\n")
		fmt.Fprintf(c.out, "Porta: %d\n", c.port)
		fmt.Fprintf(c.out, "Motivo: %v\n", err)
		if c.formatJSON {
			fmt.Fprintf(c.out, `{"status":"DOWN","port":%d,"healthy":false}`+"\n", c.port)
		}
		slog.Info("servidor inativo", "porta", c.port)
		return nil
	}

	status, _ := resp["status"].(string)
	healthy := status == "UP"

	fmt.Fprintf(c.out, "Status: ATIVO\n")
	fmt.Fprintf(c.out, "Porta: %d\n", c.port)
	fmt.Fprintf(c.out, "Health: %s\n", status)
	if pidErr == nil {
		fmt.Fprintf(c.out, "PID: %d\n", pid)
	}

	if c.formatJSON {
		pidStr := "null"
		if pidErr == nil {
			pidStr = fmt.Sprintf("%d", pid)
		}
		fmt.Fprintf(c.out, `{"status":"UP","port":%d,"healthy":%v,"pid":%s}`+"\n",
			c.port, healthy, pidStr)
	}

	slog.Info("servidor ativo", "porta", c.port, "health", status)
	return nil
}
