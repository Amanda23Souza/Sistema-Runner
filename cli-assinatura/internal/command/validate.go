package command

import (
	"flag"
	"fmt"
	"io"
	"os"
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

	// Validação de parâmetros (US-02 / RN-US002-01)
	if c.input == "" || c.signature == "" {
		fmt.Fprintf(c.out, "[MS-03] Falha: Parâmetros --input e --signature são obrigatórios.\n")
		return fmt.Errorf("missing required parameters")
	}

	// Validação de Modo (US-01 / RN-US01-01)
	if c.mode != "local" && c.mode != "http" {
		fmt.Fprintf(c.out, "[MS-03] Falha: Modo '%s' inválido. Use 'local' ou 'http'.\n", c.mode)
		return fmt.Errorf("invalid mode: %s", c.mode)
	}

	// Simulação de execução (Mock conforme US-02)
	fmt.Fprintf(c.out, "Validando assinatura (Modo: %s)...\n", c.mode)
	fmt.Fprintf(c.out, "Arquivo de entrada: %s\n", c.input)
	fmt.Fprintf(c.out, "Arquivo de assinatura: %s\n", c.signature)

	// Simulação de Sucesso (Mock da US-02)
	fmt.Fprintf(c.out, "[MS-01] Operação concluída com sucesso.\n")
	fmt.Fprintf(c.out, "[MS-04] Validação simulada concluída e resultado retornado.\n")
	fmt.Fprintf(c.out, "Status: VÁLIDA\n")

	return nil
}
