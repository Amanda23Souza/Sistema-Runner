# Relatório de Desenvolvimento — Sistema Runner
**Disciplina:** Implementação e Integração de Software — UFG  
**Estudante:** Marcello Ronald (mronald-js / mronaldjs)  
**Repositório:** https://github.com/Amanda23Souza/Sistema-Runner  
**Data do Relatório:** 2026-06-24  

---

## 1. Resumo Executivo

Resumo da minha participação Neste projeto entre **2026-03-18 e 2026-06-24**, cobrindo 4 sprints e 1 branch de release final. Fui responsável pela **implementação da CLI em Go** (cli-assinatura), pela **integração com o servidor Java** (assinador.jar) e pelo **pipeline CI/CD** completo e por boa parte dos itens do Sprint 4 listados no board do GitHub Projects.

---

## 2. Histórico de Contribuições por Data

### Sprint 1 — Kickoff e Planejamento (Mar 2026)

| Hash | Data | Descrição |
|------|------|-----------|
| `ab234b6` | 2026-03-18 | Sincronizei o upstream com nosso projeto principal |
| `3c3918e` | 2026-03-18 | Iniciei o planejamento da iteração |
| `39f7064` | 2026-03-18 | Propus melhorias na estrutura do projeto |
| `bd83508` | 2026-03-18 | Adequei o README ao contexto do grupo |
| `4ac1df6` | 2026-03-19 | Atualizei a documentação do projeto |
| `bf1da11` | 2026-03-19 | Restaurei seções do README principal que haviam sido removidas |
| `1a2cebc` | 2026-03-25 | Inicializei a implementação da cli-assinatura |
| `b1fcfa5` | 2026-03-25 | Fiz merge com upstream/main para manter sincronização |

**Foco:** Setup inicial do repositório, integração com o upstream (kyriosdata/runner), planejamento da CLI.

---

### Sprint 2 — Implementação Base da CLI (Abr 2026)

| Hash | Data | Descrição |
|------|------|-----------|
| `e7320b2` | 2026-04-01 | Criei a implementação inicial da CLI de assinatura |
| `db46a27` | 2026-04-08 | Adicionei documentação e planejamento do projeto |
| `59f718d` | 2026-04-07 | Organizei o backlog da iteração e reorganizei os tópicos |
| `f688fe5` | 2026-04-07 | Removi seção de detalhes e atualizei links de planejamento |
| `4d9c4e7` | 2026-04-07 | Removi READMEs de design redundantes e atualizei planejamento |
| `76a1372` | 2026-04-09 | Adicionei comandos de assinatura e validação ao CLI |
| `aa8b8b8` | 2026-04-15 | Merge do PR #17 |
| `b72ef95` | 2026-04-22 | Configurei o CI/CD e adicionei comandos de assinatura |
| `ed53134` | 2026-04-22 | Adicionei suporte completo a assinatura e validação de arquivos |
| `440806a` | 2026-04-29 | Implementei metadados de versão no build e no CLI |
| `969fcde` | 2026-04-29 | Removi arquivo desnecessário |

**Foco:** Estrutura base da CLI (Go), comandos `sign`, `validate`, `version`, pipeline CI/CD inicial, metadados de build injetados via ldflags.

---

### Sprint 3 — Backend Java e Integração HTTP (Mai 2026)

| Hash | Data | Descrição |
|------|------|-----------|
| `277ede8` | 2026-05-06 | Adaptei o código para não conflitar com a estrutura do professor |
| `71f67be` | 2026-05-06 | Adicionei suporte a aliases (`criar`/`validar`) e saída JSON nos comandos |
| `e7cd183` | 2026-05-20 | Implementei a REST API Java com Javalin e escrevi testes de integração |
| `e45b028` | 2026-05-20 | Implementei a estrutura inicial dos HTTP endpoints |
| `9cc4025` | 2026-05-27 | Corrigi a conformidade com os critérios A-I avaliados |
| `efa9929` | 2026-06-09 | Implementei o gerenciamento de ciclo de vida (start/stop/status) e documentei o contrato de API |

**Foco:** Servidor Java com Javalin (`/sign`, `/validate`, `/health`, `/shutdown`), testes de integração, comunicação CLI↔HTTP, aliases de comandos.

---

### Sprint 4 — Release e Conformidade (Jun 2026)

| Hash | Data | Descrição |
|------|------|-----------|
| `327ae50` | 2026-06-10 | Refatorei alinhando com upstream, corrigi CI/CD e escrevi testes para atingir 96% de conformidade |
| `e41b602` | 2026-06-10 | Atualizei README com estrutura de testes e2e e checkstyle |
| `fd2ce50` | 2026-06-10 | Corrigi falhas no CI: checkstyle Java, versão Go, split de testes e2e |
| `f8db865` | 2026-06-10 | Corrigi versão Go no go.mod e adicionei builds arm64 |
| `628eff5` | 2026-06-10 | Fixei o pin do golangci-lint para suporte ao Go 1.26 |
| `617e3ce` | 2026-06-10 | Revertei o pin do golangci-lint para latest |
| `f1428ef` | 2026-06-10 | Corrigi o install-mode do golangci-lint para goinstall |
| `5ec390a` | 2026-06-10 | Adicionei entrega-final.md e corrigi versão Go no conformidade |
| `786e3fd` | 2026-06-24 | **Implementei todos os itens pendentes do Sprint 4 (release final)** |

