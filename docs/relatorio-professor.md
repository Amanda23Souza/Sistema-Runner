# Relatório de Desenvolvimento — Sistema Runner
**Disciplina:** Implementação e Integração de Software — UFG  
**Estudante:** Marcello Ronald (mronald-js / mronaldjs)  
**Repositório:** https://github.com/Amanda23Souza/Sistema-Runner  
**Data do Relatório:** 2026-06-24  

---

## 1. Resumo Executivo

O desenvolvimento ocorreu entre **2026-03-18 e 2026-06-24**, cobrindo 4 sprints e 1 branch de release final. O estudante contribuiu com **implementação da CLI em Go** (cli-assinatura), **integração com o servidor Java** (assinador.jar), **pipeline CI/CD** completo, e todos os itens do Sprint 4 listados no board do GitHub Projects.

---

## 2. Histórico de Contribuições por Data

### Sprint 1 — Kickoff e Planejamento (Mar 2026)

| Hash | Data | Descrição |
|------|------|-----------|
| `ab234b6` | 2026-03-18 | Sincronizando upstream com projeto principal |
| `3c3918e` | 2026-03-18 | Início do planejamento |
| `39f7064` | 2026-03-18 | Proposta de melhorias |
| `bd83508` | 2026-03-18 | Adequações no README |
| `4ac1df6` | 2026-03-19 | Atualizações na documentação |
| `bf1da11` | 2026-03-19 | Fix: restaurando partes do Readme principal |
| `1a2cebc` | 2026-03-25 | Inicializar implementação da cli-assinatura |
| `b1fcfa5` | 2026-03-25 | Merge upstream/main |

**Foco:** Setup inicial do repositório, integração com upstream (kyriosdata/runner), planejamento da CLI.

---

### Sprint 2 — Implementação Base da CLI (Abr 2026)

| Hash | Data | Descrição |
|------|------|-----------|
| `e7320b2` | 2026-04-01 | Implementação inicial da CLI de assinatura |
| `db46a27` | 2026-04-08 | Documentação e planejamento do projeto |
| `59f718d` | 2026-04-07 | Backlog da iteração e reorganização |
| `f688fe5` | 2026-04-07 | Remove seção de detalhes, atualiza links |
| `4d9c4e7` | 2026-04-07 | Remove READMEs de design, atualiza planejamento |
| `76a1372` | 2026-04-09 | Adiciona comandos de assinatura e validação ao CLI |
| `aa8b8b8` | 2026-04-15 | Merge PR #17 |
| `b72ef95` | 2026-04-22 | Adiciona CI/CD e comandos de assinatura |
| `ed53134` | 2026-04-22 | Suporte a assinatura e validação de arquivos |
| `440806a` | 2026-04-29 | Metadados de versão ao build e CLI |
| `969fcde` | 2026-04-29 | Removendo arquivo desnecessário |

**Foco:** Estrutura base da CLI (Go), comandos `sign`, `validate`, `version`, pipeline CI/CD inicial, metadados de build.

---

### Sprint 3 — Backend Java e Integração HTTP (Mai 2026)

| Hash | Data | Descrição |
|------|------|-----------|
| `277ede8` | 2026-05-06 | Adaptação para não conflitar com código do professor |
| `71f67be` | 2026-05-06 | Suporte a aliases e saída JSON nos comandos CLI |
| `e7cd183` | 2026-05-20 | Implementação da REST API em Java (Javalin) + testes de integração |
| `e45b028` | 2026-05-20 | Implementação inicial dos HTTP endpoints |
| `9cc4025` | 2026-05-27 | Fix conformidade: correções critérios A-I |
| `efa9929` | 2026-06-09 | Estrutura de lifecycle management (start/stop/status) + contrato API |

**Foco:** Servidor Java com Javalin (`/sign`, `/validate`, `/health`, `/shutdown`), testes de integração Java, comunicação CLI↔HTTP, aliases `criar`/`validar`.

---

### Sprint 4 — Release e Conformidade (Jun 2026)

| Hash | Data | Descrição |
|------|------|-----------|
| `327ae50` | 2026-06-10 | Refactor: alinhamento upstream, CI/CD, testes para 96% conformidade |
| `e41b602` | 2026-06-10 | Docs: README com estrutura de testes e2e e checkstyle |
| `fd2ce50` | 2026-06-10 | Fix CI: checkstyle Java, versão Go, split de testes e2e |
| `f8db865` | 2026-06-10 | Fix CI: versão Go no go.mod + builds arm64 |
| `628eff5` | 2026-06-10 | Fix CI: pin golangci-lint v1.64.9 para suporte Go 1.26 |
| `617e3ce` | 2026-06-10 | Fix CI: reverter pin golangci-lint para latest |
| `f1428ef` | 2026-06-10 | Fix CI: install-mode goinstall no golangci-lint |
| `5ec390a` | 2026-06-10 | Docs: entrega-final.md + correção versão Go |
| `786e3fd` | 2026-06-24 | **feat(release-final): implementar todos os itens Sprint 4** |

