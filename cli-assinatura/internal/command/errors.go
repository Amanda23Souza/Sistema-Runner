// Package command define e orquestra os comandos da aplicação CLI.
package command

// ExitCodeUserError é o código de saída para erros causados pelo usuário
// (parâmetros inválidos, arquivo não encontrado pelo usuário, etc.).
const ExitCodeUserError = 2

// ExitCodeSystemError é o código de saída para erros do sistema
// (falha de I/O, rede, processo, etc.).
const ExitCodeSystemError = 1

// UserError representa um erro causado por uso incorreto pelo usuário.
// Quando propagado até main(), resulta em os.Exit(ExitCodeUserError).
type UserError struct {
	msg string
}

func (e *UserError) Error() string { return e.msg }

// IsUserError retorna true se o erro é um UserError.
func IsUserError(err error) bool {
	_, ok := err.(*UserError)
	return ok
}
