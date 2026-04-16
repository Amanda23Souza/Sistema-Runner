package cmd

import (
    "os"
    "testing"
)

func TestExecute_Version(t *testing.T) {
    oldArgs := os.Args
    defer func() { os.Args = oldArgs }()

    os.Args = []string{"assinatura", "version"}
    if err := Execute(); err != nil {
        t.Fatalf("Execute() retornou erro para version: %v", err)
    }
}

func TestExecute_Help(t *testing.T) {
    oldArgs := os.Args
    defer func() { os.Args = oldArgs }()

    os.Args = []string{"assinatura", "help"}
    if err := Execute(); err != nil {
        t.Fatalf("Execute() retornou erro para help: %v", err)
    }
}

func TestExecute_UnknownCommand(t *testing.T) {
    oldArgs := os.Args
    defer func() { os.Args = oldArgs }()

    os.Args = []string{"assinatura", "unknown"}
    if err := Execute(); err == nil {
        t.Fatal("Execute() não retornou erro para comando desconhecido")
    }
}
