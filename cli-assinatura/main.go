package main

import (
    "flag"
    "fmt"
    "os"
)

const version = "v0.1.0"

func main() {
    if len(os.Args) < 2 {
        fmt.Println("uso: assinatura <comando> [--args]")
        os.Exit(1)
    }

    cmd := os.Args[1]
    switch cmd {
    case "version":
        fmt.Println(version)
    case "sign":
        // placeholder (ver passo 4)
        fmt.Println("sign comando ainda em construção")
    case "validate":
        fmt.Println("validate comando ainda em construção")
    default:
        fmt.Printf("comando desconhecido: %s\n", cmd)
        os.Exit(1)
    }
}