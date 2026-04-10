// Package version contém informações de versão da aplicação.
package version

import "fmt"

// Version contém a versão semântica do projeto.
const Version = "v0.1.0"

// Get retorna a versão da aplicação.
func Get() string {
	return Version
}

// GetFull retorna a versão com informações adicionais (ex: buildTime, commit).
// Por enquanto, retorna apenas a versão semântica.
func GetFull() string {
	return fmt.Sprintf("assinatura %s", Version)
}
