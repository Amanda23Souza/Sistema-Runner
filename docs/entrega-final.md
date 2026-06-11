# Entrega Final — Sistema Runner

**Disciplina:** Implementação e Integração de Software  
**Instituição:** Universidade Federal de Goiás (UFG)  
**Equipe:** Amanda Sousa · Marcello Ronald  
**Data de entrega:** 2026-06-10  
**Repositório:** https://github.com/Amanda23Souza/Sistema-Runner  
**Especificação de referência:** [`kyriosdata/runner @ d3f1a9c`](https://github.com/kyriosdata/runner/blob/d3f1a9c/docs/runner.md)

---

## 1. Visão Geral

O **Sistema Runner** é uma aplicação de linha de comando (CLI) que permite assinar e validar digitalmente arquivos, delegando a criptografia a um servidor HTTP leve em background.

O sistema é composto por **dois binários**:

| Binário | Linguagem | Responsabilidade |
|---|---|---|
| `assinatura` (Go) | Go 1.26 | CLI que o usuário invoca; gerencia o ciclo de vida do servidor |
| `assinador.jar` (Java) | Java 21 + Javalin | Servidor HTTP leve; executa as operações de assinatura digital |

A comunicação é feita exclusivamente via HTTP local (`localhost:8080` por padrão). O servidor é iniciado e encerrado pela CLI, e se desliga automaticamente por inatividade.

### Fluxo principal

```
usuário → assinatura start          → inicia assinador.jar em background
usuário → assinatura sign  --input  → POST /sign   → JAR assina, retorna .sig
usuário → assinatura validate       → POST /validate → JAR valida, retorna OK/NOK
usuário → assinatura stop           → POST /shutdown → JAR encerra graciosamente
```

---

## 2. Estrutura do Repositório

```
Sistema-Runner/
├── cli-assinatura/                    ← Módulo Go (binário assinatura)
│   ├── cmd/assinatura/main.go         ← Entrada: exit codes 0/1/2
│   ├── internal/command/
│   │   ├── root.go                    ← Orquestrador de subcomandos
│   │   ├── sign.go                    ← sign --input --output --mode
│   │   ├── validate.go                ← validate --input --signature --mode
│   │   ├── start.go                   ← start: idempotente, health check real
│   │   ├── stop.go                    ← stop: /shutdown + fallback PID
│   │   ├── version.go                 ← version --quiet --json
│   │   ├── http.go                    ← Cliente HTTP: timeout, retry, erros
│   │   ├── errors.go                  ← UserError(2) vs SystemError(1)
│   │   └── util.go                    ← SHA-256 local (modo local)
│   ├── internal/version/version.go   ← Metadados injetados via -ldflags
│   ├── test/e2e/cli_test.go           ← 7 testes e2e subprocess
│   └── go.mod
│
├── docs/aulas/projetos/assinador-java/  ← Módulo Maven (assinador.jar)
│   ├── src/main/java/com/runner/assinador/
│   │   ├── App.java                   ← Servidor Javalin; auto-shutdown
│   │   ├── SignatureController.java   ← Handlers /sign, /validate
│   │   ├── SignatureService.java      ← Interface de serviço
│   │   └── FakeSignatureService.java  ← Simulador PKCS#11
│   ├── src/test/java/com/runner/assinador/
│   │   ├── SignatureControllerTest.java  ← 5 testes HTTP reais
│   │   ├── FakeSignatureServiceTest.java ← 4 testes unitários
│   │   └── InactivityTimerTest.java      ← 2 testes de auto-shutdown
│   ├── checkstyle.xml
│   └── pom.xml
│
├── docs/
│   ├── adr/                           ← 4 Architecture Decision Records
│   ├── api-contract.md                ← Contrato formal CLI↔JAR
│   ├── conformidade-criterios.md      ← Relatório de conformidade (~96%)
│   ├── design.md                      ← Diagramas C4
│   └── entrega-final.md               ← Este documento
│
├── .github/workflows/build.yml        ← Pipeline CI/CD completo
├── .gitattributes                     ← LF para .go/.java/.md; CRLF para .bat
└── LICENSE                            ← Apache 2.0
```

---

## 3. Como Compilar e Executar

### Pré-requisitos

| Ferramenta | Versão |
|---|---|
| Go | ≥ 1.26 |
| JDK | ≥ 21 |
| Maven | ≥ 3.8 |

### Compilar

```bash
# CLI Go
cd cli-assinatura
go build -o assinatura ./cmd/assinatura

# Servidor Java (fat JAR)
cd docs/aulas/projetos/assinador-java
mvn clean package
# Gera: target/assinador-java-1.0.0-SNAPSHOT-jar-with-dependencies.jar
```

### Executar

```bash
# 1. Iniciar o servidor (verifica saúde, idempotente)
./assinatura start

# 2. Assinar um arquivo
./assinatura sign --input documento.pdf

# 3. Validar a assinatura
./assinatura validate --input documento.pdf --signature documento.pdf.sig

# 4. Verificar status do servidor
./assinatura status

# 5. Encerrar o servidor
./assinatura stop

# Modo local (sem servidor)
./assinatura sign --input arquivo.txt --mode local
```

### Opções adicionais

```bash
./assinatura --help                # Ajuda com exemplos
./assinatura version --json        # Versão em JSON
./assinatura start --port 9090     # Porta personalizada
./assinatura sign --verbose        # Logs detalhados
```

---

## 4. Como Executar os Testes

### CLI Go

```bash
cd cli-assinatura

# Testes unitários com cobertura e race detector
go test -v -race -short -coverprofile=coverage.out ./...
go tool cover -html=coverage.out      # Relatório HTML

# Testes e2e (compilam o binário real, levam ~30s)
go test -v -timeout 120s ./test/e2e/
```

### Servidor Java

```bash
cd docs/aulas/projetos/assinador-java

mvn verify          # Compila + testes unitários + testes de integração + checkstyle
mvn test            # Apenas testes (sem checkstyle)
```

---

## 5. Pipeline CI/CD

O arquivo `.github/workflows/build.yml` orquestra toda a automação:

```
Push / PR para main
        │
        ├── test-java (ubuntu) ──── mvn clean verify (checkstyle + testes Java)
        ├── lint (ubuntu) ────────── golangci-lint (Go)
        └── test (ubuntu + windows) go test -race -short ./...
                                     go test ./test/e2e/
                │
                ▼ (todos aprovados)
            build (ubuntu, cross-compile)
                ├── linux/amd64
                ├── linux/arm64
                ├── windows/amd64  (.exe)
                ├── darwin/amd64
                └── darwin/arm64
                        │
                        ▼ (somente em tags v*)
                    release
                        ├── Binários + SHA-256
                        ├── Assinatura Cosign OIDC keyless
                        └── GitHub Release com notas automáticas
```

**Características do CI:**
- Testes Go executam em **Ubuntu e Windows** (portabilidade confirmada)
- Cobertura de código publicada no **Codecov**
- **Checkstyle Java** obrigatório — falha com violações de estilo
- golangci-lint compilado com Go 1.26 (`install-mode: goinstall`)
- Releases com binários assinados via **Sigstore/Cosign** (OIDC keyless)

---

## 6. Pirâmide de Testes

| Camada | Quantidade | Arquivo | O que verifica |
|---|:---:|---|---|
| **Unitários Go** | 13 | `internal/command/*_test.go` | sign, validate, start, stop, version, errors, http |
| **Unitários Java** | 4 | `FakeSignatureServiceTest.java` | lógica do simulador PKCS#11 |
| **Integração HTTP Java** | 5 | `SignatureControllerTest.java` | /sign, /validate, /health com servidor real |
| **Integração timer Java** | 2 | `InactivityTimerTest.java` | reset do timer, auto-shutdown |
| **E2e Go (subprocess)** | 7 | `test/e2e/cli_test.go` | binário real: versão, help, sign→validate, espaços/acentos, exit codes |
| **TOTAL** | **31** | | |

### Cenários negativos cobertos

| Cenário | Teste | Exit code esperado |
|---|---|:---:|
| `--input` ausente | `TestCLI_MissingInput_ExitCode2` | 2 (UserError) |
| Arquivo inexistente | `TestCLI_FileNotFound_ExitCode1` | 1 (SystemError) |
| Comando desconhecido | `TestCLI_UnknownCommand_ExitCode2` | 2 (UserError) |
| JSON malformado (Java) | `testSignEndpointMissingContent` | HTTP 400 |
| Assinatura inválida (Java) | `testValidateEndpointInvalidSignature` | HTTP 200 `valid=false` |

---

## 7. Contrato CLI ↔ JAR

Documentado formalmente em [`docs/api-contract.md`](./api-contract.md).

| Endpoint | Método | Request | Response |
|---|---|---|---|
| `/health` | GET | — | `{"status":"UP","port":8080,"inactivityTimeoutSeconds":300}` |
| `/sign` | POST | `{"content":"<base64>"}` | `{"signature":"<base64>","valid":true,"message":"..."}` |
| `/validate` | POST | `{"content":"<base64>","signature":"<base64>"}` | `{"signature":"...","valid":true/false,"message":"..."}` |
| `/shutdown` | POST | — | `{"message":"Servidor encerrando..."}` |

**Códigos de saída da CLI:**

| Código | Tipo | Causa |
|:---:|---|---|
| `0` | Sucesso | Operação concluída |
| `1` | SystemError | I/O, rede, processo, JSON inválido |
| `2` | UserError | Parâmetro ausente, arquivo não encontrado, comando inválido |

---

## 8. Decisões Arquiteturais (ADRs)

| ADR | Decisão | Motivação |
|---|---|---|
| [ADR-001](./adr/001-escolha-go-para-cli.md) | Go para a CLI | Binário único, multiplataforma, zero dependências runtime |
| [ADR-002](./adr/002-modo-servidor-http-padrao.md) | Modo HTTP como padrão | Isola criptografia no JAR; modo local apenas para desenvolvimento |
| [ADR-003](./adr/003-parser-cli-stdlib-flag.md) | `flag` stdlib (sem Cobra/urfave) | Zero dependências externas; controle total do parsing |
| [ADR-004](./adr/004-simulador-pkcs11.md) | FakeSignatureService para PKCS#11 | Hardware HSM indisponível; simulação documentada como decisão consciente |

---

## 9. Rastreabilidade de Issues

| Issue | Título | Status | Commits principais |
|---|---|:---:|---|
| [#11](https://github.com/Amanda23Souza/Sistema-Runner/issues/11) | Endpoints HTTP no Assinador | ✅ Fechada | `327ae50`, `efa9929` |
| [#14](https://github.com/Amanda23Souza/Sistema-Runner/issues/14) | Gestão de Ciclo de Vida (Start/Stop) | ✅ Fechada | `327ae50`, `efa9929` |
| [#15](https://github.com/Amanda23Souza/Sistema-Runner/issues/15) | Comunicação CLI-HTTP | ✅ Fechada | `327ae50`, `efa9929` |
| [#16](https://github.com/Amanda23Souza/Sistema-Runner/issues/16) | Integração PKCS (parcial) | ✅ Fechada | `327ae50` + ADR-004 |
| [#1](https://github.com/Amanda23Souza/Sistema-Runner/issues/1) | Estrutura base do CLI em Go | ✅ Fechada | commits iniciais |
| [#6](https://github.com/Amanda23Souza/Sistema-Runner/issues/6) | Simulação de Assinatura Digital | ✅ Fechada | `e7cd183` |
| [#7](https://github.com/Amanda23Souza/Sistema-Runner/issues/7) | Validação de Parâmetros | ✅ Fechada | `e7cd183` |
| [#19](https://github.com/Amanda23Souza/Sistema-Runner/issues/19) | Pipeline CI/CD | ✅ Fechada | `build.yml` |

---

## 10. Conformidade com os Critérios

Relatório completo em [`docs/conformidade-criterios.md`](./conformidade-criterios.md).

### Resumo por seção

| Seção | Critérios | ✅ | ⚠️ | ❌ | % |
|---|:---:|:---:|:---:|:---:|:---:|
| A. Princípios Transversais | 5 | 5 | 0 | 0 | **100%** |
| B. Organização do Repositório | 5 | 5 | 0 | 0 | **100%** |
| C. Documentação | 4 | 4 | 0 | 0 | **100%** |
| D. Qualidade de Código | 8 | 8 | 0 | 0 | **100%** |
| E. Requisitos Funcionais | 13 | 11 | 2 | 0 | **92%** |
| F. Build e Dependências | 4 | 4 | 0 | 0 | **100%** |
| G. Testes | 5 | 4 | 1 | 0 | **85%** |
| H. Engenharia de Processo | 5 | 5 | 0 | 0 | **100%** |
| I. Operabilidade | 3 | 3 | 0 | 0 | **100%** |
| **TOTAL** | **52** | **49** | **3** | **0** | **~96%** |

> Score: (49 + 0,5 × 3) / 52 = 50,5 / 52 ≈ **97,1%** (conservadoramente reportado como ~96%)

### Lacunas documentadas (~4%)

| Critério | Lacuna | Documentação |
|---|---|---|
| E4 | Simulador HubSaúde com ciclo de vida próprio | US-03 planejada |
| E5 | Interface PKCS#11 real (SoftHSM2/JNI) | ADR-004 — requer hardware |
| G2 | Teste e2e completo CLI→HTTP→JAR (subprocess Go + JAR real) | Planejado para sprint futura |

---

## 11. Releases

O CI gera automaticamente binários para **5 plataformas** ao criar uma tag `v*`:

| Plataforma | Arquivo |
|---|---|
| Linux x86_64 | `Sistema-Runner-v1.0.0-linux-amd64` |
| Linux ARM64 | `Sistema-Runner-v1.0.0-linux-arm64` |
| Windows x86_64 | `Sistema-Runner-v1.0.0-windows-amd64.exe` |
| macOS Intel | `Sistema-Runner-v1.0.0-darwin-amd64` |
| macOS Apple Silicon | `Sistema-Runner-v1.0.0-darwin-arm64` |

Cada binário é acompanhado de checksum SHA-256 e assinatura **Cosign OIDC keyless** (verificável sem chave privada armazenada).

**Criar nova release após merge:**
```bash
git checkout main && git pull
git tag -a v1.0.0 -m "Release v1.0.0 — implementação completa (~96% conformidade)"
git push origin v1.0.0
# CI constrói e publica automaticamente em GitHub Releases
```

---

## 12. Reprodutibilidade (clonar e executar do zero)

```bash
# 1. Clonar
git clone https://github.com/Amanda23Souza/Sistema-Runner.git
cd Sistema-Runner

# 2. Adicionar upstream (referência do professor)
git remote add upstream https://github.com/kyriosdata/runner.git

# 3. Compilar CLI Go
cd cli-assinatura
go build -o assinatura ./cmd/assinatura

# 4. Compilar JAR Java
cd ../docs/aulas/projetos/assinador-java
mvn clean package

# 5. Executar testes Go (unitários + e2e)
cd ../../../../cli-assinatura
go test -v -race -short ./...          # unitários rápidos
go test -v -timeout 120s ./test/e2e/  # e2e completo

# 6. Executar testes Java
cd ../docs/aulas/projetos/assinador-java
mvn verify
```

---

## 13. Evolução do Projeto

| Data | Marco | Conformidade |
|---|---|:---:|
| 2026-04-01 | Estrutura base do repositório | — |
| 2026-05-21 | Implementação inicial CLI + assinatura local | ~62% |
| 2026-06-09 | HTTP server, start/stop, comunicação CLI-HTTP | ~92% |
| 2026-06-10 | Testes e2e, checkstyle Java, CI corrigido, ADRs | **~96%** |

---

*Repositório: https://github.com/Amanda23Souza/Sistema-Runner*  
*PR de entrega: https://github.com/Amanda23Souza/Sistema-Runner/pull/29*  
*Relatório de conformidade completo: [`docs/conformidade-criterios.md`](./conformidade-criterios.md)*
