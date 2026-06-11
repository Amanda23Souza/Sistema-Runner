// Package e2e contains end-to-end tests that compile the assinatura binary
// and run it as a real subprocess, verifying observable CLI behaviour.
//
// These tests are skipped when running with -short (unit tests only).
// Traceability: US-01 (Invocar Assinador via CLI), critério G1 (pirâmide de testes).
package e2e_test

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

// buildBinary compiles the assinatura binary into a temp dir and returns its path.
func buildBinary(t *testing.T) string {
	t.Helper()

	suffix := ""
	if runtime.GOOS == "windows" {
		suffix = ".exe"
	}

	binary := filepath.Join(t.TempDir(), "assinatura"+suffix)
	moduleRoot := findModuleRoot(t)

	cmd := exec.Command("go", "build", "-o", binary, "./cmd/assinatura")
	cmd.Dir = moduleRoot
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("build falhou: %v\n%s", err, out)
	}

	return binary
}

// findModuleRoot walks up from the test's working directory to find go.mod.
func findModuleRoot(t *testing.T) string {
	t.Helper()

	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("os.Getwd falhou: %v", err)
	}

	for {
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir
		}
		parent := filepath.Dir(dir)
		if parent == dir {
			t.Fatalf("go.mod não encontrado a partir de %s", dir)
		}
		dir = parent
	}
}

// TestCLI_Version verifica que 'assinatura version' retorna saída não vazia.
// US-01 / I2: versão acessível via CLI.
func TestCLI_Version(t *testing.T) {
	if testing.Short() {
		t.Skip("pulando teste e2e em modo -short")
	}

	binary := buildBinary(t)

	out, err := exec.Command(binary, "version").Output()
	if err != nil {
		t.Fatalf("version falhou: %v", err)
	}

	output := strings.TrimSpace(string(out))
	if output == "" {
		t.Fatal("version deve produzir saída não vazia")
	}
	if !strings.Contains(output, "assinatura") {
		t.Fatalf("saída de version deve conter 'assinatura', obteve: %q", output)
	}
}

// TestCLI_Help verifica que '--help' sai com código 0.
// US-01 / I1: --help com exemplos, não lista de flags.
func TestCLI_Help(t *testing.T) {
	if testing.Short() {
		t.Skip("pulando teste e2e em modo -short")
	}

	binary := buildBinary(t)

	cmd := exec.Command(binary, "--help")
	out, err := cmd.Output()
	if err != nil {
		t.Fatalf("--help falhou: %v", err)
	}

	if !strings.Contains(string(out), "sign") || !strings.Contains(string(out), "validate") {
		t.Fatalf("--help deve listar os comandos sign e validate, obteve: %s", out)
	}
}

// TestCLI_SignAndValidate_LocalMode verifica o fluxo sign→validate em modo local.
// US-01: invocar assinador via CLI. US-02: simular assinatura digital.
// E1.1: funciona independente do diretório atual.
func TestCLI_SignAndValidate_LocalMode(t *testing.T) {
	if testing.Short() {
		t.Skip("pulando teste e2e em modo -short")
	}

	binary := buildBinary(t)
	dir := t.TempDir()

	// Cria arquivo de entrada
	inputFile := filepath.Join(dir, "documento.txt")
	if err := os.WriteFile(inputFile, []byte("conteúdo de teste para assinatura"), 0644); err != nil {
		t.Fatalf("falha ao criar arquivo de entrada: %v", err)
	}
	sigFile := inputFile + ".sig"

	// Sign
	cmd := exec.Command(binary, "sign", "--input", inputFile, "--output", sigFile, "--mode", "local")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("sign falhou (exit=%v):\n%s", err, out)
	}

	if _, err := os.Stat(sigFile); err != nil {
		t.Fatalf("arquivo de assinatura não foi criado: %v", err)
	}

	// Validate
	cmd = exec.Command(binary, "validate", "--input", inputFile, "--signature", sigFile, "--mode", "local")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("validate falhou (exit=%v):\n%s", err, out)
	}
	if !strings.Contains(string(out), "VÁLIDA") {
		t.Fatalf("validate deve reportar assinatura VÁLIDA, obteve:\n%s", out)
	}
}