**Foco:** CI multi-plataforma (Ubuntu + Windows + macOS + arm64), cobertura de testes, conformidade com critérios, release final.

---

## 3. Implementações do Release Final (2026-06-24)

Commit `786e3fd5ae18e05d417f1414a99d8d3e6636d7b4`:

### 3.1 Pacote `cmd/` — Parsing de Comandos (#8)
Adição dos arquivos que seguem a estrutura esperada pelo professor:

```
cli-assinatura/cmd/
├── root.go       ← Execute() com switch/case dispatcher
├── sign.go       ← runSign() → internal/command.SignCmd
├── validate.go   ← runValidate() → internal/command.ValidateCmd
└── root_test.go  ← 9 testes do Execute()
```

O `Execute()` em `cmd/root.go` aceita os mesmos comandos e aliases que `internal/command`:

```go
switch args[0] {
case "sign", "criar":   return runSign(args[1:])
case "validate", "validar": return runValidate(args[1:])
// ...
}
```

### 3.2 Provisionamento Automático do JDK (#10)
Novo arquivo `cli-assinatura/internal/command/jdk.go` com 3 estratégias em cascata:

```
1. exec.LookPath("java")       → usa java do PATH
2. ~/.assinador/jdk/bin/java   → usa cache local  
3. downloadAndProvisionJDK()   → baixa JDK 21 da Adoptium
```

Download via `https://api.adoptium.net/v3/binary/latest/21/ga/{OS}/{ARCH}/jdk/...`  
Extração nativa em Go: `extractTarGZ` (Linux/Mac) + `extractZip` (Windows).

### 3.3 Invocação Real do assinador.jar (#9)
Correção de bug crítico nos modos HTTP:

**Antes (sign.go L120):**
```go
body := fmt.Sprintf(`{"content": "%s"}`, c.input)  // enviava o CAMINHO, não o conteúdo!
```

**Depois:**
```go
inputData, _ := os.ReadFile(c.input)
bodyBytes, _ := json.Marshal(map[string]string{"content": string(inputData)})
httpPost(url, string(bodyBytes))  // envia o conteúdo real do arquivo
```

O mesmo fix foi aplicado em `validate.go`.

### 3.4 Testes Adicionais (#20)
- `TestExecute_Version`, `TestExecute_Help`, `TestExecute_NoArgs`, ...
- `TestLocalJDKBin_ReturnsValidPath`, `TestLocalJDKDir_ContainsAssinadorDir`
- `TestStartCmd_ResolveJava_UsesPathIfAvailable`
- `TestSignCmd_Run_HTTPMode_ConnRefused`, `TestValidateCmd_Run_HTTPMode_ConnRefused`

### 3.5 Compilação Automática do JAR
`start.go` agora compila o assinador.jar via Maven se o JAR não for encontrado:
```go
func (c *StartCmd) buildJar() error {
    mvn, _ := exec.LookPath("mvn")
    cmd := exec.Command(mvn, "clean", "package", "-DskipTests", ...)
    cmd.Dir = "docs/aulas/projetos/assinador-java"
    return cmd.Run()
}
```

---

## 4. Arquitetura Implementada

```
┌─────────────────────────────────────────────────────┐
│             CLI (Go) — cli-assinatura/               │
│                                                      │
│  cmd/assinatura/main.go  ←── ponto de entrada        │
│  cmd/root.go (Execute)   ←── dispatcher (switch/case)│
│                                                      │
│  internal/command/                                   │
│    root.go    ← orquestra todos os subcomandos       │
│    sign.go    ← sign/criar (HTTP ou local)           │
│    validate.go← validate/validar (HTTP ou local)     │
│    start.go   ← inicia JAR + resolveJava() + PID     │
│    stop.go    ← /shutdown + kill por PID             │
│    jdk.go     ← provisionamento automático JDK 21    │
│    http.go    ← cliente HTTP c/ timeout              │
└────────────────────┬────────────────────────────────┘
                     │ POST /sign, /validate, /health
                     ▼
┌─────────────────────────────────────────────────────┐
│         assinador.jar (Java 21 + Javalin)           │
│                                                      │
│  App.java            ← servidor HTTP porta 8080      │
│  SignatureController ← /sign, /validate endpoints    │
│  FakeSignatureService← simulação PKCS#11             │
│                                                      │
│  Auto-shutdown: timer 5 min reiniciado a cada req    │
└─────────────────────────────────────────────────────┘
```

---

## 5. Fluxo Completo (Modo HTTP)

```bash
# 1. CLI detecta java / provisiona JDK 21 automaticamente
# 2. Compila assinador.jar se necessário (mvn package)
# 3. Inicia servidor em background, aguarda /health
assinatura start

# 4. Lê conteúdo do arquivo → JSON → POST /sign
# 5. Recebe {"signature": "MOCKED_SIGNATURE_BASE64_=="} → salva .sig
assinatura sign --input documento.pdf

# 6. Lê conteúdo + assinatura → POST /validate → valid: true/false
assinatura validate --input documento.pdf --signature documento.pdf.sig

# 7. Encerra via /shutdown (graceful) ou PID (fallback)
assinatura stop
```

