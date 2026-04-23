package command

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
	"strings"
)

// ValidateCmd implementa o comando "validate" (validação de assinatura).
type ValidateCmd struct {
	out       io.Writer
	input     string
	signature string
	mode      string
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
  --help            Show this help message

Example:
  assinatura validate --input document.pdf --signature document.sig`
}

// Run executa o comando de validação.
func (c *ValidateCmd) Run(args []string) error {
	fs := flag.NewFlagSet("validate", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	fs.StringVar(&c.input, "input", "", "Path to the file to validate")
	fs.StringVar(&c.signature, "signature", "", "Path to the signature file")
	fs.StringVar(&c.mode, "mode", "local", "Invocation mode: 'local' or 'http'")

	err := fs.Parse(args)
	if err != nil {
		fmt.Fprintf(c.out, "[MS-03] Falha: Parâmetro obrigatório ausente ou sintaxe de comando inválida.\n")
		fmt.Fprintln(c.out, c.Help())
		return err
	}

	if c.input == "" || c.signature == "" {
		fmt.Fprintf(c.out, "[MS-03] Falha: Parâmetros --input e --signature são obrigatórios.\n")
		return fmt.Errorf("missing required parameters")
	}

	if c.mode != "local" && c.mode != "http" {
		fmt.Fprintf(c.out, "[MS-03] Falha: Modo '%s' inválido. Use 'local' ou 'http'.\n", c.mode)
		return fmt.Errorf("invalid mode: %s", c.mode)
	}

	if c.mode == "http" {
		fmt.Fprintf(c.out, "[MS-05] Modo HTTP ainda não está implementado. Use --mode local.\n")
		return fmt.Errorf("http mode not supported yet")
	}

	inputData, err := os.ReadFile(c.input)
	if err != nil {
		fmt.Fprintf(c.out, "[MS-03] Falha: não foi possível ler o arquivo de entrada '%s'.\n", c.input)
		return err
	}

	sigData, err := os.ReadFile(c.signature)
	if err != nil {
		fmt.Fprintf(c.out, "[MS-03] Falha: não foi possível ler o arquivo de assinatura '%s'.\n", c.signature)
		return err
	}

	sigText := strings.TrimSpace(string(sigData))
	sigBytes, err := hex.DecodeString(sigText)
	if err != nil {
		fmt.Fprintf(c.out, "[MS-03] Falha: o arquivo de assinatura '%s' está em formato inválido.\n", c.signature)
		return err
	}

	hash := sha256.Sum256(inputData)

	fmt.Fprintf(c.out, "Validando assinatura (Modo: %s)...\n", c.mode)
	fmt.Fprintf(c.out, "Arquivo de entrada: %s\n", c.input)
	fmt.Fprintf(c.out, "Arquivo de assinatura: %s\n", c.signature)

	if !equalBytes(hash[:], sigBytes) {
		fmt.Fprintf(c.out, "[MS-01] Operação concluída com erro.\n")
		fmt.Fprintf(c.out, "Status: INVÁLIDA\n")
		return fmt.Errorf("signature invalid")
	}

	fmt.Fprintf(c.out, "[MS-01] Operação concluída com sucesso.\n")
	fmt.Fprintf(c.out, "Status: VÁLIDA\n")

	return nil
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
