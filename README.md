# Sistema-Runner

Repositório dedicado para desenvolvimento do **projeto Runner**, da disciplina **Implementação e Integração de Software** (UFG).

O Sistema Runner é uma aplicação CLI (Command-Line Interface) em Go que facilita operações de assinatura digital através da linha de comando, sem necessidade de conhecimento aprofundado em configurações Java.

### [📌 Nosso planejamento](https://github.com/Amanda23Souza/Sistema-Runner/blob/main/docs/planejamento/nossoPlanejamento.md)


---

## 📋 Requisitos

- **Go 1.26.1+** (ou superior)
- **JDK** instalado (para operações de assinatura digital com `assinador.jar`)

---

## 📁 Estrutura do Projeto

```
cli-assinatura/
├── cmd/assinatura/
│   └── main.go              ← Ponto de entrada (executável)
├── internal/
│   ├── command/
│   │   ├── interface.go      ← Interface Command (padrão)
│   │   ├── root.go           ← Parser e orquestrador
│   │   ├── version.go        ← Comando de versão
│   │   ├── sign.go           ← Comando de assinatura (stub)
│   │   └── validate.go       ← Comando de validação (stub)
│   ├── version/
│   │   └── version.go        ← Gerenciamento de versão centralizado
│   └── app/                  ← Reservado para futuro
├── go.mod                   ← Definição de módulo Go
└── assinatura              ← Executável compilado
```

---

## 🚀 Como Compilar

```bash
cd cli-assinatura
go build -o assinatura ./cmd/assinatura
```

---

## 💻 Como Usar

### Comando `version`

Exibe a versão do assinatura CLI:

```bash
# Modo padrão
./assinatura version
# Saída: assinatura v0.1.0

# Modo quiet (apenas número da versão)
./assinatura version --quiet
# Saída: v0.1.0

# Modo JSON
./assinatura version --json
# Saída: { "version": "v0.1.0" }
```

### Comando `version` - Ajuda

```bash
./assinatura version --help
```

### Ajuda Geral

```bash
./assinatura --help
# ou
./assinatura
```

---

## 📊 Status de Implementação

### ✅ Fase 1 - Estrutura Base (Completo)

- ✓ Estrutura de pacotes conforme padrão Go (cmd/, internal/)
- ✓ Interface `Command` padronizada para todos os comandos
- ✓ Ponto de entrada minimalista em `cmd/assinatura/main.go`
- ✓ Gerenciamento centralizado de versão

### ✅ Fase 2 - Comando Version (Completo)

- ✓ Comando `version` totalmente funcional
- ✓ Flags: `--quiet`, `--json`, `--help`
- ✓ Exit codes corretos (0 para sucesso, 1 para erro)

### ⏳ Fase 3 - Preparação para Próximos Comandos (Em Progresso)

- ✓ Stubs criados para `sign` e `validate` com estrutura pronta
- ⏳ Implementação de `sign` (criar assinatura digital)
- ⏳ Implementação de `validate` (validar assinatura)
- ⏳ Comandos `start`/`stop` (gerenciador do servidor HTTP)

---

## 📖 Documentação

- **[User Story 01 - Invocar Assinador via CLI](./docs/US-01%20-%20Invocar%20Assinador%20via%20CLI.md)** — Requisitos funcionais e critérios de aceitação
- **[Planejamento do Projeto](./docs/planejamento/nossoPlanejamento.md)** — Visão geral e timeline
- **[Design de Arquitetura](./docs/design.md)** — Padrões e decisões técnicas

---

## 🔧 Próximos Passos

1. **Implementar comando `sign`**
   - Parser de flags: `--input`, `--output`, `--mode`
   - Validação de parâmetros obrigatórios
   - Invocação local (java -jar assinador.jar)
   - Invocação HTTP (fallback se servidor ativo)

2. **Implementar comando `validate`**
   - Parser de flags: `--input`, `--signature`, `--mode`
   - Lógica de validação criptográfica

3. **Implementar comandos `start`/`stop`**
   - Iniciar servidor HTTP do assinador.jar
   - Encerrar processo com detecção de instância ativa
   - Timeout de inatividade

4. **Melhorias Técnicas**
   - Setup de `Makefile` para build multi-plataforma
   - Injeção de versão via `ldflags` em compile-time
   - Logging estruturado com `log/slog`
   - Testes unitários e integração

---
