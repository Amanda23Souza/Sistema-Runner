package command

import (
	"flag"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

// SignCmd implementa o comando "sign" (assinatura digital).
type SignCmd struct {
	out    io.Writer
	input  string
	output string
	mode   string
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
	fs := flag.NewFlagSet("sign", flag.ContinueOnError)
	fs.SetOutput(io.Discard)

	fs.StringVar(&c.input, "input", "", "Path to the file to sign")
	fs.StringVar(&c.output, "output", "", "Path to save the signature")
	fs.StringVar(&c.mode, "mode", "local", "Invocation mode: 'local' or 'http'")

	err := fs.Parse(args)
	if err != nil {
		fmt.Fprintf(c.out, "[MS-03] Falha: Parâmetro obrigatório ausente ou sintaxe de comando inválida.\n")
		fmt.Fprintln(c.out, c.Help())
		return err
	}

	// Validação de parâmetros (US-02 / RN-US002-01)
	if c.input == "" {
		fmt.Fprintf(c.out, "[MS-03] Falha: O parâmetro --input é obrigatório.\n")
		return fmt.Errorf("missing required parameter: --input")
	}

	// Definir output padrão se não fornecido
	if c.output == "" {
		c.output = c.input + ".sig"
	}

	// Validação de Modo (US-01 / RN-US01-01)
	if c.mode != "local" && c.mode != "http" {
		fmt.Fprintf(c.out, "[MS-03] Falha: Modo '%s' inválido. Use 'local' ou 'http'.\n", c.mode)
		return fmt.Errorf("invalid mode: %s", c.mode)
	}

	// Simulação de execução (Mock conforme US-02)
	fmt.Fprintf(c.out, "Processando assinatura (Modo: %s)...\n", c.mode)
	fmt.Fprintf(c.out, "Arquivo de entrada: %s\n", c.input)

	// Simulação de Sucesso (Mock da US-02)
	fmt.Fprintf(c.out, "[MS-01] Operação concluída com sucesso.\n")
	fmt.Fprintf(c.out, "[MS-03] Criação de assinatura simulada concluída com sucesso.\n")
	fmt.Fprintf(c.out, "Assinatura gerada em: %s\n", c.output)

	return nil
}
