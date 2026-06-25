// Package command contém testes unitários e de integração dos comandos CLI.
// Rastreabilidade:
//   - US-01: Invocar Assinador via CLI (sign, validate, start, stop, status)
//   - US-02: Simular Assinatura Digital com Validação de Parâmetros
package command

import (
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// ─────────────────────────────────────────────────────────
// SignCmd — US-01, US-02
// ─────────────────────────────────────────────────────────

// TestSignCmd_Run_CreatesSignatureFile verifica que o comando sign cria
// o arquivo de assinatura corretamente no modo local.
// US-01: sign --input <arquivo> --mode local
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
	cmd.errOut = &bytes.Buffer{}

	if err := cmd.Run([]string{"--input", inputPath, "--output", outputPath, "--mode", "local"}); err != nil {
		t.Fatalf("SignCmd.Run retornou erro: %v", err)
	}

	sigBytes, err := os.ReadFile(outputPath)
	if err != nil {
		t.Fatalf("falha ao ler arquivo de assinatura: %v", err)
	}

	expected := computeSHA256(content)
	if strings.TrimSpace(string(sigBytes)) != expected {
		t.Fatalf("assinatura gerada incorreta: esperava %s, obteve %s", expected, strings.TrimSpace(string(sigBytes)))
	}
}

// TestSignCmd_Run_MissingInput verifica que --input ausente retorna UserError.
// US-02: validação de parâmetro obrigatório
func TestSignCmd_Run_MissingInput(t *testing.T) {
	cmd := NewSignCmd()
	var errBuf bytes.Buffer
	cmd.out = &bytes.Buffer{}
	cmd.errOut = &errBuf

	err := cmd.Run([]string{"--mode", "local"})
	if err == nil {
		t.Fatal("esperava erro para --input ausente, mas não houve erro")
	}
	if !IsUserError(err) {
		t.Fatalf("esperava UserError (código 2), obteve: %T: %v", err, err)
	}
	if !strings.Contains(errBuf.String(), "--input") {
		t.Fatalf("mensagem de erro deveria mencionar --input, obteve: %q", errBuf.String())
	}
}

// TestSignCmd_Run_InvalidMode verifica que --mode inválido retorna UserError.
func TestSignCmd_Run_InvalidMode(t *testing.T) {
	tmp := t.TempDir()
	inputPath := filepath.Join(tmp, "doc.txt")
	_ = os.WriteFile(inputPath, []byte("test"), 0644)

	cmd := NewSignCmd()
	cmd.out = &bytes.Buffer{}
	cmd.errOut = &bytes.Buffer{}

	err := cmd.Run([]string{"--input", inputPath, "--mode", "invalido"})
	if err == nil {
		t.Fatal("esperava erro para --mode inválido")
	}
	if !IsUserError(err) {
		t.Fatalf("esperava UserError para modo inválido, obteve: %T", err)
	}
}

// TestSignCmd_Run_FileWithSpacesAndAccents verifica E1.2:
// passagem de argumentos preserva espaços e acentos.
func TestSignCmd_Run_FileWithSpacesAndAccents(t *testing.T) {
	tmp := t.TempDir()
	inputPath := filepath.Join(tmp, "documento com espaços é ação.txt")
	outputPath := filepath.Join(tmp, "documento com espaços é ação.sig")
	content := []byte("conteúdo com acentuação: ção, ã, é, ü")

	if err := os.WriteFile(inputPath, content, 0644); err != nil {
		t.Fatalf("falha ao criar arquivo com acentos/espaços: %v", err)
	}

	cmd := NewSignCmd()
	cmd.out = &bytes.Buffer{}
	cmd.errOut = &bytes.Buffer{}

	if err := cmd.Run([]string{"--input", inputPath, "--output", outputPath, "--mode", "local"}); err != nil {
		t.Fatalf("falha ao assinar arquivo com espaços/acentos: %v", err)
	}

	if _, err := os.Stat(outputPath); err != nil {
		t.Fatalf("arquivo de saída não criado: %v", err)
	}
}

// TestSignCmd_Run_InputFileNotFound verifica E1.3: erro de sistema (não UserError)
// quando o arquivo de entrada não existe.
func TestSignCmd_Run_InputFileNotFound(t *testing.T) {
	cmd := NewSignCmd()
	var errBuf bytes.Buffer
	cmd.out = &bytes.Buffer{}
	cmd.errOut = &errBuf

	err := cmd.Run([]string{"--input", "/tmp/arquivo-que-nao-existe-xyz.txt", "--mode", "local"})
	if err == nil {
		t.Fatal("esperava erro para arquivo inexistente")
	}
	if IsUserError(err) {
		t.Fatal("arquivo inexistente deve ser SystemError (exit 1), não UserError (exit 2)")
	}
}

