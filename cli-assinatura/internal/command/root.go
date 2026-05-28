// Package command define e orquestra os comandos da aplicação CLI.
package command

import (
	"fmt"
	"io"
	"log/slog"
	"os"
)

// RootCmd é o comando raiz que orquestra todos os subcomandos.
type RootCmd struct {
	out      io.Writer
	errOut   io.Writer
	commands map[string]Command
}

// NewRootCmd cria uma nova instância do comando raiz.
func NewRootCmd() *RootCmd {
	root := &RootCmd{
		out:      os.Stdout,
		errOut:   os.Stderr,
		commands: make(map[string]Command),
	}

	// Registra os comandos disponíveis
	root.commands["version"] = NewVersionCmd()
	root.commands["sign"] = NewSignCmd()
	root.commands["criar"] = root.commands["sign"]
	root.commands["validate"] = NewValidateCmd()
	root.commands["validar"] = root.commands["validate"]
	root.commands["start"] = NewStartCmd()
	root.commands["stop"] = NewStopCmd()
	root.commands["status"] = NewStatusCmd()

	return root
}

// Run executa o comando apropriado baseado nos argumentos fornecidos.
func (c *RootCmd) Run(args []string) error {
	if len(args) == 0 {
		fmt.Fprintln(c.out, c.Help())
		return fmt.Errorf("nenhum comando especificado")
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
		fmt.Fprintf(c.errOut, "Erro: comando desconhecido '%s'\n", cmdName)
		fmt.Fprintln(c.errOut, "Use 'assinatura --help' para ver a lista de comandos.")
		slog.Error("comando desconhecido", "comando", cmdName)
		return &UserError{msg: fmt.Sprintf("comando não encontrado: %s", cmdName)}
	}

	// Executa o comando com os argumentos restantes
	return cmd.Run(args[1:])
}

// Help retorna a mensagem de uso geral da aplicação.
func (c *RootCmd) Help() string {
	return `assinatura - CLI para operações de assinatura digital

Usage: assinatura <command> [OPTIONS]

Commands:
  version              Exibe a versão do CLI
  sign / criar         Cria uma assinatura digital para um arquivo
  validate / validar   Valida uma assinatura digital
  start                Inicia o servidor assinador em background
  stop                 Encerra o servidor assinador
  status               Verifica o status do servidor assinador

Global Options:
  --help       Exibe esta mensagem de ajuda
  --version    Exibe informações de versão
  --verbose    Habilita saída detalhada (logs de debug)

Examples:
  assinatura start
  assinatura start --port 9090
  assinatura sign --input documento.pdf
  assinatura criar --input documento.pdf --mode local
  assinatura validate --input documento.pdf --signature documento.sig
  assinatura stop
  assinatura status

Use 'assinatura <command> --help' para mais informações sobre um comando.`
}