**Foco:** CI multi-plataforma (Ubuntu + Windows + macOS + arm64), cobertura de testes, conformidade com critérios, release final.

---

## 3. O que Implementei no Release Final (2026-06-24)

Commit `786e3fd5ae18e05d417f1414a99d8d3e6636d7b4`:

### 3.1 Pacote `cmd/` — Parsing de Comandos (#8)
Adicionei os arquivos que seguem a estrutura esperada pelo professor:

```
cli-assinatura/cmd/
├── root.go       ← Execute() com switch/case dispatcher
├── sign.go       ← runSign() → internal/command.SignCmd
├── validate.go   ← runValidate() → internal/command.ValidateCmd
└── root_test.go  ← 9 testes do Execute()
```

O `Execute()` em `cmd/root.go` implementa o dispatcher com switch/case conforme esperado:

```go
switch args[0] {
case "sign", "criar":       return runSign(args[1:])
case "validate", "validar": return runValidate(args[1:])
// ...
}
```

### 3.2 Provisionamento Automático do JDK (#10)
Criei `cli-assinatura/internal/command/jdk.go` com 3 estratégias em cascata:

```
1. exec.LookPath("java")       → usa java do PATH do sistema
2. ~/.assinador/jdk/bin/java   → usa cache local provisionado
3. downloadAndProvisionJDK()   → baixa JDK 21 da Adoptium automaticamente
```

Download via `https://api.adoptium.net/v3/binary/latest/21/ga/{OS}/{ARCH}/jdk/...`  
Extração nativa em Go: `extractTarGZ` (Linux/Mac) + `extractZip` (Windows) — sem dependências externas.

### 3.3 Correção da Invocação Real do assinador.jar (#9)
Corrigi um bug crítico nos modos HTTP de sign e validate:

**Antes (sign.go L120) — bug:**
```go
body := fmt.Sprintf(`{"content": "%s"}`, c.input)  // enviava o CAMINHO, não o conteúdo!
```

**Depois — correto:**
```go
inputData, _ := os.ReadFile(c.input)
bodyBytes, _ := json.Marshal(map[string]string{"content": string(inputData)})
httpPost(url, string(bodyBytes))  // envia o conteúdo real do arquivo com JSON válido
```

Apliquei o mesmo fix em `validate.go`.

### 3.4 Testes Elaborados (#20)
Adicionei:
- 9 testes do dispatcher `Execute()` em `cmd/root_test.go`
- `TestLocalJDKBin_ReturnsValidPath` e `TestLocalJDKDir_ContainsAssinadorDir`
- `TestStartCmd_ResolveJava_UsesPathIfAvailable`
- `TestSignCmd_Run_HTTPMode_ConnRefused` e `TestValidateCmd_Run_HTTPMode_ConnRefused`
- `TestSignCmd_Run_HTTPMode_FileNotFound`

### 3.5 Compilação Automática do JAR (#21 Refinamento)
Adicionei `buildJar()` ao `start.go`, que compila o assinador.jar via Maven automaticamente:
```go
func (c *StartCmd) buildJar() error {
    mvn, _ := exec.LookPath("mvn")
    cmd := exec.Command(mvn, "clean", "package", "-DskipTests", ...)
    cmd.Dir = "docs/aulas/projetos/assinador-java"
    return cmd.Run()
}
```

---

## 4. Arquitetura que Implementei

```
┌─────────────────────────────────────────────────────┐
│             CLI (Go) — cli-assinatura/               │
│                                                      │
│  cmd/assinatura/main.go  ← ponto de entrada          │
│  cmd/root.go (Execute)   ← dispatcher (switch/case)  │
│                                                      │
│  internal/command/                                   │
│    root.go    ← orquestra todos os subcomandos       │
│    sign.go    ← sign/criar (HTTP ou local)           │
│    validate.go← validate/validar (HTTP ou local)     │
│    start.go   ← inicia JAR + resolveJava() + PID    │
│    stop.go    ← /shutdown + kill por PID             │
│    jdk.go     ← provisionamento automático JDK 21    │
│    http.go    ← cliente HTTP com timeout             │
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

## 5. Fluxo Completo que Implementei (Modo HTTP)

```bash
# 1. Detecto java no PATH / provisiono JDK 21 automaticamente
# 2. Compilo assinador.jar se não encontrado (mvn package)
# 3. Inicio servidor em background, aguardo /health
assinatura start

# 4. Leio conteúdo do arquivo → serializo JSON → POST /sign
# 5. Recebo {"signature": "MOCKED_SIGNATURE_BASE64_=="} → salvo em .sig
assinatura sign --input documento.pdf

# 6. Leio conteúdo + assinatura → POST /validate → valid: true/false
assinatura validate --input documento.pdf --signature documento.pdf.sig

