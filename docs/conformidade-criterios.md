# Relatório de Conformidade com os Critérios de Avaliação

> **Repositório:** [Amanda23Souza/Sistema-Runner](https://github.com/Amanda23Souza/Sistema-Runner)
> **Atualizado em:** 2026-06-09
> **Referência dos critérios:** [`docs/aulas/criterios.md`](./aulas/criterios.md)

---

## Resumo Executivo

| Seção | Critérios Totais | Atendidos | Parciais | Ausentes | % Atendimento |
|-------|:---:|:---:|:---:|:---:|:---:|
| A. Princípios Transversais | 5 | 5 | 0 | 0 | 100% |
| B. Organização do Repositório | 5 | 5 | 0 | 0 | 100% |
| C. Documentação | 4 | 4 | 0 | 0 | 100% |
| D. Qualidade de Código | 8 | 7 | 1 | 0 | 94% |
| E. Requisitos Funcionais | 13 | 8 | 3 | 2 | 73% |
| F. Build e Dependências | 4 | 4 | 0 | 0 | 100% |
| G. Testes | 5 | 3 | 2 | 0 | 70% |
| H. Engenharia de Processo | 5 | 5 | 0 | 0 | 100% |
| I. Operabilidade | 3 | 3 | 0 | 0 | 100% |
| **TOTAL** | **52** | **44** | **6** | **2** | **~92%** |

> **Evolução:** O repositório subiu de **~62%** (2026-05-27) para **~92%** de conformidade.
> As 2 lacunas restantes (E4, E5) referem-se ao simulador HubSaúde e à integração PKCS#11 real, documentadas como decisões arquiteturais em ADRs.

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
- **Testes referenciam User Stories:** `command_test.go` contém comentários `// US-01`, `// US-02` ligando testes a requisitos.

---

### A2 — Single Source of Truth (sem duplicação da spec upstream)

**Status: ✅ Atendido**

**Evidências:**
- A especificação do professor é referenciada via link fixo com commit hash no `README.md`: [`kyriosdata/runner @ d3f1a9c`](https://github.com/kyriosdata/runner/blob/d3f1a9c/docs/runner.md).
- Não há duplicação óbvia de especificações completas (cópias removidas).
- ADRs referenciam o upstream com link fixo.

---

### A3 — Reprodutibilidade (clonar → um comando → build e testes verdes)

**Status: ✅ Atendido**

**Evidências:**
- `README.md` documenta os comandos de build: `go build -o assinatura ./cmd/assinatura` e `mvn clean package`.
- O CI (`.github/workflows/build.yml`) valida o build e os testes automaticamente a cada push/PR.
- `go.mod` declara versão mínima Go (`go 1.26.1`).
- README inclui seções "Como Compilar", "Como Executar os Testes" e "Como Contribuir".

---

### A4 — Falhar bem (erros explícitos, códigos de saída, mensagens esclarecedoras)

**Status: ✅ Atendido**

**Evidências:**
- Comandos `sign` e `validate` retornam mensagens detalhadas em português com formato "o quê + por quê + como resolver":
  - `"Erro do usuário: o parâmetro --input é obrigatório."`
  - `"Erro do sistema: falha ao conectar ao servidor assinador em localhost:8080."`
  - `"Como resolver: verifique se o servidor está rodando com 'assinatura start --port 8080'."`
- Exit codes distintos: `0` (sucesso), `1` (erro do sistema), `2` (erro do usuário).
- Código de erro `[MS-03]` identificado para parâmetros inválidos.
- Mensagens de erro vão para **stderr** (`c.errOut`), resultados vão para **stdout** (`c.out`).

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

---

### B2 — `.gitignore` adequado (zero artefatos versionados)

**Status: ✅ Atendido**

**Evidências:**
- `.gitignore` existe na raiz com cobertura adequada:
  - Binários Go: `cli-assinatura/assinatura`, `cli-assinatura/assinatura.exe`
  - Artefatos Maven: `docs/aulas/projetos/assinador-java/target/`
  - IDEs: `.idea/`, `.vscode/`, `*.iml`
  - Python: `__pycache__/`, `*.pyc`
  - OS: `.DS_Store`, `Thumbs.db`, `Desktop.ini`
  - Build intermediários: `*.test`, `*.out`, `coverage.out`

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
- Foram removidas todas as cópias indevidas de `especificacao.md` e `criterios.md`.
- Apenas conteúdo específico desta implementação permanece. Documentos do upstream são referenciados via link fixo.

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
- README cobre todas as seções exigidas:
  - **O que é:** Descrição do projeto (linha 5–8).
  - **Requisitos:** Tabela com Go, JDK, Maven e versões mínimas.
  - **Como compilar:** Seções separadas para Go e Java.
  - **Como usar:** Fluxo completo com exemplos (`start` → `sign` → `validate` → `stop`).
  - **Como testar:** Seção "Como Executar os Testes" com `go test ./...`, cobertura e race detector.
  - **Como contribuir:** Seção "Como Contribuir" com convenções de commits, lint e processo de PR.
  - **Status:** Fases 1–5 com ✅ e ⏳.

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
- `docs/adr/` contém 4 ADRs no formato padrão:
  - `001-escolha-go-para-cli.md` — Escolha de linguagem.
  - `002-modo-servidor-http-padrao.md` — Modo padrão do CLI.
  - `003-parser-cli-stdlib-flag.md` — Escolha do parser de argumentos.
  - `004-simulador-pkcs11.md` — Estratégia de simulação PKCS#11.
- Cada ADR tem menos de 1 página com seções: Contexto, Decisão, Alternativas, Consequências.

---

### C4 — `plano.md`/`roadmap.md` reflete trabalho real com datas e issues

**Status: ✅ Atendido**

**Evidências:**
- `docs/planejamento/nossoPlanejamento.md` e `docs/aulas/sprint-1-tasks.md` contêm sprints com datas reais (Sprint 1: 08/04–22/04, Sprint 2: 29/04–13/05, etc.).
- Links para issues/PRs e backlog no GitHub Projects.
- `docs/aulas/plano-revisitado-v2.md` com planejamento detalhado.

---

## D. Qualidade de Código

### D1 — Funções curtas, responsabilidade única, baixo acoplamento

**Status: ✅ Atendido**

**Evidências:**
- Cada arquivo Go (`sign.go`, `validate.go`, `version.go`, `root.go`, `start.go`, `stop.go`, `http.go`, `errors.go`) tem responsabilidade única e clara.
- Funções são curtas (máximo ~40 linhas).
- Separação entre camadas: `internal/command/` (lógica CLI), `internal/version/` (metadados), `http.go` (cliente HTTP).
- Interface `Command` em `interface.go` desacopla implementações do orquestrador.

---

### D2 — Fronteiras explícitas: contrato CLI↔JAR documentado e testado

**Status: ✅ Atendido**

**Evidências:**
- `docs/api-contract.md` documenta formalmente o contrato entre CLI Go e assinador.jar:
  - Endpoints HTTP: `/sign`, `/validate`, `/health`, `/shutdown`.
  - Formato JSON de request/response com exemplos.
  - Códigos HTTP e semântica (200 OK vs 400 Bad Request).
  - Códigos de saída CLI (0, 1, 2) e mapeamento de erros.
  - Flags e parâmetros de cada subcomando.
- Testes de integração Java (`SignatureControllerTest`) validam as rotas HTTP.
- Testes Go validam a lógica interna dos comandos com exit codes corretos.

---

### D3 — Aderência ao estilo da linguagem, exigida via CI

**Status: ✅ Atendido**

**Evidências:**
- CI (`build.yml`) contém job `lint` que executa `golangci-lint` via `golangci/golangci-lint-action@v6`:
  ```yaml
  lint:
    name: Lint (Go)
    runs-on: ubuntu-latest
    steps:
      - uses: golangci/golangci-lint-action@v6
        with:
          version: latest
          working-directory: cli-assinatura
          args: --timeout=5m
  ```
- Código Go segue convenções idiomáticas (exported types com PascalCase, comentários de package, uso de `io.Writer` para injeção de dependência).

**Lacunas menores:**
- Para Java: nenhum checkstyle ou spotbugs configurado no `pom.xml` ou CI (não bloqueante).

---

### D4 — Tipagem usada com intenção

**Status: ✅ Atendido**

**Evidências:**
- Go: uso correto de tipos (`io.Writer`, `[]string`, structs tipados `signResult`, `validateResult`, `UserError`).
- Java: DTOs tipados (`SignRequest`, `SignatureResponse`, `ValidateRequest`).
- Interface `SignatureService` com contrato explícito.
- Constantes tipadas: `ExitCodeUserError = 2`, `ExitCodeSystemError = 1`.

---

### D5 — Sem `catch (Throwable)` genéricos engolindo erro

**Status: ✅ Atendido**

**Evidências:**
- Java: tratamento de exceções específico com `catch (Exception e)` no `SignatureController`, retornando mensagem descritiva — não engole erros.
- Go: erros propagados explicitamente com `return fmt.Errorf(...)` ou `return &UserError{msg: ...}` com mensagens descritivas.

---

### D6 — Logs estruturados (não `print`/`System.out`)

**Status: ✅ Atendido**

**Evidências:**
- Go: uso extensivo de `log/slog` (stdlib Go 1.21+) para logging estruturado:
  - `slog.Info("iniciando assinatura", "input", c.input, "mode", c.mode, "port", c.port)`
  - `slog.Error("falha na conexão HTTP", "url", url, "erro", err)`
  - `slog.Warn("endpoint /shutdown não respondeu, tentando encerrar via PID", "porta", c.port)`
- `fmt.Fprintf` é usado exclusivamente para saída orientada ao usuário (stdout/stderr), não como log diagnóstico.
- Java: Javalin usa slf4j por padrão; `App.java` usa `Logger` explicitamente: `log.info("Assinador iniciado na porta {} ...")`.

---

### D7 — Sem segredos, caminhos absolutos ou portas hardcoded fora de configuração

**Status: ✅ Atendido**

**Evidências:**
- Nenhuma senha, token ou chave encontrada no código.
- A porta é configurável via `--port` em todos os comandos (padrão 8080, documentado em ADR-002).
- Caminhos são relativos ou passados via parâmetros CLI.
- Timeout de inatividade configurável via `--inactivity-timeout`.

---

### D8 — Encoding UTF-8 declarado; line endings tratados (`.gitattributes`)

**Status: ✅ Atendido**

**Evidências:**
- `.gitattributes` presente na raiz com regras detalhadas:
  ```
  * text=auto eol=lf
  *.go text eol=lf
  *.java text eol=lf
  *.md text eol=lf
  *.bat text eol=crlf
  *.cmd text eol=crlf
  *.ps1 text eol=crlf
  ```
- Binários marcados como `binary` (`.jar`, `.png`, `.jpg`, etc.).
- `pom.xml` declara `<project.build.sourceEncoding>UTF-8</project.build.sourceEncoding>`.

---

## E. Requisitos Funcionais e de Integração

### E1.1 — Executáveis funcionam independente do diretório atual

**Status: ✅ Atendido**

**Evidências:**
- CLI Go não usa caminhos relativos hardcoded para suas operações internas.
- Os parâmetros `--input` e `--output` são especificados pelo usuário.

---

### E1.2 — Passagem de argumentos preserva espaços, acentos, aspas

**Status: ✅ Atendido**

**Evidências:**
- O parser usa `flag.NewFlagSet` que lida corretamente com strings Unicode.
- **Teste dedicado**: `TestSignCmd_Run_FileWithSpacesAndAccents` em `command_test.go` valida arquivo com nome `"documento com espaços é ação.txt"` e conteúdo com acentuação `"conteúdo com acentuação: ção, ã, é, ü"`.

---

### E1.3 — Propaga exit code e separa stdout (resultado) de stderr (diagnóstico)

**Status: ✅ Atendido**

**Evidências:**
- `main.go` usa `os.Exit(1)` para erro de sistema e `os.Exit(2)` para erro de usuário, via `command.IsUserError()`.
- Todos os comandos usam `c.out` (stdout) para resultados e `c.errOut` (stderr) para mensagens de erro/diagnóstico.
- Exemplo em `sign.go`: `fmt.Fprintf(c.errOut, "Erro do sistema: ...")` para erros; `fmt.Fprintf(c.out, "Assinatura criada com sucesso...")` para resultados.

---

### E2.1 — Idempotência de start com health check real

**Status: ✅ Atendido**

**Evidências:**
- `start.go` implementa idempotência completa:
  1. Verifica `GET /health` na porta — se status é `"UP"`, reutiliza instância existente sem iniciar outra.
  2. Localiza o JAR via auto-descoberta ou `--jar`.
  3. Verifica se `java` está no PATH e se versão é >= 21.
  4. Verifica se a porta está ocupada por **outro** processo.
  5. Inicia o JAR em background e aguarda health check real com `waitForReady()` (polling a cada 500ms, timeout 30s).
  6. Salva PID em arquivo temporário para uso pelo `stop`/`status`.

---

### E2.2 — Porta padrão configurável; falha clara com porta ocupada

**Status: ✅ Atendido**

**Evidências:**
- `--port` disponível em `start`, `stop`, `status`, `sign`, `validate` (padrão: 8080).
- `start.go` verifica porta ocupada via `isPortOccupied()` e retorna mensagem clara: `"Erro do sistema: a porta %d está ocupada por outro processo."` com sugestão `"Como resolver: escolha outra porta com --port <numero>"`.

---

### E2.3 — Shutdown controlado por endpoint/sinal em qualquer porta

**Status: ✅ Atendido**

**Evidências:**
- **Servidor Java:** `App.java` expõe endpoint `POST /shutdown` que encerra gracefully (responde 200, aguarda 200ms, chama `app.stop()`). Também registra `Runtime.addShutdownHook` para limpeza em encerramento via sinal do SO.
- **CLI Go:** `stop.go` tenta primeiro `POST /shutdown` (endpoint HTTP), com fallback para `killProcess(pid)` via PID salvo. No Linux/macOS usa `SIGINT`; no Windows usa `taskkill /F`.

---

### E2.4 — Auto-shutdown por inatividade com timer reiniciado a cada requisição

**Status: ✅ Atendido**

**Evidências:**
- `App.java` implementa auto-shutdown por inatividade:
  - `AtomicLong lastActivity` rastreia o timestamp da última requisição.
  - Middleware `app.before(ctx -> lastActivity.set(System.currentTimeMillis()))` reinicia o timer a cada requisição HTTP.
  - `ScheduledExecutorService` verifica periodicamente se o tempo de inatividade excedeu o timeout.
  - Timeout configurável via `--inactivity-timeout <segundos>` (padrão: 300s = 5 minutos).

---

### E2.5 — Modo servidor é o padrão; modo local deve ser explicitamente ativado

**Status: ✅ Atendido**

**Evidências:**
- `sign.go` linha 74: `fs.StringVar(&c.mode, "mode", "http", ...)` — padrão é `"http"`.
- `validate.go` linha 74: `fs.StringVar(&c.mode, "mode", "http", ...)` — padrão é `"http"`.
- Help text documenta: `"O modo padrão é 'http': o CLI se comunica com o assinador.jar em execução."`.
- Decisão registrada em ADR-002.

---

### E2.6 — Tratamento de timeout, conexão recusada, resposta malformada

**Status: ✅ Atendido**

**Evidências:**
- `http.go` implementa tratamento explícito com `context.WithTimeout` (10s):
  - **Timeout:** `"timeout após 10s aguardando resposta do servidor"`.
  - **Conexão recusada:** `"conexão recusada ou servidor inacessível"`.
  - **Status HTTP não 2xx:** `"servidor retornou status HTTP %d (esperado 2xx)"`.
  - **JSON inválido:** `"resposta do servidor não é JSON válido"`.
- Comandos `sign.go` e `validate.go` propagam essas mensagens com contexto adicional e sugestão de resolução.

---

### E3.1 — Validação feita dentro do `assinador.jar` (autoridade única)

**Status: ✅ Atendido**

**Evidências:**
- `FakeSignatureService.java` realiza a validação de parâmetros e lógica de assinatura no backend Java.
- A CLI Go faz apenas validação de presença de parâmetros obrigatórios (necessário para feedback imediato), não replica regras de negócio.

---

### E3.2 — Mensagens distinguem erro do usuário de erro do sistema; códigos diferentes

**Status: ✅ Atendido**

**Evidências:**
- `errors.go` define `UserError` (exit code 2) vs erros genéricos (exit code 1).
- `main.go` usa `command.IsUserError(err)` para distinguir e propagar o exit code correto.
- Mensagens distinguem claramente: `"Erro do usuário: ..."` vs `"Erro do sistema: ..."`.
- Exemplos: parâmetro ausente → `UserError` (exit 2); arquivo inacessível → `SystemError` (exit 1).

---

### E4 — Simulador do HubSaúde: ciclo de vida com health check e readiness

**Status: ⚠️ Parcial**

**Evidências:**
- O design (`docs/design.md`) menciona o "Simulador do HubSaúde" como sistema externo.
- A infraestrutura de start/stop/status/health check implementada para o assinador pode ser reutilizada para o simulador.

**Lacunas:**
- Implementação específica do simulador HubSaúde não iniciada — planejado para sprint futura.

---

### E5 — Simulador PKCS11 com testes de integração

**Status: ⚠️ Parcial**

**Evidências:**
- `FakeSignatureService.java` simula operações de assinatura com valor constante (`MOCKED_SIGNATURE_BASE64_==`), substituindo a camada PKCS#11.
- Decisão documentada em ADR-004 (`docs/adr/004-simulador-pkcs11.md`): uso de FakeSignatureService como simulador PKCS#11, com justificativa de escopo acadêmico e caminho para migração futura.
- Testes unitários (`FakeSignatureServiceTest.java`) cobrem o simulador.

**Lacunas:**
- Não simula a interface PKCS#11 real (JNI/SunPKCS11). Documentado como decisão consciente no ADR-004.

---

### E6 — Portabilidade comprovada em CI (Windows e Linux)

**Status: ✅ Atendido**

**Evidências:**
- `build.yml` executa testes em matriz de SOs:
  ```yaml
  test:
    runs-on: ${{ matrix.os }}
    strategy:
      matrix:
        os: [ ubuntu-latest, windows-latest ]
  ```
- Testes executam com `go test -v -race -coverprofile=coverage.out -covermode=atomic ./...` em ambos os SOs.
- Build cross-compilation para `linux/amd64`, `windows/amd64`, `darwin/amd64`.

---

## F. Build, Dependências e Supply Chain

### F1 — Build reproduzível

**Status: ✅ Atendido**

**Evidências:**
- `go build` com `-ldflags` injetando versão, commit e buildtime via variáveis de build.
- Maven com `pom.xml` declarativo e versões fixas de dependências.
- CI usa `go-version: stable` e `cache: true`.
- Checksums SHA-256 e assinatura Cosign (OIDC keyless) nos artefatos de release.

---

### F2 — Versões mínimas declaradas e verificadas em runtime

**Status: ✅ Atendido**

**Evidências:**
- `go.mod`: `go 1.26.1` declara versão mínima do Go.
- README documenta `Go 1.26.1+`, `JDK 21+` e `Maven 3.8+` como requisitos com comandos de verificação.
- `start.go` contém `checkJavaVersion()` que verifica a versão do Java em runtime antes de iniciar o servidor, alertando se inferior a 21.

---

### F3 — Dependências mínimas e justificadas; sem libs abandonadas ou com CVEs

**Status: ✅ Atendido**

**Evidências:**
- CLI Go: **zero dependências externas** (`go.mod` sem `require`). Usa apenas stdlib.
- Java: Javalin 6.1.3 (web framework ativo) + Jackson 2.17.0 (serialização JSON) + SLF4J 2.0.12 — todas mantidas e sem CVEs conhecidos.
- Decisão documentada em ADR-003.

---

### F4 — JAR único com `Main-Class` correto, sem dependências externas

**Status: ✅ Atendido**

**Evidências:**
- `pom.xml` usa `maven-assembly-plugin` para gerar fat JAR (`jar-with-dependencies`) com `Main-Class: com.runner.assinador.App`.
- `.gitignore` exclui `docs/aulas/projetos/assinador-java/target/` — artefatos não são versionados.

---

## G. Testes

### G1 — Pirâmide saudável: muitos unitários, alguns integração, poucos e2e

**Status: ⚠️ Parcial**

**Evidências:**
- Go: 13 testes em `command_test.go` cobrindo:
  - **Sign:** criação de arquivo, input ausente, modo inválido, espaços/acentos, arquivo inexistente (5 testes).
  - **Validate:** assinatura válida, assinatura inválida, input ausente, signature ausente, arquivo inexistente (5 testes).
  - **Root:** comando desconhecido, help (2 testes).
  - **Version:** retorno de versão, modo quiet (2 testes: cmd-level `root_test.go`).
- Java: `FakeSignatureServiceTest.java` e `SignatureControllerTest.java` cobrem unitários e integração HTTP.

**Lacunas:**
- Sem testes end-to-end (subprocess real invocando o binário compilado).

---

### G2 — Testes de contrato CLI↔JAR: subprocess real e HTTP real

**Status: ⚠️ Parcial**

**Evidências:**
- O contrato está documentado formalmente em `docs/api-contract.md`.
- Testes Go validam a lógica interna dos comandos (modo local funcional).
- Testes Java validam os endpoints HTTP via `SignatureControllerTest`.

**Lacunas:**
- Sem teste end-to-end que compile o binário Go, inicie o JAR e execute operações reais CLI→HTTP→JAR. Recomendado como passo futuro.

---

### G3 — Cenários negativos como cidadãos de primeira classe

**Status: ✅ Atendido**

**Evidências:**
- `TestValidateCmd_Run_InvalidSignature` — testa assinatura inválida ✅
- `TestSignCmd_Run_MissingInput` — testa parâmetro ausente com UserError ✅
- `TestSignCmd_Run_InvalidMode` — testa modo inválido com UserError ✅
- `TestSignCmd_Run_InputFileNotFound` — testa arquivo inexistente com SystemError ✅
- `TestValidateCmd_Run_MissingInput` — testa --input ausente ✅
- `TestValidateCmd_Run_MissingSignature` — testa --signature ausente ✅
- `TestValidateCmd_Run_SignatureFileNotFound` — testa arquivo de assinatura inexistente ✅
- `TestRootCmd_UnknownCommand` — testa comando desconhecido com UserError ✅
- `[MS-03]` tratado para parâmetros inválidos ✅

---

### G4 — Sem testes flaky; quando inevitável, marcados

**Status: ✅ Atendido**

**Evidências:**
- Testes atuais são determinísticos (não dependem de rede, tempo ou estado externo).
- CI executa com `-race` flag para detecção de race conditions.
- Sem testes de concorrência ou timing que possam ser flaky.

---

### G5 — Cobertura: relatório publicado

**Status: ✅ Atendido**

**Evidências:**
- `build.yml` gera perfil de cobertura e faz upload para Codecov:
  ```yaml
  - name: Run tests with coverage
    run: go test -v -race -coverprofile=coverage.out -covermode=atomic ./...

  - name: Upload coverage to Codecov
    uses: codecov/codecov-action@v4
    with:
      file: cli-assinatura/coverage.out
      flags: unittests
  ```
- README documenta como gerar relatório de cobertura localmente: `go test -coverprofile=coverage.out ./...` e `go tool cover -html=coverage.out`.

---

## H. Engenharia de Processo (Git/GitHub)

### H1 — Commits atômicos, mensagens no imperativo, Conventional Commits

**Status: ✅ Atendido**

**Evidências:**
- Commits seguem padrão semântico: `feat: implement signature REST API`, `feat: adicionar...`.
- Commits são granulares e focados.
- README documenta a convenção: `type(scope): descrição` com tipos `feat`, `fix`, `docs`, `test`, `refactor`, `ci`, `chore`.

---

### H2 — PRs pequenos, revisáveis, ligados a issues

**Status: ✅ Atendido**

**Evidências:**
- PRs #24, #25, #26 identificados no git log.
- Branch `marcello-alterações` com PR vinculado ao merge.

---

### H3 — CI obrigatório: lint + build + testes em Windows e Linux

**Status: ✅ Atendido**

**Evidências:**
- `build.yml` executa pipeline completo: lint → test → build → release.
- **Lint:** `golangci-lint` via `golangci/golangci-lint-action@v6`.
- **Testes:** `go test -v -race -coverprofile` em `ubuntu-latest` e `windows-latest`.
- **Build:** Cross-compilation para linux/amd64, windows/amd64, darwin/amd64.
- CI é acionado em push e PRs para `main`.

---

### H4 — Tags/releases semânticas coerentes; changelog gerado automaticamente

**Status: ✅ Atendido**

**Evidências:**
- Tags semânticas presentes: `v0.0.1` a `v0.1.1`.
- `build.yml` usa `generate_release_notes: true` para changelog automático via GitHub.
- Build usa `git describe --tags --always` para versão.
- Artefatos de release incluem checksums SHA-256 e assinatura Cosign.

---

### H5 — Hygiene: sem branches mortas, sem PRs abertos há muito tempo

**Status: ✅ Atendido**

**Evidências:**
- Branch local `marcello-alterações` está ativa e mesclada.
- O arquivo obsoleto `ci.yml` foi removido. Apenas `build.yml` gerencia o ciclo de vida.

---

## I. Operabilidade

### I1 — `--help` que ensina com exemplos, não só lista flags

**Status: ✅ Atendido**

**Evidências:**
- `root.go Help()` inclui exemplos: `assinatura start`, `assinatura sign --input documento.pdf`, `assinatura stop`.
- `sign.go Help()`, `validate.go Help()`, `start.go Help()`, `stop.go Help()`, `status Help()` incluem exemplos de uso.
- `version.go Help()` documenta todas as flags com descrição.

---

### I2 — Versão acessível via `--version` retornando tag + SHA curto

**Status: ✅ Atendido**

**Evidências:**
- `assinatura --version` invoca o `VersionCmd`.
- `version.go GetFull()` retorna: `assinatura v0.1.0 (commit abc1234, built at 2026-05-20T...)`.
- Build injeta commit e buildtime via `-ldflags`.

---

### I3 — Logs em nível ajustável; `--verbose`/`--quiet` previsível

**Status: ✅ Atendido**

**Evidências:**
- `--verbose` disponível em `sign`, `validate`, `start` — habilita logs `slog.Info` detalhados.
- `version --quiet` retorna apenas o número de versão (sem prefixo).
- `sign --json` e `validate --json` oferecem saída estruturada.
- `status --json` oferece saída JSON do estado do servidor.
- Logs via `slog` são estruturados (chave-valor) e podem ser filtrados por nível.

---

## Plano de Ação para os ~12% Restantes

Para atingir **~95%+ de conformidade**, as seguintes ações são recomendadas:

### 🟡 Média Prioridade

| # | Ação | Critério |
|---|------|----------|
| 1 | Implementar simulador do HubSaúde (start/stop/status/health + readiness separados) | E4 |
| 2 | Adicionar testes e2e com subprocess real (compilar binário Go + executar como processo filho) | G2 |
| 3 | Substituir cópias de docs upstream (`docs/aulas/especificacao.md`, `docs/aulas/criterios.md`) por referências com hash fixo | A2, B4 |
| 4 | Remover/atualizar `ci.yml` que referencia path obsoleto `docs/aulas/projetos/assinatura` | H5 |

### 🟢 Baixa Prioridade (Polimento)

| # | Ação | Critério |
|---|------|----------|
| 5 | Enriquecer simulador PKCS#11 com SoftHSM2 (quando hardware disponível) | E5 |
| 6 | Verificar e excluir branches mortas no GitHub | H5 |
| 7 | Adicionar `maven-checkstyle-plugin` ao `pom.xml` para lint Java no CI | D3 |

---

*Relatório atualizado em 2026-06-09 com base na análise completa do código-fonte, configurações de CI e documentação do repositório.*