// ─────────────────────────────────────────────────────────
// ValidateCmd — US-01, US-02
// ─────────────────────────────────────────────────────────

// TestValidateCmd_Run_ValidSignature verifica validação de assinatura correta.
// US-01: validate --input <arquivo> --signature <arquivo.sig> --mode local
func TestValidateCmd_Run_ValidSignature(t *testing.T) {
	tmp := t.TempDir()
	inputPath := filepath.Join(tmp, "document.txt")
	sigPath := filepath.Join(tmp, "document.sig")
	content := []byte("conteudo de teste")

	if err := os.WriteFile(inputPath, content, 0644); err != nil {
		t.Fatalf("falha ao criar arquivo de entrada: %v", err)
	}

	hash := computeSHA256(content)
	if err := os.WriteFile(sigPath, []byte(hash), 0644); err != nil {
		t.Fatalf("falha ao criar arquivo de assinatura: %v", err)
	}

	cmd := NewValidateCmd()
	var out bytes.Buffer
	cmd.out = &out
	cmd.errOut = &bytes.Buffer{}

	if err := cmd.Run([]string{"--input", inputPath, "--signature", sigPath, "--mode", "local"}); err != nil {
		t.Fatalf("ValidateCmd.Run retornou erro inesperado: %v", err)
	}

	if !strings.Contains(out.String(), "VÁLIDA") {
		t.Fatalf("mensagem de validação esperada não encontrada: %s", out.String())
	}
}

// TestValidateCmd_Run_InvalidSignature verifica que assinatura inválida
// retorna erro e mensagem clara.
// US-02: cenário negativo — assinatura incorreta
func TestValidateCmd_Run_InvalidSignature(t *testing.T) {
	tmp := t.TempDir()
	inputPath := filepath.Join(tmp, "document.txt")
	sigPath := filepath.Join(tmp, "document.sig")
	content := []byte("conteudo de teste")

	if err := os.WriteFile(inputPath, content, 0644); err != nil {
		t.Fatalf("falha ao criar arquivo de entrada: %v", err)
	}
	if err := os.WriteFile(sigPath, []byte("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa"), 0644); err != nil {
		t.Fatalf("falha ao criar arquivo de assinatura inválida: %v", err)
	}

	cmd := NewValidateCmd()
	var out bytes.Buffer
	cmd.out = &out
	cmd.errOut = &bytes.Buffer{}

	if err := cmd.Run([]string{"--input", inputPath, "--signature", sigPath, "--mode", "local"}); err == nil {
		t.Fatal("ValidateCmd.Run deve retornar erro para assinatura inválida")
	}

	if !strings.Contains(out.String(), "INVÁLIDA") {
		t.Fatalf("mensagem de assinatura inválida esperada não encontrada: %s", out.String())
	}
}

// TestValidateCmd_Run_MissingInput verifica erro para --input ausente.
func TestValidateCmd_Run_MissingInput(t *testing.T) {
	cmd := NewValidateCmd()
	var errBuf bytes.Buffer
	cmd.out = &bytes.Buffer{}
	cmd.errOut = &errBuf

	err := cmd.Run([]string{"--signature", "doc.sig", "--mode", "local"})
	if err == nil {
		t.Fatal("esperava erro para --input ausente")
	}
	if !IsUserError(err) {
		t.Fatalf("esperava UserError, obteve: %T", err)
	}
}

// TestValidateCmd_Run_MissingSignature verifica erro para --signature ausente.
func TestValidateCmd_Run_MissingSignature(t *testing.T) {
	cmd := NewValidateCmd()
	var errBuf bytes.Buffer
	cmd.out = &bytes.Buffer{}
	cmd.errOut = &errBuf

	err := cmd.Run([]string{"--input", "doc.txt", "--mode", "local"})
	if err == nil {
		t.Fatal("esperava erro para --signature ausente")
	}
	if !IsUserError(err) {
		t.Fatalf("esperava UserError, obteve: %T", err)
	}
}

