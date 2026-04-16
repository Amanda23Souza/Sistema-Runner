package cmd

import (
    "errors"
    "fmt"
    "os"
)

// Version variable can be overridden at build time with -ldflags "-X github.com/Amanda23Souza/Sistema-Runner/cli-assinatura/cmd.Version=vX.Y.Z".
var Version = "v0.1.0"

// Execute processes CLI arguments and dispatch commands.
func Execute() error {
    if len(os.Args) < 2 {
        printHelp()
        return errors.New("nenhum comando fornecido")
    }

    cmd := os.Args[1]
    switch cmd {
    case "version":
        return runVersion()
    case "sign":
        return runSign(os.Args[2:])
    case "validate":
        return runValidate(os.Args[2:])
    case "help", "--help", "-h":
        printHelp()
        return nil
    default:
        return fmt.Errorf("comando desconhecido: %s", cmd)
    }
}

func runVersion() error {
    fmt.Println(Version)
    return nil
}

func printHelp() {
    fmt.Println("Sistema Runner - CLI de Assinatura")
    fmt.Println("Uso: assinatura <comando> [opções]")
    fmt.Println("Comandos:")
    fmt.Println("  version           Exibe a versão do CLI")
    fmt.Println("  sign              Envia dados para assinar via assinador.jar")
    fmt.Println("  validate          Valida assinatura via assinador.jar")
    fmt.Println("  help              Mostra esta ajuda")
}
