package command

import (
	"crypto/sha256"
	"encoding/hex"
)

// computeSHA256 calcula o hash SHA-256 de um slice de bytes
// e retorna como string hexadecimal. Usado pelo modo local
// como simulação de assinatura até integração real com assinador.jar.
func computeSHA256(data []byte) string {
	hash := sha256.Sum256(data)
	return hex.EncodeToString(hash[:])
}
