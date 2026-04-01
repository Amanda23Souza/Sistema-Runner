package main

import (
    "log"
    "github.com/Amanda23Souza/Sistema-Runner/cli-assinatura/cmd"
)

func main() {
    if err := cmd.Execute(); err != nil {
        log.Fatalf("Erro: %v", err)
    }
}
