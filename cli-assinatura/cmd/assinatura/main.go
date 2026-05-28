package main

import (
	"os"

	"github.com/Amanda23Souza/Sistema-Runner/cli-assinatura/internal/command"
)

func main() {
	root := command.NewRootCmd()
	err := root.Run(os.Args[1:])

	if err != nil {
		if command.IsUserError(err) {
			os.Exit(command.ExitCodeUserError) // 2 = erro do usuário
		}
		os.Exit(command.ExitCodeSystemError) // 1 = erro do sistema
	}
}
