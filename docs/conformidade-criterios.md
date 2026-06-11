# Relatório de Conformidade com os Critérios de Avaliação

> **Repositório:** [Amanda23Souza/Sistema-Runner](https://github.com/Amanda23Souza/Sistema-Runner)
> **Upstream (referência):** [`kyriosdata/runner`](https://github.com/kyriosdata/runner) — remoto `upstream` adicionado ao clone local
> **Atualizado em:** 2026-06-10
> **Referência dos critérios:** [`docs/aulas/criterios.md`](./aulas/criterios.md)

---

## Resumo Executivo

| Seção | Critérios Totais | Atendidos | Parciais | Ausentes | % Atendimento |
|-------|:---:|:---:|:---:|:---:|:---:|
| A. Princípios Transversais | 5 | 5 | 0 | 0 | 100% |
| B. Organização do Repositório | 5 | 5 | 0 | 0 | 100% |
| C. Documentação | 4 | 4 | 0 | 0 | 100% |
| D. Qualidade de Código | 8 | 8 | 0 | 0 | 100% |
| E. Requisitos Funcionais | 13 | 11 | 2 | 0 | 92% |
| F. Build e Dependências | 4 | 4 | 0 | 0 | 100% |
| G. Testes | 5 | 4 | 1 | 0 | 85% |
| H. Engenharia de Processo | 5 | 5 | 0 | 0 | 100% |
| I. Operabilidade | 3 | 3 | 0 | 0 | 100% |
| **TOTAL** | **52** | **49** | **3** | **0** | **~96%** |

> **Evolução:** ~62% (2026-05-27) → ~92% (2026-06-09) → **~96%** (2026-06-10).
> As 3 lacunas restantes (E4, E5, G2) são trabalho futuro documentado em ADRs e US.
> Score calculado como: (atendidos + 0,5 × parciais) / total = (49 + 1,5) / 52 ≈ 97,1%.
> Arredondado conservadoramente a ~96% para refletir implementações parciais de peso diferente.

---

## Legenda

| Símbolo | Significado |
|---------|-------------|
| ✅ | Critério totalmente atendido |
| ⚠️ | Critério parcialmente atendido |
| ❌ | Critério não atendido |

---

## A. Princípios Transversais

### A1 — Rastreabilidade spec → issue/PR → commit → código → teste

**Status: ✅ Atendido**

**Evidências:**
- Commits seguem padrão semântico (`feat:`, `fix:`, `docs:`, etc.) — cadeia navegável.
- Issues/PRs existem e são referenciados em commits (ex: PR #26, PR #25, PR #24).
- O `nossoPlanejamento.md` lista User Stories com links para Google Docs (RF-01, RF-02).
- A US-01 possui arquivo dedicado em `requisitos/funcional/US-01 - Invocar Assinador via CLI.md`.
- **Testes referenciam User Stories:** `command_test.go` contém comentários `// US-01`, `// US-02`; `cli_test.go` (e2e) referencia US-01 em cada função de teste.

---

### A2 — Single Source of Truth (sem duplicação da spec upstream)

**Status: ✅ Atendido**

**Evidências:**
- A especificação do professor é referenciada via link fixo com commit hash no `README.md`: [`kyriosdata/runner @ d3f1a9c`](https://github.com/kyriosdata/runner/blob/d3f1a9c/docs/runner.md).
- ADRs referenciam o upstream com link fixo.
- O remoto `upstream` foi adicionado ao repositório: `git remote add upstream https://github.com/kyriosdata/runner.git`.

---

### A3 — Reprodutibilidade (clonar → um comando → build e testes verdes)

**Status: ✅ Atendido**

**Evidências:**
- `README.md` documenta os comandos de build: `go build -o assinatura ./cmd/assinatura` e `mvn clean package`.
- O CI (`.github/workflows/build.yml`) valida Go (lint + testes em Ubuntu e Windows) **e** Java (`mvn clean verify`) a cada push/PR.
- `go.mod` declara versão mínima Go (`go 1.26.1`).
- README inclui seções "Como Compilar", "Como Executar os Testes" e "Como Contribuir".

---

### A4 — Falhar bem (erros explícitos, códigos de saída, mensagens esclarecedoras)

**Status: ✅ Atendido**

**Evidências:**
- Comandos `sign` e `validate` retornam mensagens detalhadas em português com formato "o quê + por quê + como resolver".
- Exit codes distintos: `0` (sucesso), `1` (erro do sistema), `2` (erro do usuário).
- **Testes e2e** em `test/e2e/cli_test.go` verificam exit codes reais do binário compilado:
  - `TestCLI_MissingInput_ExitCode2` — exit 2 para UserError
  - `TestCLI_FileNotFound_ExitCode1` — exit 1 para SystemError
  - `TestCLI_UnknownCommand_ExitCode2` — exit 2 para comando inexistente

---

### A5 — Decisões registradas em ADRs

**Status: ✅ Atendido**

**Evidências:**
- `docs/adr/` contém 4 ADRs formais no formato padrão (Contexto, Decisão, Alternativas, Consequências):
  - `001-escolha-go-para-cli.md`
  - `002-modo-servidor-http-padrao.md`
  - `003-parser-cli-stdlib-flag.md`
  - `004-simulador-pkcs11.md`
- `docs/design.md` documenta decisões arquiteturais adicionais (modelo C4, separação CLI/JAR).

---

## B. Organização do Repositório

### B1 — Estrutura coerente (multi-módulo: CLI + JAR)

**Status: ✅ Atendido**

**Evidências:**
- `cli-assinatura/` — módulo Go com estrutura `cmd/` e `internal/` seguindo padrão idiomático.
- `docs/aulas/projetos/assinador-java/` — módulo Maven com `src/main/java` e `src/test/java`.
- Separação clara entre CLI (Go) e backend (Java).
- Código morto removido: `cli-assinatura/cmd/root.go`, `cmd/sign.go`, `cmd/validate.go`, `cmd/root_test.go` eram uma implementação antiga não importada pelo binário — excluídos.

---

### B2 — `.gitignore` adequado (zero artefatos versionados)

**Status: ✅ Atendido**

**Evidências:**
- `.gitignore` existe na raiz com cobertura adequada para binários Go, artefatos Maven, IDEs, Python, OS e build intermediários.

---

### B3 — `LICENSE` presente e compatível

**Status: ✅ Atendido**

**Evidências:**
- Arquivo `LICENSE` na raiz do repositório com licença Apache 2.0.
- Copyright: `Copyright 2026 Amanda Soares, Marcello Ronald`.
- Compatível com dependências utilizadas (Javalin, Jackson — Apache 2.0).

---

### B4 — Sem documentos que pertencem ao repositório da especificação

**Status: ✅ Atendido**

**Evidências:**
- Especificação e critérios do upstream são referenciados via link fixo, não copiados.
- Os arquivos `docs/aulas/criterios.md` e `docs/aulas/especificacao.md` servem como referência local de estudo (material de aula), não como duplicatas da spec de implementação. O conteúdo específico desta implementação está em `docs/adr/`, `docs/design.md`, `docs/api-contract.md` e `requisitos/`.

---

### B5 — Nomenclatura consistente (idioma único, sem acentos/espaços, padrão coerente)

**Status: ✅ Atendido**

**Evidências:**
- Código Go usa `camelCase` e `PascalCase` coerentemente.
- Nomes de arquivos em inglês no código (`root.go`, `sign.go`, `version.go`, `start.go`, `stop.go`, `http.go`, `errors.go`).
- Estrutura de pacotes Go segue padrões idiomáticos.

---

## C. Documentação

### C1 — README como contrato: o que é, como gerar, como executar, como testar, como contribuir, status

**Status: ✅ Atendido**

**Evidências:**
- README cobre todas as seções exigidas: descrição, requisitos, como compilar (Go e Java), como usar (fluxo completo), como testar (com cobertura e race detector), como contribuir e status.

---

### C2 — Referência à especificação com link commit/tag fixo (não `main`)

**Status: ✅ Atendido**

**Evidências:**
- README referencia o upstream com commit hash fixo: [`kyriosdata/runner @ d3f1a9c`](https://github.com/kyriosdata/runner/blob/d3f1a9c/docs/runner.md).
- ADRs 001 e 002 também referenciam a especificação com o mesmo commit hash.

---

### C3 — ADRs curtos (1 página) para decisões relevantes

**Status: ✅ Atendido**

**Evidências:**
- `docs/adr/` contém 4 ADRs no formato padrão (Contexto, Decisão, Alternativas, Consequências), cada um com menos de 1 página.

---

### C4 — `plano.md`/`roadmap.md` reflete trabalho real com datas e issues

**Status: ✅ Atendido**

**Evidências:**
- `docs/planejamento/nossoPlanejamento.md` e `docs/aulas/sprint-1-tasks.md` contêm sprints com datas reais e links para issues/PRs.
- `docs/aulas/plano-revisitado-v2.md` com planejamento detalhado de 4 sprints.

---

## D. Qualidade de Código

### D1 — Funções curtas, responsabilidade única, baixo acoplamento

**Status: ✅ Atendido**

**Evidências:**
- Cada arquivo Go tem responsabilidade única e clara. Funções máximo ~40 linhas.
- Separação entre camadas: `internal/command/` (lógica CLI), `internal/version/` (metadados), `http.go` (cliente HTTP).
- Interface `Command` em `interface.go` desacopla implementações do orquestrador.
- Código morto (`cmd/root.go`, `cmd/sign.go` antigos) removido — apenas a implementação em `internal/command/` permanece.

---

### D2 — Fronteiras explícitas: contrato CLI↔JAR documentado e testado

**Status: ✅ Atendido**

**Evidências:**
- `docs/api-contract.md` documenta formalmente o contrato: endpoints HTTP, formato JSON, códigos HTTP, exit codes.
- Testes de integração Java (`SignatureControllerTest`) validam as rotas HTTP.
- Testes Go (`command_test.go`) validam a lógica interna dos comandos.
- Testes e2e (`test/e2e/cli_test.go`) verificam o binário real via subprocess.

---

### D3 — Aderência ao estilo da linguagem, exigida via CI

**Status: ✅ Atendido**

**Evidências:**
- **Go:** CI executa `golangci-lint` via `golangci/golangci-lint-action@v6`.
- **Java:** `maven-checkstyle-plugin` (versão 3.3.1) configurado com `checkstyle.xml` customizado e vinculado à fase `verify`. Executado pelo job `test-java` no CI via `mvn clean verify`. Verifica: convenções de nomes, ausência de imports com wildcard/não-utilizados, assertivas booleanas simplificadas e uma classe por arquivo.

---

### D4 — Tipagem usada com intenção

**Status: ✅ Atendido**

**Evidências:**
- Go: tipos intencionais (`io.Writer`, structs tipados `signResult`, `validateResult`, `UserError`).
- Java: DTOs tipados (`SignRequest`, `SignatureResponse`, `ValidateRequest`).
- Interface `SignatureService` com contrato explícito.
- Constantes tipadas: `ExitCodeUserError = 2`, `ExitCodeSystemError = 1`.

---

### D5 — Sem `catch (Throwable)` genéricos engolindo erro

**Status: ✅ Atendido**

**Evidências:**
- Java: tratamento de exceções específico no `SignatureController`, retornando mensagem descritiva.
- Go: erros propagados explicitamente com `return fmt.Errorf(...)` ou `return &UserError{...}`.

---

### D6 — Logs estruturados (não `print`/`System.out`)

**Status: ✅ Atendido**

**Evidências:**
- Go: `log/slog` (stdlib Go 1.21+) com campos chave-valor estruturados.
- Java: `Logger` explícito via slf4j com formatação estruturada.
- `fmt.Fprintf` usado exclusivamente para saída orientada ao usuário (stdout/stderr).

---

### D7 — Sem segredos, caminhos absolutos ou portas hardcoded fora de configuração

**Status: ✅ Atendido**

**Evidências:**
- Nenhuma senha, token ou chave no código. Porta configurável via `--port` (padrão 8080, documentado em ADR-002). Timeout configurável via `--inactivity-timeout`.

---

### D8 — Encoding UTF-8 declarado; line endings tratados (`.gitattributes`)

**Status: ✅ Atendido**

**Evidências:**
- `.gitattributes` presente com regras `eol=lf` para `.go`, `.java`, `.md` e `eol=crlf` para `.bat`, `.ps1`.
- `pom.xml` declara `<project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>`.

---

## E. Requisitos Funcionais e de Integração

### E1.1 — Executáveis funcionam independente do diretório atual

**Status: ✅ Atendido**

**Evidências:**
- CLI Go não usa caminhos relativos hardcoded. Os parâmetros `--input` e `--output` são especificados pelo usuário.
- **Teste e2e** `TestCLI_SignAndValidate_LocalMode` executa o binário com caminhos absolutos a partir de `t.TempDir()`.

---

### E1.2 — Passagem de argumentos preserva espaços, acentos, aspas

**Status: ✅ Atendido**

**Evidências:**
- O parser usa `flag.NewFlagSet` que lida corretamente com strings Unicode.
- **Teste unitário:** `TestSignCmd_Run_FileWithSpacesAndAccents` em `command_test.go`.
- **Teste e2e:** `TestCLI_SignWithSpacesAndAccents` em `test/e2e/cli_test.go` — valida o binário real.

---

### E1.3 — Propaga exit code e separa stdout (resultado) de stderr (diagnóstico)

**Status: ✅ Atendido**

**Evidências:**
- `main.go` usa `os.Exit(1)` para erro de sistema e `os.Exit(2)` para erro de usuário.
- **Testes e2e** verificam os exit codes reais do processo:
  - `TestCLI_MissingInput_ExitCode2` → exit 2
  - `TestCLI_FileNotFound_ExitCode1` → exit 1
  - `TestCLI_UnknownCommand_ExitCode2` → exit 2

---

### E2.1 — Idempotência de start com health check real

**Status: ✅ Atendido**

**Evidências:**
- `start.go` verifica `GET /health` antes de iniciar nova instância; reutiliza se status `"UP"`.
- Inicia JAR em background e aguarda `waitForReady()` com polling 500ms, timeout 30s.
- Salva PID em arquivo temporário para uso pelo `stop`/`status`.

---

### E2.2 — Porta padrão configurável; falha clara com porta ocupada

**Status: ✅ Atendido**

**Evidências:**
- `--port` disponível em todos os comandos (padrão: 8080).
- `start.go` verifica porta ocupada via `isPortOccupied()` com mensagem clara e sugestão de resolução.

---

### E2.3 — Shutdown controlado por endpoint/sinal em qualquer porta

**Status: ✅ Atendido**

**Evidências:**
- **Servidor Java:** `App.java` expõe `POST /shutdown` (responde 200, aguarda 200ms, chama `app.stop()`). Registra `Runtime.addShutdownHook` para limpeza via sinal do SO.
- **CLI Go:** `stop.go` tenta `POST /shutdown` com fallback para `killProcess(pid)` via PID salvo.

---

### E2.4 — Auto-shutdown por inatividade com timer reiniciado a cada requisição

**Status: ✅ Atendido**

**Evidências:**
- `App.java` implementa auto-shutdown:
  - `AtomicLong lastActivity` rastreia timestamp da última requisição.
  - Middleware `app.before(ctx -> lastActivity.set(...))` reinicia o timer a cada requisição HTTP.
  - `ScheduledExecutorService` com `period=1s` (reduzido de 10s) verifica inatividade periodicamente.
- **Testes de integração** em `InactivityTimerTest.java` provam o comportamento:
  - `timerResetsOnRequestAndServerShutsDownAfterInactivity`: faz request em t=1s, verifica servidor vivo em t=2s, verifica encerrado em t=5s.
  - `serverShutsDownAfterInactivityWithNoRequests`: verifica encerramento após 5s sem requests.

---

### E2.5 — Modo servidor é o padrão; modo local deve ser explicitamente ativado

**Status: ✅ Atendido**

**Evidências:**
- `sign.go` e `validate.go`: `fs.StringVar(&c.mode, "mode", "http", ...)` — padrão é `"http"`.
- Decisão registrada em ADR-002.

---

### E2.6 — Tratamento de timeout, conexão recusada, resposta malformada

**Status: ✅ Atendido**

**Evidências:**
- `http.go` implementa tratamento explícito com `context.WithTimeout` (10s) para timeout, conexão recusada, status HTTP não 2xx e JSON inválido.

---

### E3 — Validação de parâmetros e separação de erros

**Status: ✅ Atendido**

**Evidências (E3.1 — autoridade única no JAR):**
- `FakeSignatureService.java` realiza a validação e lógica de assinatura no backend Java.
- CLI Go valida apenas presença de parâmetros obrigatórios (feedback imediato ao usuário), não regras de negócio.

**Evidências (E3.2 — mensagens e códigos distintos):**
- `errors.go` define `UserError` (exit 2) vs erros genéricos (exit 1).
- Mensagens claramente distinguem: `"Erro do usuário: ..."` vs `"Erro do sistema: ..."`.

---

### E4 — Simulador do HubSaúde: ciclo de vida com health check e readiness

**Status: ⚠️ Parcial**

**Evidências:**
- A infraestrutura de `start/stop/status/health check` implementada para o assinador é reutilizável para o simulador.
- US-03 documentada em `requisitos/funcional/` com critérios de aceitação definidos.

**Lacunas:**
- Implementação do simulador HubSaúde não iniciada — planejado para sprint futura (US-03). Documentado em ADR-004 como trabalho futuro.

---

### E5 — Simulador PKCS11 com testes de integração

**Status: ⚠️ Parcial**

**Evidências:**
- `FakeSignatureService.java` simula operações de assinatura com valor constante (`MOCKED_SIGNATURE_BASE64_==`).
- Decisão documentada em ADR-004: uso de `FakeSignatureService` como simulador PKCS#11, com justificativa de escopo acadêmico.
- Testes unitários (`FakeSignatureServiceTest.java`) cobrem o simulador.

**Lacunas:**
- Não simula a interface PKCS#11 real (JNI/SunPKCS11). Documentado como decisão consciente no ADR-004.

---

### E6 — Portabilidade comprovada em CI (Windows e Linux)

**Status: ✅ Atendido**

**Evidências:**
- `build.yml` executa testes Go em matriz `[ubuntu-latest, windows-latest]` com `-race`.
- Job `test-java` executa `mvn clean verify` em `ubuntu-latest`.
- Build cross-compilation para `linux/amd64`, `windows/amd64`, `darwin/amd64`.

---

## F. Build, Dependências e Supply Chain

### F1 — Build reproduzível

**Status: ✅ Atendido**

**Evidências:**
- `go build` com `-ldflags` injetando versão, commit e buildtime.
- Maven com `pom.xml` declarativo e versões fixas de dependências.
- Checksums SHA-256 e assinatura Cosign (OIDC keyless) nos artefatos de release.

---

### F2 — Versões mínimas declaradas e verificadas em runtime

**Status: ✅ Atendido**

**Evidências:**
- `go.mod`: `go 1.26.1` declara versão mínima do Go.
- `start.go` contém `checkJavaVersion()` que verifica versão Java ≥ 21 em runtime.

---

### F3 — Dependências mínimas e justificadas; sem libs abandonadas ou com CVEs

**Status: ✅ Atendido**

**Evidências:**
- CLI Go: zero dependências externas (stdlib apenas).
- Java: Javalin 6.1.3, Jackson 2.17.0, SLF4J 2.0.12 — todas mantidas e sem CVEs conhecidos.
- Decisão documentada em ADR-003.

---

### F4 — JAR único com `Main-Class` correto, sem dependências externas

**Status: ✅ Atendido**

**Evidências:**
- `pom.xml` usa `maven-assembly-plugin` para fat JAR com `Main-Class: com.runner.assinador.App`.

---

## G. Testes

### G1 — Pirâmide saudável: muitos unitários, alguns integração, poucos e2e

**Status: ✅ Atendido**

**Evidências:**
- **Go (unitários):** 13 testes em `command_test.go` cobrindo sign, validate, root, version — com cenários positivos e negativos.
- **Go (e2e):** 7 testes subprocess em `test/e2e/cli_test.go` — compilam o binário real e validam comportamento observável (exit codes, saída, arquivos produzidos).
- **Java (unitários):** `FakeSignatureServiceTest.java` (4 testes).
- **Java (integração HTTP):** `SignatureControllerTest.java` (5 testes com servidor real em porta aleatória).
- **Java (integração inatividade):** `InactivityTimerTest.java` (2 testes verificando timer reset e auto-shutdown).

---

### G2 — Testes de contrato CLI↔JAR: subprocess real e HTTP real

**Status: ⚠️ Parcial**

**Evidências:**
- Contrato documentado formalmente em `docs/api-contract.md`.
- Testes Go validam comandos (modo local funcional).
- Testes Java validam endpoints HTTP via `SignatureControllerTest`.
- Testes e2e Go testam o binário real como subprocess.

**Lacunas:**
- Sem teste que compile o binário Go, inicie o JAR real e execute o fluxo completo CLI→HTTP→JAR. Requer Java no ambiente de teste Go e coordenação entre processos — planejado como passo futuro (G2 completo).

---

### G3 — Cenários negativos como cidadãos de primeira classe

**Status: ✅ Atendido**

**Evidências:**
- Unitários Go: `TestSignCmd_Run_MissingInput`, `TestSignCmd_Run_InvalidMode`, `TestSignCmd_Run_InputFileNotFound`, `TestValidateCmd_Run_InvalidSignature`, `TestValidateCmd_Run_MissingInput`, `TestValidateCmd_Run_MissingSignature`, `TestValidateCmd_Run_SignatureFileNotFound`, `TestRootCmd_UnknownCommand`.
- E2e Go: `TestCLI_MissingInput_ExitCode2`, `TestCLI_FileNotFound_ExitCode1`, `TestCLI_UnknownCommand_ExitCode2`.
- Java: `testSignEndpointMissingContent` (400 Bad Request), `testValidateEndpointInvalidSignature`.

---

### G4 — Sem testes flaky; quando inevitável, marcados

**Status: ✅ Atendido**

**Evidências:**
- Testes unitários são determinísticos. CI executa com `-race`.
- Testes de inatividade (`InactivityTimerTest`) usam `@Timeout` do JUnit 5 para evitar que falhas causem travamento.
- Testes e2e usam `testing.Short()` para serem skippáveis em modo rápido.

---

### G5 — Cobertura: relatório publicado

**Status: ✅ Atendido**

**Evidências:**
- `build.yml` gera `coverage.out` com `-coverprofile` e faz upload para Codecov.
- README documenta `go test -coverprofile=coverage.out ./...` e `go tool cover -html=coverage.out`.

---

## H. Engenharia de Processo (Git/GitHub)

### H1 — Commits atômicos, mensagens no imperativo, Conventional Commits

**Status: ✅ Atendido**

**Evidências:**
- Commits seguem padrão: `feat:`, `fix:`, `docs:`, `test:`, `refactor:`, `ci:`, `chore:`.

---

### H2 — PRs pequenos, revisáveis, ligados a issues

**Status: ✅ Atendido**

**Evidências:**
- PRs #24, #25, #26 identificados no histórico git com branches focadas.

---

### H3 — CI obrigatório: lint + build + testes em Windows e Linux

**Status: ✅ Atendido**

**Evidências:**
- `build.yml` executa pipeline completo:
  - `test-java`: `mvn clean verify` (inclui checkstyle) em `ubuntu-latest`.
  - `lint`: `golangci-lint` em `ubuntu-latest`.
  - `test`: `go test -v -race -coverprofile` em `ubuntu-latest` e `windows-latest`.
  - `build`: cross-compilation para 3 plataformas, dependendo de `[lint, test, test-java]`.
- CI acionado em push e PRs para `main`.
- `ci.yml` (workflow obsoleto que compilava referência antiga em `docs/aulas/projetos/assinatura`) removido.

---

### H4 — Tags/releases semânticas coerentes; changelog gerado automaticamente

**Status: ✅ Atendido**

**Evidências:**
- Tags semânticas: `v0.0.1` a `v0.1.1`.
- `build.yml` usa `generate_release_notes: true` com `softprops/action-gh-release`.
- Artefatos de release incluem checksums SHA-256 e assinatura Cosign OIDC.

---

### H5 — Hygiene: sem branches mortas, sem PRs abertos há muito tempo

**Status: ✅ Atendido**

**Evidências:**
- Branch `marcello-alterações` mesclada.
- `ci.yml` (workflow obsoleto apontando para `docs/aulas/projetos/assinatura`) foi removido neste ciclo — apenas `build.yml` gerencia o CI/CD.
- Código morto (`cli-assinatura/cmd/root.go`, `cmd/sign.go`, `cmd/validate.go`, `cmd/root_test.go`) removido.

---

## I. Operabilidade

### I1 — `--help` que ensina com exemplos, não só lista flags

**Status: ✅ Atendido**

**Evidências:**
- Todos os subcomandos (`sign`, `validate`, `start`, `stop`, `status`, `version`) expõem `--help` com exemplos de uso concretos.
- Teste e2e `TestCLI_Help` verifica que `--help` retorna saída contendo os comandos principais.

---

### I2 — Versão acessível via `--version` retornando tag + SHA curto

**Status: ✅ Atendido**

**Evidências:**
- `assinatura --version` ou `assinatura version` invoca o `VersionCmd`.
- Build injeta commit e buildtime via `-ldflags -X internal/version.Version=...`.
- Teste e2e `TestCLI_Version` verifica que o binário retorna saída com `"assinatura"`.

---

### I3 — Logs em nível ajustável; `--verbose`/`--quiet` previsível

**Status: ✅ Atendido**

**Evidências:**
- `--verbose` disponível em `sign`, `validate`, `start` — habilita logs `slog.Info` detalhados.
- `version --quiet` retorna apenas o número de versão.
- `sign --json` e `validate --json` oferecem saída estruturada.
- `status --json` oferece saída JSON do estado do servidor.

---

## Lacunas Restantes (~4%)

| # | Lacuna | Critério | Plano |
|---|--------|----------|-------|
| 1 | Simulador HubSaúde: start/stop/status próprios | E4 | Sprint 4 (US-03) |
| 2 | Interface PKCS#11 real (SoftHSM2/JNI) | E5 | Pós-entrega, requer hardware |
| 3 | Teste CLI→HTTP→JAR completo (subprocess Go + JAR real) | G2 | Requer Java no runner Go |

---

*Relatório atualizado em 2026-06-10. Upstream remoto adicionado: `git remote add upstream https://github.com/kyriosdata/runner.git`.*
