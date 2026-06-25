package command

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log/slog"
	"os"
	"strings"
)

// ValidateCmd implementa o comando "validate" (validação de assinatura).
type ValidateCmd struct {
	out        io.Writer
	errOut     io.Writer
	input      string
	signature  string
	mode       string
	port       int
	formatJSON bool
	verbose    bool
}

type validateResult struct {
	Operation string `json:"operation"`
	Mode      string `json:"mode"`
	Input     string `json:"input"`
	Signature string `json:"signature"`
	Execution string `json:"execution"`
	Status    string `json:"status"`
	Valid     bool   `json:"valid"`
}

// NewValidateCmd cria uma nova instância do comando de validação.
func NewValidateCmd() *ValidateCmd {
	return &ValidateCmd{
		out:    os.Stdout,
		errOut: os.Stderr,
	}
}

// Help retorna a descrição de uso do comando validate.
func (c *ValidateCmd) Help() string {
	return `Usage: assinatura validate [OPTIONS]

Valida a assinatura digital de um arquivo.

O modo padrão é 'http': o CLI se comunica com o assinador.jar em execução.
Para validar localmente (sem servidor), use --mode local.

Options:
  --input FILE      Caminho do arquivo a ser validado (obrigatório)
  --signature FILE  Caminho do arquivo de assinatura (obrigatório)
  --mode MODE       Modo de invocação: 'http' (padrão) ou 'local'
  --port PORT       Porta do servidor assinador em modo http (padrão: 8080)
  --json            Saída em formato JSON estruturado
  --verbose         Habilita saída detalhada
  --help            Exibe esta mensagem de ajuda

Examples:
  assinatura validate --input documento.pdf --signature documento.sig
  assinatura validar --input documento.pdf --signature documento.sig --json
  assinatura validate --input documento.pdf --signature documento.sig --mode local`
}

// Run executa o comando de validação.
func (c *ValidateCmd) Run(args []string) error {
	fs := flag.NewFlagSet("validate", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	fs.StringVar(&c.input, "input", "", "Caminho do arquivo a ser validado")
	fs.StringVar(&c.signature, "signature", "", "Caminho do arquivo de assinatura")
	fs.StringVar(&c.mode, "mode", "http", "Modo de invocação: 'http' (padrão) ou 'local'")
	fs.IntVar(&c.port, "port", 8080, "Porta do servidor assinador (modo http)")
	fs.BoolVar(&c.formatJSON, "json", false, "Saída em formato JSON estruturado")
	fs.BoolVar(&c.verbose, "verbose", false, "Habilita saída detalhada")

	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			fmt.Fprintln(c.out, c.Help())
			return nil
		}
	}

	if err := fs.Parse(args); err != nil {
		fmt.Fprintf(c.errOut, "[MS-03] Erro do usuário: parâmetro inválido ou sintaxe incorreta.\n")
		fmt.Fprintln(c.errOut, c.Help())
		return &UserError{msg: "sintaxe de comando inválida"}
	}

	if c.input == "" {
		fmt.Fprintf(c.errOut, "Erro do usuário: o parâmetro --input é obrigatório.\n")
		fmt.Fprintf(c.errOut, "Dica: assinatura validate --input <arquivo> --signature <arquivo.sig>\n")
		return &UserError{msg: "parâmetro obrigatório ausente: --input"}
	}
	if c.signature == "" {
		fmt.Fprintf(c.errOut, "Erro do usuário: o parâmetro --signature é obrigatório.\n")
		fmt.Fprintf(c.errOut, "Dica: assinatura validate --input <arquivo> --signature <arquivo.sig>\n")
		return &UserError{msg: "parâmetro obrigatório ausente: --signature"}
	}

	if c.mode != "local" && c.mode != "http" {
		fmt.Fprintf(c.errOut, "Erro do usuário: --mode deve ser 'http' ou 'local' (recebido: %q).\n", c.mode)
		return &UserError{msg: fmt.Sprintf("modo inválido: %s", c.mode)}
	}

	if c.verbose {
		slog.Info("iniciando validação", "input", c.input, "signature", c.signature, "mode", c.mode)
	}

	if c.mode == "http" {
		return c.runHTTP()
	}
	return c.runLocal()
}