// TestValidateCmd_Run_SignatureFileNotFound verifica erro de sistema
// quando o arquivo de assinatura não existe.
func TestValidateCmd_Run_SignatureFileNotFound(t *testing.T) {
	tmp := t.TempDir()
	inputPath := filepath.Join(tmp, "doc.txt")
	_ = os.WriteFile(inputPath, []byte("test"), 0644)

	cmd := NewValidateCmd()
	cmd.out = &bytes.Buffer{}
	cmd.errOut = &bytes.Buffer{}

	err := cmd.Run([]string{"--input", inputPath, "--signature", "/tmp/nao-existe.sig", "--mode", "local"})
	if err == nil {
		t.Fatal("esperava erro para arquivo de assinatura inexistente")
	}
	if IsUserError(err) {
		t.Fatal("arquivo inexistente deve ser SystemError, não UserError")
	}
}

// ─────────────────────────────────────────────────────────
// StartCmd
// ─────────────────────────────────────────────────────────

// TestStartCmd_ParseTimeoutFlag verifica se a flag --timeout é parseada corretamente.
func TestStartCmd_ParseTimeoutFlag(t *testing.T) {
	cmd := NewStartCmd()
	cmd.out = &bytes.Buffer{}
	cmd.errOut = &bytes.Buffer{}

	// Passando a flag --timeout 15, não queremos executar o Run completo (que tentaria iniciar o java),
	// então usaremos a validação inicial do Parse via Run. O Run vai falhar porque não acha o JAR na nossa temp.
	// Mas o parse terá ocorrido! Então vamos apenas verificar se cmd.timeout foi setado corretamente.
	// Outra forma é verificar que não dá erro de sintaxe.
	
	// Para não travar no exec.LookPath ou Run real, vamos apenas mockar os args se possível, ou testar o erro
	err := cmd.Run([]string{"--timeout", "15", "--jar", "nonexistent.jar"})
	
	// Esperamos um erro porque "nonexistent.jar" não existe ou Java não está no PATH, mas o timeout deve ser 15
	if cmd.timeout != 15 {
		t.Fatalf("esperava timeout 15, obteve: %d (erro retornado: %v)", cmd.timeout, err)
	}
}

// ─────────────────────────────────────────────────────────
// RootCmd — help e roteamento
// ─────────────────────────────────────────────────────────

// TestRootCmd_UnknownCommand verifica que comando desconhecido retorna UserError.
func TestRootCmd_UnknownCommand(t *testing.T) {
	root := NewRootCmd()
	var out, errOut bytes.Buffer
	root.out = &out
	root.errOut = &errOut

	err := root.Run([]string{"comando-inexistente"})
	if err == nil {
		t.Fatal("esperava erro para comando desconhecido")
	}
	if !IsUserError(err) {
		t.Fatalf("esperava UserError para comando desconhecido, obteve: %T", err)
	}
	if !strings.Contains(errOut.String(), "comando desconhecido") {
		t.Fatalf("mensagem de erro esperada não encontrada no stderr: %q", errOut.String())
	}
}

// TestRootCmd_Help verifica que --help não retorna erro.
func TestRootCmd_Help(t *testing.T) {
	root := NewRootCmd()
	var out bytes.Buffer
	root.out = &out
	root.errOut = &bytes.Buffer{}

	if err := root.Run([]string{"--help"}); err != nil {
		t.Fatalf("--help não deve retornar erro: %v", err)
	}
	if !strings.Contains(out.String(), "assinatura") {
		t.Fatalf("help deveria conter 'assinatura', obteve: %q", out.String())
	}
}

// ─────────────────────────────────────────────────────────
// VersionCmd
// ─────────────────────────────────────────────────────────

// TestVersionCmd_Run_ReturnsVersion verifica que version retorna sem erro.
func TestVersionCmd_Run_ReturnsVersion(t *testing.T) {
	cmd := NewVersionCmd()
	var out bytes.Buffer
	cmd.out = &out

	if err := cmd.Run([]string{}); err != nil {
		t.Fatalf("version não deve retornar erro: %v", err)
	}
	if out.String() == "" {
		t.Fatal("version deve produzir saída não vazia")
	}
}

// TestVersionCmd_Run_Quiet verifica que --quiet retorna apenas o número de versão.
func TestVersionCmd_Run_Quiet(t *testing.T) {
	cmd := NewVersionCmd()
	var out bytes.Buffer
	cmd.out = &out

	if err := cmd.Run([]string{"--quiet"}); err != nil {
		t.Fatalf("version --quiet não deve retornar erro: %v", err)
	}
	output := strings.TrimSpace(out.String())
	if strings.Contains(output, " ") {
		t.Fatalf("--quiet deve retornar apenas o número de versão sem espaços, obteve: %q", output)
	}
}

// ─────────────────────────────────────────────────────────
// JDK Provisioning
// ─────────────────────────────────────────────────────────

