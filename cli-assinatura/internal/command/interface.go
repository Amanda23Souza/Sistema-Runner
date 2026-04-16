// Package command, contém a definição e implementação dos comandos CLI.
package command

// Command define a interface que todo comando da CLI deve implementar.
type Command interface {
	// Run executa o comando com os argumentos fornecidos.
	// Retorna um erro se a execução falhar.
	Run(args []string) error

	// Help retorna uma descrição de uso do comando.
	Help() string
}