// TestCLI_SignWithSpacesAndAccents verifica E1.2: passagem de argumentos preserva
// espaços, acentos e caracteres especiais nos nomes de arquivos.
func TestCLI_SignWithSpacesAndAccents(t *testing.T) {
	if testing.Short() {
		t.Skip("pulando teste e2e em modo -short")
	}

	binary := buildBinary(t)
	dir := t.TempDir()

	inputFile := filepath.Join(dir, "documento com espaços é ação.txt")
	content := []byte("conteúdo: ção, ã, é, ü, ñ")
	if err := os.WriteFile(inputFile, content, 0644); err != nil {
		t.Fatalf("falha ao criar arquivo com acentos/espaços: %v", err)
	}

	cmd := exec.Command(binary, "sign", "--input", inputFile, "--mode", "local")
	if out, err := cmd.CombinedOutput(); err != nil {
		t.Fatalf("sign com espaços/acentos falhou:\n%s", out)
	}
}

// TestCLI_MissingInput_ExitCode2 verifica E1.3: --input ausente retorna exit code 2
// (UserError) e escreve mensagem no stderr.
func TestCLI_MissingInput_ExitCode2(t *testing.T) {
	if testing.Short() {
		t.Skip("pulando teste e2e em modo -short")
	}

	binary := buildBinary(t)

	cmd := exec.Command(binary, "sign", "--mode", "local")
	out, err := cmd.CombinedOutput()
	if err == nil {
		t.Fatal("esperava exit code != 0 para --input ausente")
	}

	exitErr, ok := err.(*exec.ExitError)
	if !ok {
		t.Fatalf("esperava *exec.ExitError, obteve: %T", err)
	}
	if exitErr.ExitCode() != 2 {
		t.Fatalf("esperava exit code 2 (UserError), obteve %d\noutput: %s", exitErr.ExitCode(), out)
	}
}

// TestCLI_UnknownCommand_ExitCode2 verifica que comando desconhecido retorna exit code 2.
func TestCLI_UnknownCommand_ExitCode2(t *testing.T) {
	if testing.Short() {
		t.Skip("pulando teste e2e em modo -short")
	}

	binary := buildBinary(t)

	cmd := exec.Command(binary, "comando-que-nao-existe")
	err := cmd.Run()
	if err == nil {
		t.Fatal("esperava exit code != 0 para comando desconhecido")
	}

	exitErr, ok := err.(*exec.ExitError)
	if !ok {
		t.Fatalf("esperava *exec.ExitError, obteve: %T", err)
	}
	if exitErr.ExitCode() != 2 {
		t.Fatalf("esperava exit code 2 (UserError) para comando desconhecido, obteve %d", exitErr.ExitCode())
	}
}

// TestCLI_FileNotFound_ExitCode1 verifica E1.3: arquivo de entrada inexistente
// retorna exit code 1 (SystemError), não exit code 2 (UserError).
func TestCLI_FileNotFound_ExitCode1(t *testing.T) {
	if testing.Short() {
		t.Skip("pulando teste e2e em modo -short")
	}

	binary := buildBinary(t)

	cmd := exec.Command(binary, "sign", "--input", "/tmp/arquivo-que-nao-existe-e2e.txt", "--mode", "local")
	err := cmd.Run()
	if err == nil {
		t.Fatal("esperava exit code != 0 para arquivo inexistente")
	}

	exitErr, ok := err.(*exec.ExitError)
	if !ok {
		t.Fatalf("esperava *exec.ExitError, obteve: %T", err)
	}
	if exitErr.ExitCode() != 1 {
		t.Fatalf("arquivo inexistente deve retornar exit code 1 (SystemError), obteve %d", exitErr.ExitCode())
	}
}
