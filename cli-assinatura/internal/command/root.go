// Package command define e orquestra os comandos da aplicação CLI.
package command

import (
	"fmt"
	"io"
	"os"
)

// RootCmd é o comando raiz que orquestra todos os subcomandos.
type RootCmd struct {
	out      io.Writer
	commands map[string]Command
}

// NewRootCmd cria uma nova instância do comando raiz.
func NewRootCmd() *RootCmd {
	root := &RootCmd{
		out:      os.Stdout,
		commands: make(map[string]Command),
	}

	// Registra os comandos disponíveis
	root.commands["version"] = NewVersionCmd()
	root.commands["sign"] = NewSignCmd()
	root.commands["validate"] = NewValidateCmd()

	return root
}

// Run executa o comando apropriado baseado nos argumentos fornecidos.
func (c *RootCmd) Run(args []string) error {
	if len(args) == 0 {
		fmt.Fprintln(c.out, c.Help())
		return fmt.Errorf("no command specified")
	}

	cmdName := args[0]

	// Verifica se é solicitado help geral
	if cmdName == "help" || cmdName == "--help" || cmdName == "-h" {
		fmt.Fprintln(c.out, c.Help())
		return nil
	}

	// Suporte a opção global de versão
	if cmdName == "--version" {
		return c.commands["version"].Run([]string{})
	}

	// Busca o comando registrado
	cmd, exists := c.commands[cmdName]
	if !exists {
		fmt.Fprintf(c.out, "Erro: comando desconhecido '%s'\n", cmdName)
		fmt.Fprintln(c.out, "Use 'assinatura --help' para ver a lista de comandos.")
		return fmt.Errorf("command not found: %s", cmdName)
	}

	// Executa o comando com os argumentos restantes
	return cmd.Run(args[1:])
}

// Help retorna a mensagem de uso geral da aplicação.
func (c *RootCmd) Help() string {
	return `assinatura - CLI para operações de assinatura digital

Usage: assinatura <command> [OPTIONS]

Commands:
  version      Display the version of assinatura CLI
  sign         Create a digital signature for a file
  validate     Validate a digital signature

Global Options:
  --help       Show this help message
  --version    Display version information

Examples:
  assinatura version
  assinatura sign --input document.pdf
  assinatura validate --input document.pdf --signature document.sig

Use 'assinatura <command> --help' for more information about a command.`
}
