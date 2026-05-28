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

// SignCmd implementa o comando "sign" (assinatura digital).
type SignCmd struct {
	out        io.Writer
	errOut     io.Writer
	input      string
	output     string
	mode       string
	port       int
	formatJSON bool
	verbose    bool
}

type signResult struct {
	Operation string `json:"operation"`
	Mode      string `json:"mode"`
	Input     string `json:"input"`
	Output    string `json:"output"`
	Signature string `json:"signature"`
	Execution string `json:"execution"`
	Status    string `json:"status"`
}

// NewSignCmd cria uma nova instância do comando de assinatura.
func NewSignCmd() *SignCmd {
	return &SignCmd{
		out:    os.Stdout,
		errOut: os.Stderr,
	}
}

// Help retorna a descrição de uso do comando sign.
func (c *SignCmd) Help() string {
	return `Usage: assinatura sign [OPTIONS]

Cria uma assinatura digital para um arquivo.

O modo padrão é 'http': o CLI se comunica com o assinador.jar em execução.
Para invocar o JAR diretamente (sem servidor), use --mode local.

Options:
  --input FILE      Caminho do arquivo a ser assinado (obrigatório)
  --output FILE     Caminho para salvar a assinatura (padrão: <input>.sig)
  --mode MODE       Modo de invocação: 'http' (padrão) ou 'local'
  --port PORT       Porta do servidor assinador em modo http (padrão: 8080)
  --json            Saída em formato JSON estruturado
  --verbose         Habilita saída detalhada
  --help            Exibe esta mensagem de ajuda

Examples:
  assinatura sign --input documento.pdf
  assinatura sign --input documento.pdf --output documento.sig --json
  assinatura criar --input documento.pdf --mode local`
}

// Run executa o comando de assinatura.
func (c *SignCmd) Run(args []string) error {
	fs := flag.NewFlagSet("sign", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	fs.StringVar(&c.input, "input", "", "Caminho do arquivo a ser assinado")
	fs.StringVar(&c.output, "output", "", "Caminho para salvar a assinatura")
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
		fmt.Fprintf(c.errOut, "Dica: assinatura sign --input <arquivo>\n")
		return &UserError{msg: "parâmetro obrigatório ausente: --input"}
	}

	if c.mode != "local" && c.mode != "http" {
		fmt.Fprintf(c.errOut, "Erro do usuário: --mode deve ser 'http' ou 'local' (recebido: %q).\n", c.mode)
		return &UserError{msg: fmt.Sprintf("modo inválido: %s", c.mode)}
	}

	if c.output == "" {
		c.output = c.input + ".sig"
	}

	if c.verbose {
		slog.Info("iniciando assinatura", "input", c.input, "mode", c.mode, "port", c.port)
	}

	if c.mode == "http" {
		return c.runHTTP()
	}
	return c.runLocal()
}

// runHTTP invoca o servidor assinador via HTTP.
func (c *SignCmd) runHTTP() error {
	url := fmt.Sprintf("http://localhost:%d/sign", c.port)
	body := fmt.Sprintf(`{"content": "%s"}`, strings.ReplaceAll(c.input, `"`, `\"`))

	if c.verbose {
		slog.Info("enviando requisição HTTP", "url", url)
	}

	resp, err := httpPost(url, body)
	if err != nil {
		fmt.Fprintf(c.errOut, "Erro do sistema: falha ao conectar ao servidor assinador em localhost:%d.\n", c.port)
		fmt.Fprintf(c.errOut, "Causa: %v\n", err)
		fmt.Fprintf(c.errOut, "Como resolver: verifique se o servidor está rodando com 'assinatura start --port %d'.\n", c.port)
		slog.Error("falha na conexão HTTP", "url", url, "erro", err)
		return fmt.Errorf("falha na conexão com o servidor: %w", err)
	}

	signature, ok := resp["signature"].(string)
	if !ok || signature == "" {
		fmt.Fprintf(c.errOut, "Erro do sistema: resposta malformada do servidor assinador.\n")
		fmt.Fprintf(c.errOut, "Resposta recebida: %v\n", resp)
		slog.Error("resposta malformada do servidor", "resposta", resp)
		return fmt.Errorf("resposta malformada do servidor")
	}

	if err := os.WriteFile(c.output, []byte(signature), 0644); err != nil {
		fmt.Fprintf(c.errOut, "Erro do sistema: não foi possível escrever o arquivo de saída %q.\n", c.output)
		fmt.Fprintf(c.errOut, "Causa: %v\n", err)
		slog.Error("falha ao escrever arquivo de saída", "output", c.output, "erro", err)
		return fmt.Errorf("falha ao gravar assinatura: %w", err)
	}

	c.printSuccess("http", signature)

	if c.formatJSON {
		return c.outputJSON(signResult{
			Operation: "sign",
			Mode:      "http",
			Input:     c.input,
			Output:    c.output,
			Signature: signature,
			Execution: fmt.Sprintf("servidor HTTP localhost:%d", c.port),
			Status:    "success",
		})
	}
	return nil
}

// runLocal invoca o assinador em modo local (SHA-256 simulado).
func (c *SignCmd) runLocal() error {
	inputData, err := os.ReadFile(c.input)
	if err != nil {
		fmt.Fprintf(c.errOut, "Erro do sistema: não foi possível ler o arquivo de entrada %q.\n", c.input)
		fmt.Fprintf(c.errOut, "Causa: %v\n", err)
		fmt.Fprintf(c.errOut, "Como resolver: verifique se o arquivo existe e se você tem permissão de leitura.\n")
		slog.Error("falha ao ler arquivo de entrada", "input", c.input, "erro", err)
		return fmt.Errorf("falha ao ler arquivo de entrada: %w", err)
	}

	signature := computeSHA256(inputData)

	if err := os.WriteFile(c.output, []byte(signature), 0644); err != nil {
		fmt.Fprintf(c.errOut, "Erro do sistema: não foi possível escrever o arquivo de saída %q.\n", c.output)
		fmt.Fprintf(c.errOut, "Causa: %v\n", err)
		slog.Error("falha ao escrever arquivo de saída", "output", c.output, "erro", err)
		return fmt.Errorf("falha ao gravar assinatura: %w", err)
	}

	c.printSuccess("local", signature)

	if c.formatJSON {
		return c.outputJSON(signResult{
			Operation: "sign",
			Mode:      "local",
			Input:     c.input,
			Output:    c.output,
			Signature: signature,
			Execution: "simulação local (SHA-256)",
			Status:    "success",
		})
	}
	return nil
}

func (c *SignCmd) printSuccess(mode, signature string) {
	fmt.Fprintf(c.out, "Assinatura criada com sucesso (modo %s).\n", mode)
	fmt.Fprintf(c.out, "Arquivo de entrada: %s\n", c.input)
	fmt.Fprintf(c.out, "Assinatura salva em: %s\n", c.output)
	if c.verbose {
		fmt.Fprintf(c.out, "Assinatura: %s\n", signature)
	}
}

func (c *SignCmd) outputJSON(result signResult) error {
	encoder := json.NewEncoder(c.out)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}
