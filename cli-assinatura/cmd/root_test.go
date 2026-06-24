package cmd

import (
	"os"
	"testing"
)

func TestExecute_Version(t *testing.T) {
	old := os.Args
	defer func() { os.Args = old }()

	os.Args = []string{"assinatura", "version"}
	if err := Execute(); err != nil {
		t.Fatalf("Execute() retornou erro para 'version': %v", err)
	}
}

func TestExecute_Help(t *testing.T) {
	old := os.Args
	defer func() { os.Args = old }()

	os.Args = []string{"assinatura", "help"}
	if err := Execute(); err != nil {
		t.Fatalf("Execute() retornou erro para 'help': %v", err)
	}
}

func TestExecute_HelpFlag(t *testing.T) {
	old := os.Args
	defer func() { os.Args = old }()

	os.Args = []string{"assinatura", "--help"}
	if err := Execute(); err != nil {
		t.Fatalf("Execute() retornou erro para '--help': %v", err)
	}
}

func TestExecute_UnknownCommand(t *testing.T) {
	old := os.Args
	defer func() { os.Args = old }()

	os.Args = []string{"assinatura", "comando-inexistente"}
	if err := Execute(); err == nil {
		t.Fatal("Execute() deveria retornar erro para comando desconhecido")
	}
}

func TestExecute_NoArgs(t *testing.T) {
	old := os.Args
	defer func() { os.Args = old }()

	os.Args = []string{"assinatura"}
	if err := Execute(); err == nil {
		t.Fatal("Execute() deveria retornar erro quando nenhum comando é fornecido")
	}
}

func TestExecute_Sign_HelpFlag(t *testing.T) {
	old := os.Args
	defer func() { os.Args = old }()

	os.Args = []string{"assinatura", "sign", "--help"}
	if err := Execute(); err != nil {
		t.Fatalf("Execute() retornou erro para 'sign --help': %v", err)
	}
}

func TestExecute_Validate_HelpFlag(t *testing.T) {
	old := os.Args
	defer func() { os.Args = old }()

	os.Args = []string{"assinatura", "validate", "--help"}
	if err := Execute(); err != nil {
		t.Fatalf("Execute() retornou erro para 'validate --help': %v", err)
	}
}

func TestExecute_Criar_Alias(t *testing.T) {
	old := os.Args
	defer func() { os.Args = old }()

	os.Args = []string{"assinatura", "criar", "--help"}
	if err := Execute(); err != nil {
		t.Fatalf("Execute() retornou erro para alias 'criar --help': %v", err)
	}
}

func TestExecute_Validar_Alias(t *testing.T) {
	old := os.Args
	defer func() { os.Args = old }()

	os.Args = []string{"assinatura", "validar", "--help"}
	if err := Execute(); err != nil {
		t.Fatalf("Execute() retornou erro para alias 'validar --help': %v", err)
	}
}
