package command

import (
	"fmt"
	"io"
	"os"
)

// ValidateCmd implementa o comando "validate" (validação de assinatura).
// Atualmente é um stub aguardando implementação completa.
type ValidateCmd struct {
	out io.Writer
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
	fmt.Fprintln(c.out, "Comando 'validate' ainda em construção.")
	fmt.Fprintln(c.out, "Use 'assinatura validate --help' para mais informações.")
	return nil
}
