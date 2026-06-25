package cmd

import "github.com/Amanda23Souza/Sistema-Runner/cli-assinatura/internal/command"

// runSign despacha o comando sign para a implementação em internal/command.
func runSign(args []string) error {
	return command.NewSignCmd().Run(args)
}