// runHTTP lê o arquivo de entrada e a assinatura, enviando ambos ao servidor via HTTP.
func (c *ValidateCmd) runHTTP() error {
	inputData, err := os.ReadFile(c.input)
	if err != nil {
		fmt.Fprintf(c.errOut, "Erro do sistema: não foi possível ler o arquivo de entrada %q.\n", c.input)
		fmt.Fprintf(c.errOut, "Causa: %v\n", err)
		slog.Error("falha ao ler arquivo de entrada (modo http)", "input", c.input, "erro", err)
		return fmt.Errorf("falha ao ler arquivo de entrada: %w", err)
	}

	sigData, err := os.ReadFile(c.signature)
	if err != nil {
		fmt.Fprintf(c.errOut, "Erro do sistema: não foi possível ler o arquivo de assinatura %q.\n", c.signature)
		fmt.Fprintf(c.errOut, "Causa: %v\n", err)
		slog.Error("falha ao ler arquivo de assinatura", "signature", c.signature, "erro", err)
		return fmt.Errorf("falha ao ler arquivo de assinatura: %w", err)
	}

	url := fmt.Sprintf("http://localhost:%d/validate", c.port)
	bodyBytes, err := json.Marshal(map[string]string{
		"content":   string(inputData),
		"signature": strings.TrimSpace(string(sigData)),
	})
	if err != nil {
		return fmt.Errorf("erro ao serializar requisição: %w", err)
	}

	if c.verbose {
		slog.Info("enviando requisição HTTP", "url", url)
	}

	resp, err := httpPost(url, string(bodyBytes))
	if err != nil {
		fmt.Fprintf(c.errOut, "Erro do sistema: falha ao conectar ao servidor assinador em localhost:%d.\n", c.port)
		fmt.Fprintf(c.errOut, "Causa: %v\n", err)
		fmt.Fprintf(c.errOut, "Como resolver: verifique se o servidor está rodando com 'assinatura start --port %d'.\n", c.port)
		slog.Error("falha na conexão HTTP", "url", url, "erro", err)
		return fmt.Errorf("falha na conexão com o servidor: %w", err)
	}

	valid, _ := resp["valid"].(bool)
	return c.printValidationResult(valid, "http")
}

// runLocal valida localmente (SHA-256 simulado).
func (c *ValidateCmd) runLocal() error {
	inputData, err := os.ReadFile(c.input)
	if err != nil {
		fmt.Fprintf(c.errOut, "Erro do sistema: não foi possível ler o arquivo de entrada %q.\n", c.input)
		fmt.Fprintf(c.errOut, "Causa: %v\n", err)
		fmt.Fprintf(c.errOut, "Como resolver: verifique se o arquivo existe e se você tem permissão de leitura.\n")
		slog.Error("falha ao ler arquivo de entrada", "input", c.input, "erro", err)
		return fmt.Errorf("falha ao ler arquivo de entrada: %w", err)
	}

	sigData, err := os.ReadFile(c.signature)
	if err != nil {
		fmt.Fprintf(c.errOut, "Erro do sistema: não foi possível ler o arquivo de assinatura %q.\n", c.signature)
		fmt.Fprintf(c.errOut, "Causa: %v\n", err)
		slog.Error("falha ao ler arquivo de assinatura", "signature", c.signature, "erro", err)
		return fmt.Errorf("falha ao ler arquivo de assinatura: %w", err)
	}

	expected := computeSHA256(inputData)
	actual := strings.TrimSpace(string(sigData))
	valid := expected == actual

	return c.printValidationResult(valid, "local")
}

func (c *ValidateCmd) printValidationResult(valid bool, mode string) error {
	fmt.Fprintf(c.out, "Validação concluída (modo %s).\n", mode)
	fmt.Fprintf(c.out, "Arquivo de entrada: %s\n", c.input)
	fmt.Fprintf(c.out, "Arquivo de assinatura: %s\n", c.signature)

	if valid {
		fmt.Fprintf(c.out, "Resultado: assinatura VÁLIDA.\n")
		fmt.Fprintf(c.out, "Status: VÁLIDA\n")
	} else {
		fmt.Fprintf(c.out, "Resultado: assinatura INVÁLIDA.\n")
		fmt.Fprintf(c.out, "Status: INVÁLIDA\n")
	}

	if c.formatJSON {
		status := "valid"
		if !valid {
			status = "invalid"
		}
		return c.outputJSON(validateResult{
			Operation: "validate",
			Mode:      mode,
			Input:     c.input,
			Signature: c.signature,
			Execution: fmt.Sprintf("modo %s", mode),
			Status:    status,
			Valid:      valid,
		})
	}

	if !valid {
		return fmt.Errorf("assinatura inválida")
	}
	return nil
}

func (c *ValidateCmd) outputJSON(result validateResult) error {
	encoder := json.NewEncoder(c.out)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}
