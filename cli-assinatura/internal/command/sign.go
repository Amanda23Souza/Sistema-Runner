package command

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
)

// SignCmd implementa o comando "sign" (assinatura digital).
type SignCmd struct {
	out        io.Writer
	input      string
	output     string
	mode       string
	formatJSON bool
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
		out: os.Stdout,
	}
}

// Help retorna a descrição de uso do comando sign.
func (c *SignCmd) Help() string {
	return `Usage: assinatura sign [OPTIONS]

Create a digital signature for a file.

Options:
  --input FILE      Path to the file to sign (required)
  --output FILE     Path to save the signature (default: <input>.sig)
  --mode MODE       Invocation mode: 'local' or 'http' (default: 'local')
	--json            Output a structured JSON summary
  --help            Show this help message

Example:
	assinatura sign --input document.pdf --output document.sig --json`
}

// Run executa o comando de assinatura.
func (c *SignCmd) Run(args []string) error {
	fs := flag.NewFlagSet("sign", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	fs.StringVar(&c.input, "input", "", "Path to the file to sign")
	fs.StringVar(&c.output, "output", "", "Path to save the signature")
	fs.StringVar(&c.mode, "mode", "local", "Invocation mode: 'local' or 'http'")
	fs.BoolVar(&c.formatJSON, "json", false, "Output a structured JSON summary")

	err := fs.Parse(args)
	if err != nil {
		fmt.Fprintf(c.out, "[MS-03] Falha: Parâmetro obrigatório ausente ou sintaxe de comando inválida.\n")
		fmt.Fprintln(c.out, c.Help())
		return err
	}

	if c.input == "" {
		fmt.Fprintf(c.out, "Erro de validação: o campo --input é obrigatório.\n")
		return fmt.Errorf("missing required parameter: --input")
	}

	if c.output == "" {
		c.output = c.input + ".sig"
	}

	if c.mode != "local" && c.mode != "http" {
		fmt.Fprintf(c.out, "Erro de validação: o campo --mode deve ser 'local' ou 'http' (valor recebido: %q).\n", c.mode)
		return fmt.Errorf("invalid mode: %s", c.mode)
	}

	inputData, err := os.ReadFile(c.input)
	if err != nil {
		fmt.Fprintf(c.out, "Erro de validação: não foi possível ler o arquivo de entrada %q.\n", c.input)
		return err
	}

	hash := sha256.Sum256(inputData)
	signature := hex.EncodeToString(hash[:])

	if err := os.WriteFile(c.output, []byte(signature), 0644); err != nil {
		fmt.Fprintf(c.out, "Erro de validação: não foi possível escrever o arquivo de saída %q.\n", c.output)
		return err
	}

	execution := "simulação local"
	if c.mode == "http" {
		execution = "simulação HTTP interna"
	}

	fmt.Fprintf(c.out, "Processando assinatura (%s)...\n", execution)
	fmt.Fprintf(c.out, "Arquivo de entrada: %s\n", c.input)
	fmt.Fprintf(c.out, "Assinatura gerada em: %s\n", c.output)
	fmt.Fprintf(c.out, "Resultado: assinatura simulada com sucesso.\n")

	if c.formatJSON {
		return c.outputJSON(signResult{
			Operation: "sign",
			Mode:      c.mode,
			Input:     c.input,
			Output:    c.output,
			Signature: signature,
			Execution: execution,
			Status:    "success",
		})
	}

	return nil
}

func (c *SignCmd) outputJSON(result signResult) error {
	encoder := json.NewEncoder(c.out)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}
