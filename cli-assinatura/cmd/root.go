// Package cmd fornece o dispatcher principal da CLI.
// Segue a estrutura esperada pelo professor: switch/case em Execute()
// delegando para as implementações em internal/command.
package cmd

import (
	"fmt"
	"os"

	"github.com/Amanda23Souza/Sistema-Runner/cli-assinatura/internal/command"
)

// Version pode ser sobrescrita em tempo de compilação:
// go build -ldflags "-X github.com/Amanda23Souza/Sistema-Runner/cli-assinatura/cmd.Version=v1.0.0"
var Version = "v0.1.0"

// Execute processa os argumentos da CLI e despacha para o comando correto.
func Execute() error {
	args := os.Args[1:]
	if len(args) == 0 {
		printHelp()
		return fmt.Errorf("nenhum comando fornecido")
	}

	switch args[0] {
	case "version":
		return command.NewVersionCmd().Run([]string{})
	case "sign", "criar":
		return runSign(args[1:])
	case "validate", "validar":
		return runValidate(args[1:])
	case "start":
		return command.NewStartCmd().Run(args[1:])
	case "stop":
		return command.NewStopCmd().Run(args[1:])
	case "status":
		return command.NewStatusCmd().Run(args[1:])
	case "help", "--help", "-h":
		printHelp()
		return nil
	default:
		return fmt.Errorf("comando desconhecido: %s", args[0])
	}
}

func printHelp() {
	fmt.Println("Sistema Runner - CLI de Assinatura Digital")
	fmt.Println("Uso: assinatura <comando> [opções]")
	fmt.Println()
	fmt.Println("Comandos:")
	fmt.Println("  version             Exibe a versão do CLI")
	fmt.Println("  sign / criar        Cria uma assinatura digital para um arquivo")
	fmt.Println("  validate / validar  Valida uma assinatura digital")
	fmt.Println("  start               Inicia o servidor assinador em background")
	fmt.Println("  stop                Encerra o servidor assinador")
	fmt.Println("  status              Verifica o status do servidor assinador")
	fmt.Println("  help / --help       Mostra esta ajuda")
}
