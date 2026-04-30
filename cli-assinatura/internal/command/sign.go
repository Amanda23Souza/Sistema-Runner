package command

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"io"
	"os"
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

	if c.input == "" {
		fmt.Fprintf(c.out, "[MS-03] Falha: O parâmetro --input é obrigatório.\n")
		return fmt.Errorf("missing required parameter: --input")
	}

	if c.output == "" {
		c.output = c.input + ".sig"
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

	hash := sha256.Sum256(inputData)
	signature := hex.EncodeToString(hash[:])

	if err := os.WriteFile(c.output, []byte(signature), 0644); err != nil {
		fmt.Fprintf(c.out, "[MS-03] Falha: não foi possível escrever o arquivo de assinatura '%s'.\n", c.output)
		return err
	}

	fmt.Fprintf(c.out, "Processando assinatura (Modo: %s)...\n", c.mode)
	fmt.Fprintf(c.out, "Arquivo de entrada: %s\n", c.input)
	fmt.Fprintf(c.out, "Assinatura gerada em: %s\n", c.output)
	fmt.Fprintf(c.out, "[MS-01] Operação concluída com sucesso.\n")

	return nil
}
