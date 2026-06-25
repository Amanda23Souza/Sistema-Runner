package cmd

import "github.com/Amanda23Souza/Sistema-Runner/cli-assinatura/internal/command"

// runValidate despacha o comando validate para a implementação em internal/command.
func runValidate(args []string) error {
	return command.NewValidateCmd().Run(args)
}
