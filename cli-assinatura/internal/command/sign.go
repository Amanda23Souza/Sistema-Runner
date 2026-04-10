package command

import (
	"fmt"
	"io"
	"os"
)

// SignCmd implementa o comando "sign" (assinatura digital).
// Atualmente é um stub aguardando implementação completa.
type SignCmd struct {
	out io.Writer
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
  --help            Show this help message

Example:
  assinatura sign --input document.pdf --output document.sig`
}

// Run executa o comando de assinatura.
func (c *SignCmd) Run(args []string) error {
	fmt.Fprintln(c.out, "Comando 'sign' ainda em construção.")
	fmt.Fprintln(c.out, "Use 'assinatura sign --help' para mais informações.")
	return nil
}