---

## 6. CI/CD — Matriz de Build

Pipeline em `.github/workflows/build.yml`:

| Job | Plataformas | Objetivo |
|-----|-------------|----------|
| `test-java` | ubuntu-latest | `mvn clean verify` (checkstyle + testes) |
| `lint` | ubuntu-latest | golangci-lint |
| `test` | ubuntu + windows | `go test -race -short ./...` + e2e |
| `build` | linux/amd64, linux/arm64, windows/amd64, darwin/amd64, darwin/arm64 | binários release |
| `release` | ubuntu-latest | GitHub Release (somente em tags `v*`) |

---

## 7. Prova de Integridade — Hashes SHA-1 dos Commits Principais

| Commit completo (SHA-1) | Data | Entrega |
|------------------------|------|---------|
| `786e3fd5ae18e05d417f1414a99d8d3e6636d7b4` | 2026-06-24 | **Release Final — Sprint 4** |
| `5ec390ac33ef9ba4789acb00bc6c5df2d4480c31` | 2026-06-10 | Documentação entrega final |
| `62fb5878a3072d2b57d7d971178fed5461786768` | 2026-06-10 | Merge PR #29 (Sprint 4) |
| `f1428effdd27fce42c7cbe00ebcee82d21e460e8` | 2026-06-10 | Fix CI: golangci-lint |
| `f8db865abf3faaae8a9e4cf6e2d749dd4ac1ae40` | 2026-06-10 | Fix CI: Go 1.26 + arm64 |
| `fd2ce50076c1dbb30e3a537e063f76126b467613` | 2026-06-10 | Fix CI: checkstyle + e2e |
| `327ae50caf15f692342448c1855642525a4a73f6` | 2026-06-10 | Refactor: CI + testes 96% |
| `efa9929d479248d82b3d200e99ba0a93b5d87543` | 2026-06-09 | feat: lifecycle + API contract |
| `9cc402570ec229dbd8296f4933da799dc809548a` | 2026-05-27 | fix: conformidade critérios A-I |
| `e7cd183cff17edc1a09a01f0975dc82d6a84e5cf` | 2026-05-20 | feat: REST API Java + testes |
| `e45b028102527589e48abfbb2270284f170ac83e` | 2026-05-20 | feat: HTTP endpoints inicial |
| `71f67be39c025ba0142d08ba320d50abdc94872d` | 2026-05-06 | feat: aliases + JSON output |
| `440806a8762b60faf08082ca420822a651f9507f` | 2026-04-29 | feat: metadados de versão |
| `76a1372df58bfe6db923b6eb6ef6952786fe5101` | 2026-04-09 | feat: comandos sign + validate |
| `e7320b2acf8f1abee66f2cb73bbf75344bd889f7` | 2026-04-01 | feat: implementação inicial CLI |
| `3c3918e27e803241e3ddb695c242cc8b1cac0e2d` | 2026-03-18 | docs: início do planejamento |

---

## 8. Dias de Trabalho no Repositório

| Data | Atividade Principal |
|------|---------------------|
| 2026-03-18 | Kickoff — planejamento, sync upstream, README |
| 2026-03-19 | Documentação e fixes iniciais |
| 2026-03-25 | Inicialização da cli-assinatura |
| 2026-04-01 | Implementação inicial CLI Go |
| 2026-04-07 | Reorganização de documentação |
| 2026-04-08 | Documentação do projeto |
| 2026-04-09 | Comandos sign + validate |
| 2026-04-15 | Merge e revisão |
| 2026-04-22 | CI/CD + assinatura/validação de arquivos |
| 2026-04-29 | Metadados de versão |
| 2026-05-06 | Aliases, JSON output, sync professor |
| 2026-05-20 | Backend Java (REST API Javalin) + testes |
| 2026-05-27 | Fix conformidade critérios |
| 2026-06-09 | Lifecycle management (start/stop/status) |
| 2026-06-10 | Sprint 4: CI multi-plataforma, cobertura, checkstyle |
| 2026-06-24 | **Release Final**: cmd/, JDK provisioning, PKCS fix |

**Total de dias de contribuição: 16 dias ativos** ao longo de ~100 dias de disciplina.

---

## 9. Verificação de Autenticidade

Para verificar qualquer commit listado neste relatório:

```bash
# Verificar hash completo
git show --stat 786e3fd5ae18e05d417f1414a99d8d3e6636d7b4

# Ver diff completo do release final
git diff 62fb587 786e3fd

# Ver todos os commits do estudante
git log --author="mronald" --oneline

# Verificar integridade do repositório
git fsck --full
```

---

*Relatório gerado automaticamente a partir do histórico git do repositório.*
