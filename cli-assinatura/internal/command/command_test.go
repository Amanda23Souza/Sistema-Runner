package command

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestSignCmd_Run_CreatesSignatureFile(t *testing.T) {
	tmp := t.TempDir()
	inputPath := filepath.Join(tmp, "document.txt")
	outputPath := filepath.Join(tmp, "document.sig")
	content := []byte("conteudo de teste")

	if err := os.WriteFile(inputPath, content, 0644); err != nil {
		t.Fatalf("falha ao criar arquivo de entrada: %v", err)
	}

	cmd := NewSignCmd()
	cmd.out = &bytes.Buffer{}

	if err := cmd.Run([]string{"--input", inputPath, "--output", outputPath}); err != nil {
		t.Fatalf("SignCmd.Run retornou erro: %v", err)
	}

	sigBytes, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("falha ao ler arquivo de assinatura: %v", err)
	}

	expected := sha256.Sum256(content)
	if strings.TrimSpace(string(sigBytes)) != hex.EncodeToString(expected[:]) {
		t.Fatalf("assinatura gerada incorreta: esperava %s, obteve %s", hex.EncodeToString(expected[:]), strings.TrimSpace(string(sigBytes)))
	}
}

func TestValidateCmd_Run_ValidSignature(t *testing.T) {
	tmp := t.TempDir()
	inputPath := filepath.Join(tmp, "document.txt")
	sigPath := filepath.Join(tmp, "document.sig")
	content := []byte("conteudo de teste")

	if err := os.WriteFile(inputPath, content, 0644); err != nil {
		t.Fatalf("falha ao criar arquivo de entrada: %v", err)
	}

	hash := sha256.Sum256(content)
	if err := os.WriteFile(sigPath, []byte(hex.EncodeToString(hash[:])), 0644); err != nil {
		t.Fatalf("falha ao criar arquivo de assinatura: %v", err)
	}

	cmd := NewValidateCmd()
	var out bytes.Buffer
	cmd.out = &out

	if err := cmd.Run([]string{"--input", inputPath, "--signature", sigPath}); err != nil {
		t.Fatalf("ValidateCmd.Run retornou erro inesperado: %v", err)
	}

	if !strings.Contains(out.String(), "Status: VÁLIDA") {
		t.Fatalf("mensagem de validação esperada não encontrada: %s", out.String())
	}
}

func TestValidateCmd_Run_InvalidSignature(t *testing.T) {
	tmp := t.TempDir()
	inputPath := filepath.Join(tmp, "document.txt")
	sigPath := filepath.Join(tmp, "document.sig")
	content := []byte("conteudo de teste")

	if err := os.WriteFile(inputPath, content, 0644); err != nil {
		t.Fatalf("falha ao criar arquivo de entrada: %v", err)
	}

	if err := os.WriteFile(sigPath, []byte("deadbeef"), 0644); err != nil {
		t.Fatalf("falha ao criar arquivo de assinatura: %v", err)
	}

	cmd := NewValidateCmd()
	var out bytes.Buffer
	cmd.out = &out

	if err := cmd.Run([]string{"--input", inputPath, "--signature", sigPath}); err == nil {
		t.Fatalf("ValidateCmd.Run deve retornar erro para assinatura inválida")
	}

	if !strings.Contains(out.String(), "Status: INVÁLIDA") {
		t.Fatalf("mensagem de validação inválida esperada não encontrada: %s", out.String())
	}
}
