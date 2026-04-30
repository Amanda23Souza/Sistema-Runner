// Package version contém informações de versão da aplicação.
package version

import "fmt"

var (
	// Version contém a versão semântica do projeto.
	Version = "v0.1.0"
	// Commit contém o SHA do commit usado na build.
	Commit = "unknown"
	// BuildTime contém o timestamp UTC da build.
	BuildTime = "unknown"
)

// Get retorna a versão da aplicação.
func Get() string {
	return Version
}

// GetFull retorna a versão com informações adicionais (commit e buildTime).
func GetFull() string {
	if Commit == "unknown" && BuildTime == "unknown" {
		return fmt.Sprintf("assinatura %s", Version)
	}

	return fmt.Sprintf("assinatura %s (commit %s, built at %s)", Version, Commit, BuildTime)
}
