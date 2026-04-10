package main

import (
	"os"

	"github.com/Amanda23Souza/Sistema-Runner/cli-assinatura/internal/command"
)

func main() {
	root := command.NewRootCmd()
	err := root.Run(os.Args[1:])

	if err != nil {
		os.Exit(1)
	}
}