// TestLocalJDKBin_ReturnsValidPath verifica que localJDKBin retorna caminho terminando em java.
func TestLocalJDKBin_ReturnsValidPath(t *testing.T) {
	bin := localJDKBin()
	if bin == "" {
		t.Fatal("localJDKBin() retornou string vazia")
	}
	if !strings.HasSuffix(bin, "java") && !strings.HasSuffix(bin, "java.exe") {
		t.Fatalf("localJDKBin() deve terminar com 'java' ou 'java.exe', obteve: %q", bin)
	}
}

// TestLocalJDKDir_ContainsAssinadorDir verifica que o diretório de cache contém ".assinador".
func TestLocalJDKDir_ContainsAssinadorDir(t *testing.T) {
	dir := localJDKDir()
	if !strings.Contains(dir, ".assinador") {
		t.Fatalf("localJDKDir() deve conter '.assinador', obteve: %q", dir)
	}
}

// TestStartCmd_ResolveJava_UsesPathIfAvailable verifica que resolveJava usa java do PATH quando disponível.
func TestStartCmd_ResolveJava_UsesPathIfAvailable(t *testing.T) {
	if _, err := exec.LookPath("java"); err != nil {
		t.Skip("java não encontrado no PATH — pulando teste de resolução")
	}

	cmd := NewStartCmd()
	cmd.out = &bytes.Buffer{}
	cmd.errOut = &bytes.Buffer{}

	path, err := cmd.resolveJava()
	if err != nil {
		t.Fatalf("resolveJava() retornou erro quando java está no PATH: %v", err)
	}
	if path == "" {
		t.Fatal("resolveJava() retornou caminho vazio")
	}
}

// ─────────────────────────────────────────────────────────
// SignCmd — modo HTTP (servidor mockado)
// ─────────────────────────────────────────────────────────

// TestSignCmd_Run_HTTPMode_ConnRefused verifica que modo http retorna erro de sistema
// (não UserError) quando o servidor não está rodando.
func TestSignCmd_Run_HTTPMode_ConnRefused(t *testing.T) {
	tmp := t.TempDir()
	inputPath := filepath.Join(tmp, "doc.txt")
	_ = os.WriteFile(inputPath, []byte("conteúdo para assinar"), 0644)

	cmd := NewSignCmd()
	cmd.out = &bytes.Buffer{}
	var errBuf bytes.Buffer
	cmd.errOut = &errBuf

	err := cmd.Run([]string{"--input", inputPath, "--mode", "http", "--port", "19999"})
	if err == nil {
		t.Fatal("esperava erro de conexão recusada no modo http")
	}
	if IsUserError(err) {
		t.Fatalf("conexão recusada deve ser SystemError, não UserError: %v", err)
	}
}

// TestValidateCmd_Run_HTTPMode_ConnRefused verifica que modo http retorna SystemError
// quando o servidor não está rodando.
func TestValidateCmd_Run_HTTPMode_ConnRefused(t *testing.T) {
	tmp := t.TempDir()
	inputPath := filepath.Join(tmp, "doc.txt")
	sigPath := filepath.Join(tmp, "doc.sig")
	_ = os.WriteFile(inputPath, []byte("conteúdo"), 0644)
	_ = os.WriteFile(sigPath, []byte("assinatura-fake"), 0644)

	cmd := NewValidateCmd()
	cmd.out = &bytes.Buffer{}
	cmd.errOut = &bytes.Buffer{}

	err := cmd.Run([]string{"--input", inputPath, "--signature", sigPath, "--mode", "http", "--port", "19999"})
	if err == nil {
		t.Fatal("esperava erro de conexão recusada no modo http")
	}
	if IsUserError(err) {
		t.Fatalf("conexão recusada deve ser SystemError, não UserError: %v", err)
	}
}

// TestSignCmd_Run_HTTPMode_ReadsFileContent verifica que o modo http tenta ler o arquivo
// (deve falhar com SystemError se o arquivo não existir, não UserError).
func TestSignCmd_Run_HTTPMode_FileNotFound(t *testing.T) {
	cmd := NewSignCmd()
	cmd.out = &bytes.Buffer{}
	cmd.errOut = &bytes.Buffer{}

	err := cmd.Run([]string{"--input", "/tmp/arquivo-inexistente-xyz.txt", "--mode", "http", "--port", "19999"})
	if err == nil {
		t.Fatal("esperava erro para arquivo inexistente em modo http")
	}
	if IsUserError(err) {
		t.Fatalf("arquivo inexistente deve ser SystemError, não UserError: %v", err)
	}
}
