package cmd

import (
    "fmt"
)

func runSign(args []string) error {
    fmt.Println("[TODO] sign command: implement workflow with parameters, JDK detection e assíncrono para assinador.jar")
    if len(args) == 0 {
        fmt.Println("Uso: assinatura sign --input <arquivo> --output <arquivo>")
        return fmt.Errorf("parâmetros insuficientes")
    }
    // placeholder: no parsing até US-02.1/US-01.2
    fmt.Printf("comando sign chamado com args: %v\n", args)
    return nil
}