# 7. Encerro via /shutdown (graceful) ou PID (fallback)
assinatura stop
```

---

## 6. CI/CD que Configurei

Pipeline em `.github/workflows/build.yml`:

| Job | Plataformas | Objetivo |
|-----|-------------|----------|
| `test-java` | ubuntu-latest | `mvn clean verify` (checkstyle + testes) |
| `lint` | ubuntu-latest | golangci-lint |
| `test` | ubuntu + windows | `go test -race -short ./...` + e2e |
| `build` | linux/amd64, linux/arm64, windows/amd64, darwin/amd64, darwin/arm64 | binários release |
| `release` | ubuntu-latest | GitHub Release (somente em tags `v*`) |

---

## 7. Prova de Integridade — Hashes SHA-1 dos Meus Commits

| Commit completo (SHA-1) | Data | Entrega |
|------------------------|------|---------|
| `786e3fd5ae18e05d417f1414a99d8d3e6636d7b4` | 2026-06-24 | **Release Final — Sprint 4** |
| `5ec390ac33ef9ba4789acb00bc6c5df2d4480c31` | 2026-06-10 | Documentação entrega final |
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

## 8. Meus Dias de Trabalho no Repositório

A tabela abaixo combina commits e atividade de Pull Requests (abertura e merge). Dias marcados com **[PR]** representam entregas via Pull Request sem commit direto naquela data.

| Data | Tipo | O que fiz |
|------|------|-----------|
| 2026-03-18 | commit | Iniciei o projeto, sincronizei upstream, organizei o README |
| 2026-03-19 | commit | Corrigi documentação e restaurei seções do README |
| 2026-03-25 | commit | Inicializei a implementação da cli-assinatura |
| 2026-04-01 | commit + PR #17 aberto | Criei a implementação inicial da CLI em Go e abri o PR |
| 2026-04-07 | commit | Reorganizei documentação e planejamento |
| 2026-04-08 | commit | Documentei o projeto |
| 2026-04-09 | commit | Implementei os comandos sign e validate |
| 2026-04-16 | PR #17 merged | PR da implementação inicial da CLI foi integrado ao main |
| 2026-04-22 | commit | Configurei CI/CD e adicionei suporte a assinatura/validação de arquivos |
| 2026-04-23 | PR #24 aberto | Abri PR com CI/CD e Simulador CLI |
| 2026-04-29 | commit | Implementei metadados de versão |
| 2026-04-30 | PR #24 merged | PR de CI/CD e simulador foi integrado ao main |
| 2026-05-06 | commit | Adicionei aliases e saída JSON; adaptei para estrutura do professor |
| 2026-05-07 | PR #25 merged | PR de aliases + JSON output integrado |
| 2026-05-20 | commit | Implementei o backend Java (REST API Javalin) e escrevi testes |
| 2026-05-21 | PR #26 merged | PR da Task #11 (HTTP endpoints) integrado |
| 2026-05-27 | commit | Corrigi a conformidade com critérios avaliados |
| 2026-05-28 | PR #27 merged | PR de conformidade critérios A-I integrado |
| 2026-06-09 | commit | Implementei gerenciamento de ciclo de vida (start/stop/status) |
| 2026-06-10 | commit | Finalizei Sprint 4: CI multi-plataforma, cobertura de testes, checkstyle |
| 2026-06-11 | PR #28 + PR #29 merged | PRs de lifecycle management e Sprint 4 integrados ao main |
| 2026-06-24 | commit | **Release Final**: pacote cmd/, JDK provisioning, correção PKCS, testes |

**Total de dias de contribuição ativa: 22 dias** ao longo de ~100 dias de disciplina (contando tanto commits quanto atividade de PR).

### Pull Requests abertos e mergeados por mim

| PR | Título | Aberto | Mergeado |
|----|--------|--------|----------|
| #17 | Adiciona implementação inicial da CLI de assinatura | 2026-04-01 | 2026-04-16 |
| #24 | CI/CD e Simulador CLI | 2026-04-23 | 2026-04-30 |
| #25 | Adição de suporte a aliases e saída JSON nos comandos CLI | 2026-05-07 | 2026-05-07 |
| #26 | Task #11 do board (HTTP endpoints) | 2026-05-21 | 2026-05-21 |
| #27 | fix(conformidade): aplicar correções de critérios A-I | 2026-05-28 | 2026-05-28 |
| #28 | feat: implement command structure for lifecycle management | 2026-06-11 | 2026-06-11 |
| #29 | Marcello alterações (Sprint 4) | 2026-06-11 | 2026-06-11 |

---

## 9. Como Verificar minha Autenticidade

```bash
# Ver detalhes e arquivos do release final
git show --stat 786e3fd5ae18e05d417f1414a99d8d3e6636d7b4

# Ver o diff completo do que implementei no release final
git diff 62fb587 786e3fd

# Ver todos os meus commits na disciplina
git log --author="mronald" --oneline --all

# Verificar integridade do repositório
git fsck --full
```

---

*Relatório gerado a partir do histórico git do repositório em 2026-06-24.*
