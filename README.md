# Sistema-Runner

Repositório dedicado ao desenvolvimento do **projeto Runner**, da disciplina **Implementação e Integração de Software** (UFG).

O Sistema Runner é uma aplicação CLI (Command-Line Interface) em Go que facilita operações de assinatura digital através da linha de comando, delegando a lógica criptográfica ao `assinador.jar` — um servidor HTTP leve escrito em Java.

**Especificação de referência:**
[`kyriosdata/runner @ d3f1a9c`](https://github.com/kyriosdata/runner/blob/d3f1a9c/docs/runner.md) *(link fixo — não usa `main` para evitar deriva)*

---

## 📋 Requisitos

| Ferramenta | Versão Mínima | Verificação |
|------------|:---:|---|
| Go | 1.26.1+ | `go version` |
| JDK | 21+ | `java --version` |
| Maven | 3.8+ | `mvn --version` *(apenas para compilar o JAR)* |

---

## 📁 Estrutura do Projeto

```
.
├── cli-assinatura/              ← Subprojeto CLI (Go)
│   ├── cmd/assinatura/
│   │   └── main.go              ← Ponto de entrada (exit codes distintos: 1=sistema, 2=usuário)
│   ├── internal/
│   │   ├── command/
│   │   │   ├── root.go          ← Orquestrador de subcomandos
│   │   │   ├── sign.go          ← Comando sign/criar (modo http padrão)
│   │   │   ├── validate.go      ← Comando validate/validar
│   │   │   ├── start.go         ← Comando start (health check idempotente)
│   │   │   ├── stop.go          ← Comando stop (/shutdown + PID fallback) + status
│   │   │   ├── version.go       ← Comando version (tag + SHA curto)
│   │   │   ├── http.go          ← Cliente HTTP com timeout e tratamento de erros
│   │   │   ├── errors.go        ← UserError vs SystemError
│   │   │   └── util.go          ← Utilitários (SHA-256)
│   │   └── version/
│   │       └── version.go       ← Metadados de versão (injetados via ldflags)
│   ├── test/e2e/
│   │   └── cli_test.go          ← Testes e2e: compila o binário real e valida comportamento
│   └── go.mod
├── docs/aulas/projetos/assinador-java/  ← Subprojeto Assinador (Java)
│   ├── src/main/java/com/runner/assinador/
│   │   ├── App.java             ← Servidor Javalin (auto-shutdown, /shutdown endpoint)
│   │   ├── SignatureController.java
│   │   ├── SignatureService.java
│   │   └── FakeSignatureService.java
│   └── pom.xml
├── docs/
│   ├── adr/                     ← Architecture Decision Records
│   │   ├── 001-escolha-go-para-cli.md
│   │   ├── 002-modo-servidor-http-padrao.md
│   │   ├── 003-parser-cli-stdlib-flag.md
│   │   └── 004-simulador-pkcs11.md
│   ├── api-contract.md          ← Contrato formal CLI↔JAR (endpoints, JSON, erros)
│   ├── conformidade-criterios.md ← Relatório de conformidade com critérios
│   ├── design.md                ← Diagramas C4 (contexto e contêineres)
│   └── planejamento/
│       └── nossoPlanejamento.md
├── .gitignore
├── .gitattributes
└── LICENSE                      ← Apache 2.0
```

---

## 🚀 Como Compilar

### CLI (Go)
```bash
cd cli-assinatura
go build -o assinatura ./cmd/assinatura
```

### Assinador JAR (Java)
```bash
cd docs/aulas/projetos/assinador-java
mvn clean package
# Gera: target/assinador-java-1.0.0-SNAPSHOT-jar-with-dependencies.jar
```

---

## 💻 Como Usar

O modo padrão é **HTTP** — o CLI se comunica com o servidor assinador em background.

### Fluxo recomendado

```bash
# 1. Iniciar o servidor assinador (verifica saúde, não duplica instâncias)
./assinatura start

# 2. Assinar um arquivo
./assinatura sign --input documento.pdf

# 3. Validar uma assinatura
./assinatura validate --input documento.pdf --signature documento.pdf.sig

# 4. Verificar status do servidor
./assinatura status

# 5. Encerrar o servidor
./assinatura stop
```

### Modo local (sem servidor)

```bash
./assinatura sign --input documento.pdf --mode local
./assinatura validate --input documento.pdf --signature documento.sig --mode local
```

### Opções globais

```bash
./assinatura --help      # Ajuda com exemplos
./assinatura --version   # Exibe versão + commit + buildtime
./assinatura version --json   # Versão em JSON estruturado
./assinatura version --quiet  # Apenas o número de versão
```

### Servidor com porta personalizada

```bash
./assinatura start --port 9090
./assinatura sign --input doc.pdf --port 9090
./assinatura stop --port 9090
```

### Servidor com auto-shutdown configurável

```bash
# O JAR pode ser iniciado diretamente com timeout de inatividade em segundos
java -jar assinador-java-1.0.0-SNAPSHOT-jar-with-dependencies.jar 8080 --inactivity-timeout 600
```

---

## 🧪 Como Executar os Testes

### Testes unitários da CLI (Go)

```bash
cd cli-assinatura
go test ./...
```

> Os testes e2e em `test/e2e/` compilam o binário real e levam alguns segundos a mais.
> Para rodar apenas testes unitários (rápidos), use `-short`:
> ```bash
> go test -short ./...
> ```

### Com cobertura de código

```bash
cd cli-assinatura
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out   # abre relatório no browser
```

### Testes com race detector

```bash
cd cli-assinatura
go test -race ./...
```

### Testes do Assinador Java

```bash
cd docs/aulas/projetos/assinador-java
mvn verify          # compila, testa e verifica estilo (checkstyle)
mvn test            # apenas testes, sem checkstyle
```

---

## 🤝 Como Contribuir

1. **Fork** o repositório e crie uma branch: `git checkout -b feat/minha-feature`
2. **Commits** no padrão [Conventional Commits](https://www.conventionalcommits.org/): `feat(cli): adicionar comando start`
3. **Testes**: garanta que `go test ./...` passa sem erros
4. **Lint**: execute `golangci-lint run` antes de abrir PR
5. **PR**: abra um Pull Request com descrição clara; aguarde revisão do outro membro

### Convenções de código

- Go: `gofmt`, `go vet`, `golangci-lint`
- Java: estilo padrão com `maven-checkstyle-plugin`
- Commits: `type(scope): descrição` — tipos: `feat`, `fix`, `docs`, `test`, `refactor`, `ci`, `chore`
- Mensagens de erro CLI: sempre em português, com **o quê**, **por quê** e **como resolver**

### Códigos de saída

| Código | Significado |
|:---:|---|
| `0` | Sucesso |
| `1` | Erro do sistema (I/O, rede, processo) |
| `2` | Erro do usuário (parâmetro inválido, arquivo não encontrado pelo usuário) |

---

## 📁 Estrutura do Pacote `cmd/`

O pacote `cmd/` expõe o dispatcher `Execute()` seguindo a estrutura de referência do professor:

```
cli-assinatura/
├── cmd/
│   ├── root.go          ← Execute() — dispatcher principal (switch/case)
│   ├── sign.go          ← runSign() — delega para internal/command
│   ├── validate.go      ← runValidate() — delega para internal/command
│   ├── root_test.go     ← testes do dispatcher Execute()
│   └── assinatura/
│       └── main.go      ← ponto de entrada do binário
└── internal/
    └── command/
        ├── jdk.go        ← provisionamento automático JDK 21 (Adoptium)
        └── ...
```

---

## 🔧 Provisionamento Automático do JDK

O comando `start` detecta e provisiona o JDK automaticamente em três etapas:

1. **PATH**: usa `java` do PATH do sistema se disponível
2. **Cache local**: usa `~/.assinador/jdk/bin/java` se já provisionado
3. **Download automático**: baixa JDK 21 da [Adoptium](https://adoptium.net) e extrai em `~/.assinador/jdk/`

O download é silencioso na primeira execução e as versões subsequentes usam o cache:

```bash
# Primeira execução (sem java no PATH): faz download do JDK 21 automaticamente
./assinatura start

# Execuções seguintes: usa o cache em ~/.assinador/jdk/
./assinatura start
```

---

## 🔑 Integração PKCS#11 (Simulada)

O fluxo de assinatura integra CLI → servidor JAR → `FakeSignatureService` (simulação PKCS#11):

```
assinatura sign --input doc.pdf
       │
       ▼ lê conteúdo do arquivo
       │
       ▼ POST /sign  {"content": "<conteúdo>"}
  assinador.jar (porta 8080)
       │
       ▼ FakeSignatureService.sign()
       │  retorna MOCKED_SIGNATURE_BASE64_==
       │
       ▼ salva em doc.pdf.sig
```

```
assinatura validate --input doc.pdf --signature doc.pdf.sig
       │
       ▼ lê conteúdo + assinatura
       │
       ▼ POST /validate {"content": "...", "signature": "MOCKED_SIGNATURE_BASE64_=="}
  assinador.jar
       │
       ▼ valid: true/false
```

---

## 📊 Status de Implementação (Sprint 4 — Release Final)

### ✅ Fase 1 — Estrutura Base
- ✓ Estrutura de pacotes Go (`cmd/`, `internal/`)
- ✓ Interface `Command` padronizada
- ✓ Ponto de entrada com exit codes distintos (0/1/2)
- ✓ Versão gerenciada centralmente (`internal/version/`)

### ✅ Fase 2 — Comando Version
- ✓ `version --quiet`, `--json`, `--help`
- ✓ Exibe: tag + SHA curto + buildtime (injetados via ldflags)

### ✅ Fase 3 — Backend Java Base
- ✓ Servidor HTTP Javalin com `/sign`, `/validate`, `/health`, `/shutdown`
- ✓ Auto-shutdown por inatividade (padrão: 5 min, configurável via `--inactivity-timeout`)
- ✓ Testes de integração `SignatureControllerTest` + `FakeSignatureServiceTest`

### ✅ Fase 4 — CLI Lifecycle (Start/Stop/Status)
- ✓ `start` idempotente (health check real antes de iniciar)
- ✓ `stop` com shutdown graceful via `/shutdown` + fallback PID (SIGINT/taskkill)
- ✓ `status` com verificação de prontidão real
- ✓ Compilação automática do JAR com Maven (`mvn package`) se não encontrado

### ✅ Fase 5 — Parsing de Comandos (Sprint 3/4)
- ✓ Pacote `cmd/` com dispatcher `Execute()` (switch/case, conforme estrutura do professor)
- ✓ `cmd/root.go`, `cmd/sign.go`, `cmd/validate.go`, `cmd/root_test.go`
- ✓ Aliases: `sign`/`criar`, `validate`/`validar`

### ✅ Fase 6 — Invocação do assinador.jar via CLI
- ✓ Modo HTTP padrão: lê conteúdo do arquivo e envia ao `/sign` com JSON válido
- ✓ Modo local: assinatura SHA-256 offline (sem servidor)
- ✓ `validate` lê arquivo + assinatura e envia ao `/validate`

### ✅ Fase 7 — Provisionamento Automático do JDK
- ✓ Detecção em PATH → cache local (`~/.assinador/jdk/`) → download automático
- ✓ Download do JDK 21 da Adoptium (Linux/Mac: `.tar.gz`, Windows: `.zip`)
- ✓ Extração nativa em Go (sem dependências externas)

### ✅ Fase 8 — Elaboração de Testes
- ✓ Testes unitários: `internal/command/command_test.go` (sign, validate, start, version, root)
- ✓ Testes do dispatcher: `cmd/root_test.go`
- ✓ Testes de JDK provisioning: `localJDKBin`, `localJDKDir`, `resolveJava`
- ✓ Testes de modo HTTP: verifica comportamento correto com servidor offline
- ✓ Testes e2e: `test/e2e/cli_test.go` (compila binário real)
- ✓ CI multi-plataforma: Ubuntu + Windows

### ✅ Fase 9 — Refinamento
- ✓ Bug corrigido: modo HTTP enviava caminho do arquivo em vez do conteúdo
- ✓ JSON encoding correto via `json.Marshal` (não formatação manual com `%s`)
- ✓ Separação clara `UserError` (exit 2) vs `SystemError` (exit 1)

### ✅ Fase 10 — Documentação Final
- ✓ README com fluxo completo, diagrama de integração, referência de comandos
- ✓ ADRs: Go para CLI, modo HTTP padrão, stdlib flag, simulador PKCS#11
- ✓ API contract: `docs/api-contract.md`

---

## 📖 Documentação

- [Contrato de API CLI↔JAR](./docs/api-contract.md)
- [ADR-001 — Escolha de Go para a CLI](./docs/adr/001-escolha-go-para-cli.md)
- [ADR-002 — Modo Servidor HTTP como Padrão](./docs/adr/002-modo-servidor-http-padrao.md)
- [ADR-003 — Parser de CLI: stdlib flag](./docs/adr/003-parser-cli-stdlib-flag.md)
- [ADR-004 — Simulador PKCS#11](./docs/adr/004-simulador-pkcs11.md)
- [Design de Arquitetura (C4)](./docs/design.md)
- [Planejamento do Projeto](./docs/planejamento/nossoPlanejamento.md)
- [Relatório de Conformidade](./docs/conformidade-criterios.md)

---

## 📜 Licença

Apache 2.0 — veja [LICENSE](./LICENSE).
