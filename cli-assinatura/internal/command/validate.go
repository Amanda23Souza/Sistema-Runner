package command

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// ValidateCmd implementa o comando "validate" (validação de assinatura).
type ValidateCmd struct {
	out        io.Writer
	input      string
	signature  string
	mode       string
	formatJSON bool
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
		out: os.Stdout,
	}
}

// Help retorna a descrição de uso do comando validate.
func (c *ValidateCmd) Help() string {
	return `Usage: assinatura validate [OPTIONS]

Validate a digital signature for a file.

Options:
  --input FILE      Path to the file to validate (required)
  --signature FILE  Path to the signature file (required)
  --mode MODE       Invocation mode: 'local' or 'http' (default: 'local')
	--json            Output a structured JSON summary
  --help            Show this help message

Example:
	assinatura validate --input document.pdf --signature document.sig --json`
}

// Run executa o comando de validação.
func (c *ValidateCmd) Run(args []string) error {
	fs := flag.NewFlagSet("validate", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	fs.StringVar(&c.input, "input", "", "Path to the file to validate")
	fs.StringVar(&c.signature, "signature", "", "Path to the signature file")
	fs.StringVar(&c.mode, "mode", "local", "Invocation mode: 'local' or 'http'")
	fs.BoolVar(&c.formatJSON, "json", false, "Output a structured JSON summary")

	err := fs.Parse(args)
	if err != nil {
		fmt.Fprintf(c.out, "[MS-03] Falha: Parâmetro obrigatório ausente ou sintaxe de comando inválida.\n")
		fmt.Fprintln(c.out, c.Help())
		return err
	}

	if c.input == "" || c.signature == "" {
		if c.input == "" {
			fmt.Fprintf(c.out, "Erro de validação: o campo --input é obrigatório.\n")
		}
		if c.signature == "" {
			fmt.Fprintf(c.out, "Erro de validação: o campo --signature é obrigatório.\n")
		}
		return fmt.Errorf("missing required parameters")
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

	sigData, err := os.ReadFile(c.signature)
	if err != nil {
		fmt.Fprintf(c.out, "Erro de validação: não foi possível ler o arquivo de assinatura %q.\n", c.signature)
		return err
	}

	sigText := strings.TrimSpace(string(sigData))
	sigBytes, err := hex.DecodeString(sigText)
	if err != nil {
		fmt.Fprintf(c.out, "Erro de validação: o arquivo de assinatura %q não está em formato hexadecimal válido.\n", c.signature)
		return err
	}

	hash := sha256.Sum256(inputData)
	execution := "simulação local"
	if c.mode == "http" {
		execution = "simulação HTTP interna"
	}

	fmt.Fprintf(c.out, "Validando assinatura (%s)...\n", execution)
	fmt.Fprintf(c.out, "Arquivo de entrada: %s\n", c.input)
	fmt.Fprintf(c.out, "Arquivo de assinatura: %s\n", c.signature)

	if !equalBytes(hash[:], sigBytes) {
		fmt.Fprintf(c.out, "Resultado: assinatura inválida.\n")
		fmt.Fprintf(c.out, "Status: INVÁLIDA\n")
		if c.formatJSON {
			return c.outputJSON(validateResult{
				Operation: "validate",
				Mode:      c.mode,
				Input:     c.input,
				Signature: c.signature,
				Execution: execution,
				Status:    "invalid",
				Valid:     false,
			})
		}
		return fmt.Errorf("signature invalid")
	}

	fmt.Fprintf(c.out, "Resultado: assinatura válida.\n")
	fmt.Fprintf(c.out, "Status: VÁLIDA\n")

	if c.formatJSON {
		return c.outputJSON(validateResult{
			Operation: "validate",
			Mode:      c.mode,
			Input:     c.input,
			Signature: c.signature,
			Execution: execution,
			Status:    "valid",
			Valid:     true,
		})
	}

	return nil
}

func (c *ValidateCmd) outputJSON(result validateResult) error {
	encoder := json.NewEncoder(c.out)
	encoder.SetIndent("", "  ")
	return encoder.Encode(result)
}

func equalBytes(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
