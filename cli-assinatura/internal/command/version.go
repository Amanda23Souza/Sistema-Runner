package command

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/Amanda23Souza/Sistema-Runner/cli-assinatura/internal/version"
)

// VersionCmd implementa o comando "version".
type VersionCmd struct {
	out       io.Writer // saída para escrever resultado
	quiet     bool
	formatJSON bool
}

// NewVersionCmd cria uma nova instância do comando de versão.
func NewVersionCmd() *VersionCmd {
	return &VersionCmd{
		out: os.Stdout,
	}
}

// Help retorna a descrição de uso do comando version.
func (c *VersionCmd) Help() string {
	return `Usage: assinatura version [OPTIONS]

Display the version of assinatura CLI.

Options:
  --quiet      Output only the version number (no prefix)
  --json       Output version information as JSON
  --help       Show this help message`
}

// Run executa o comando de versão.
func (c *VersionCmd) Run(args []string) error {
	fs := flag.NewFlagSet("version", flag.ContinueOnError)
	fs.SetOutput(io.Discard) // Silencia erros padrão do flag
	fs.BoolVar(&c.quiet, "quiet", false, "Output only the version number")
	fs.BoolVar(&c.formatJSON, "json", false, "Output version as JSON")

	err := fs.Parse(args)
	if err != nil {
		return err
	}

	// Manejar --help manualmente (flag não captura --help automaticamente)
	for _, arg := range args {
		if arg == "--help" || arg == "-h" {
			fmt.Fprintln(c.out, c.Help())
			return nil
		}
	}

	return c.display()
}

// display escreve a versão no formato solicitado.
func (c *VersionCmd) display() error {
	if c.formatJSON {
		return c.outputJSON()
	}

	if c.quiet {
		fmt.Fprintln(c.out, version.Get())
	} else {
		fmt.Fprintln(c.out, version.GetFull())
	}

	return nil
}

// outputJSON escreve a versão em formato JSON.
func (c *VersionCmd) outputJSON() error {
	data := map[string]string{
		"version": version.Get(),
	}

	encoder := json.NewEncoder(c.out)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}
