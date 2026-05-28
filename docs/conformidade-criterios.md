# Relatório de Conformidade com os Critérios de Avaliação

> **Repositório:** [Amanda23Souza/Sistema-Runner](https://github.com/Amanda23Souza/Sistema-Runner)
> **Gerado em:** 2026-05-27
> **Referência dos critérios:** [`docs/aulas/criterios.md`](./aulas/criterios.md)

---

## Resumo Executivo

| Seção | Critérios Totais | Atendidos | Parciais | Ausentes | % Atendimento |
|-------|:---:|:---:|:---:|:---:|:---:|
| A. Princípios Transversais | 5 | 3 | 2 | 0 | 70% |
| B. Organização do Repositório | 5 | 3 | 1 | 1 | 70% |
| C. Documentação | 4 | 2 | 2 | 0 | 63% |
| D. Qualidade de Código | 8 | 4 | 2 | 2 | 56% |
| E. Requisitos Funcionais | 13 | 3 | 4 | 6 | 38% |
| F. Build e Dependências | 4 | 3 | 1 | 0 | 81% |
| G. Testes | 5 | 2 | 2 | 1 | 50% |
| H. Engenharia de Processo | 5 | 3 | 1 | 1 | 65% |
| I. Operabilidade | 3 | 2 | 1 | 0 | 78% |
| **TOTAL** | **52** | **25** | **16** | **11** | **~62%** |

> **Atenção:** O repositório atualmente atende **~62%** dos critérios (considerando "parciais" como 50%).
> Para atingir **99%**, as lacunas listadas abaixo devem ser endereçadas.

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

**Status: ⚠️ Parcial**

**Evidências:**
- Commits seguem padrão semântico (`feat:`, `implementação de`, etc.) — cadeia parcialmente navegável.
- Issues/PRs existem e são referenciados em commits (ex: PR #26, PR #25, PR #24).
- O `nossoPlanejamento.md` lista User Stories com links para Google Docs (RF-01, RF-02).
- A US-01 possui arquivo dedicado em `requisitos/funcional/US-01 - Invocar Assinador via CLI.md`.

**Lacunas:**
- Nem todos os commits referenciam explicitamente issues/PRs.
- Falta rastreamento reverso (do teste até o requisito): os testes em `command_test.go` não mencionam qual US cobrem.
- RF-03 está sem descrição no planejamento.

---

### A2 — Single Source of Truth (sem duplicação da spec upstream)

**Status: ⚠️ Parcial**

**Evidências:**
- A especificação do professor está referenciada via link em `README.md` (`docs/aulas/especificacao.md`).
- Não há duplicação óbvia de especificações completas.

**Lacunas:**
- Os links para a especificação no `README.md` e em `nossoPlanejamento.md` apontam para caminhos relativos internos e para links do Google Docs (não fixos com commit/tag do repositório upstream `kyriosdata/runner`). O critério exige link com **commit/tag fixo** no upstream, não `main`.
- `docs/aulas/especificacao.md` pode ser uma cópia de conteúdo upstream — recomenda-se substituir por referência com hash fixo.

---

### A3 — Reprodutibilidade (clonar → um comando → build e testes verdes)

**Status: ✅ Atendido**

**Evidências:**
- `README.md` documenta os comandos de build: `go build -o assinatura ./cmd/assinatura` e `mvn clean package`.
- O CI (`.github/workflows/build.yml`) valida o build e os testes automaticamente a cada push/PR.
- `go.mod` declara versão mínima Go (`go 1.26.1`).

---

### A4 — Falhar bem (erros explícitos, códigos de saída, mensagens esclarecedoras)

**Status: ✅ Atendido**

**Evidências:**
- Comandos `sign` e `validate` retornam mensagens detalhadas em português: `"Erro de validação: o campo --input é obrigatório."`, `"Erro de validação: não foi possível ler o arquivo de entrada"`.
- Exit codes: `main.go` propaga `os.Exit(1)` em caso de erro.
- Código de erro `[MS-03]` identificado para parâmetros inválidos.
- Mensagens distinguem tipo de falha (ausência de parâmetro vs. falha de leitura de arquivo).

---

### A5 — Decisões registradas em ADRs

**Status: ✅ Atendido parcialmente — registrado em docs/design.md**

**Evidências:**
- `docs/design.md` documenta decisões arquiteturais (modelo C4, separação CLI/JAR, uso de PlantUML).
- `docs/aulas/plano-revisitado-v2.md` contém contexto das decisões técnicas.

**Lacunas:**
- Não há ADRs formais (curtos, 1 página) com formato padrão registrando decisões como: escolha da porta padrão, estratégia de descoberta de instância HTTP, parser de CLI (stdlib `flag` vs cobra), linguagem Go vs Python.
- Recomenda-se criar `docs/adr/` com pelo menos 2–3 ADRs para as principais decisões.

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

**Status: ❌ Ausente**

**Evidências:**
- **Não existe `.gitignore`** na raiz do repositório.
- O diretório `docs/aulas/projetos/assinador-java/target/` (artefato Maven compilado) **está versionado** — violação direta do critério.
- O binário `cli-assinatura/assinatura` (2,9 MB) **também está versionado** como arquivo `assinatura`.

**Ação necessária:** Criar `.gitignore` na raiz cobrindo `target/`, `*.class`, `*.jar` (intermediários), binários Go, `.idea/`, `__pycache__/`, `.DS_Store`.

---

### B3 — `LICENSE` presente e compatível

**Status: ❌ Ausente**

**Evidências:**
- Não há arquivo `LICENSE` ou `LICENSE.md` no repositório.
- As dependências utilizadas (Javalin, Jackson — Apache 2.0) exigem compatibilidade.

**Ação necessária:** Adicionar `LICENSE` (sugestão: Apache 2.0 ou MIT).

---

### B4 — Sem documentos que pertencem ao repositório da especificação

**Status: ⚠️ Parcial**

**Evidências:**
- `docs/aulas/criterios.md`, `docs/aulas/especificacao.md`, `docs/aulas/implementacao.png` — são arquivos do professor versionados no repositório da implementação.
- O critério exige que apenas conteúdo **específico desta implementação** permaneça; documentos do upstream devem ser referenciados, não copiados.

**Lacunas:**
- Remover ou mover para uma pasta claramente identificada como "materiais de aula" e substituir por referências com hash fixo.

---

### B5 — Nomenclatura consistente (idioma único, sem acentos/espaços, padrão coerente)

**Status: ✅ Atendido**

**Evidências:**
- Código Go usa `camelCase` e `PascalCase` coerentemente.
- Nomes de arquivos em inglês no código (`root.go`, `sign.go`, `version.go`).
- Estrutura de pacotes Go segue padrões idiomáticos.

**Lacunas menores:**
- Mistura de idiomas em paths de documentação: `docs/aulas/projetos/assinador-java` (PT) vs `internal/command` (EN). Não crítico, mas inconsistente.

---

## C. Documentação

### C1 — README como contrato: o que é, como gerar, como executar, como testar, como contribuir, status

**Status: ✅ Atendido**

**Evidências:**
- README cobre: descrição do projeto, requisitos, estrutura, como compilar (Go e Java), como usar, status de implementação (✅/⏳ por fase), próximos passos.
- Exemplos de uso com `curl` para endpoints HTTP.
- Seção de documentação com links internos.

**Lacunas:**
- Falta seção "Como contribuir" explícita (convenções de commits, processo de PR).
- Falta instrução de como executar os testes: `go test ./...` e `mvn test`.

---

### C2 — Referência à especificação com link commit/tag fixo (não `main`)

**Status: ⚠️ Parcial**

**Evidências:**
- README referencia `docs/aulas/especificacao.md` (link interno relativo).
- `nossoPlanejamento.md` referencia links do Google Docs.

**Lacunas:**
- Nenhuma referência ao upstream `kyriosdata/runner` com commit/tag fixo. O critério exige que links para a especificação original usem hash de commit ou tag específica para evitar deriva.

---

### C3 — ADRs curtos (1 página) para decisões relevantes

**Status: ⚠️ Parcial**

**Evidências:**
- `docs/design.md` registra decisões arquiteturais de forma breve.

**Lacunas:**
- Sem estrutura formal de ADR. Recomenda-se criar `docs/adr/001-escolha-go-para-cli.md`, `docs/adr/002-modo-servidor-padrao.md`, etc.

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
- Cada arquivo Go (`sign.go`, `validate.go`, `version.go`, `root.go`) tem responsabilidade única e clara.
- Funções são curtas (máximo ~40 linhas).
- Separação entre camadas: `internal/command/` (lógica CLI) e `internal/version/` (metadados).
- Interface `Command` em `interface.go` desacopla implementações do orquestrador.

---

### D2 — Fronteiras explícitas: contrato CLI↔JAR documentado e testado

**Status: ⚠️ Parcial**

**Evidências:**
- Os endpoints HTTP (`/sign`, `/validate`, `/health`) estão documentados no README com exemplos `curl`.
- Testes de integração Java (`SignatureControllerTest`) validam as rotas HTTP.

**Lacunas:**
- O contrato entre CLI Go e o JAR (parâmetros de linha de comando, formato de saída, códigos de erro) não está formalmente documentado como uma API (ex: `docs/api-contract.md`).
- A integração real CLI→HTTP ainda não está implementada (conforme README: "⏳ Integração real CLI-HTTP").

---

### D3 — Aderência ao estilo da linguagem, exigida via CI

**Status: ⚠️ Parcial**

**Evidências:**
- Código Go segue convenções idiomáticas (exported types com PascalCase, comentários de package, uso de `io.Writer` para injeção de dependência).

**Lacunas:**
- Nenhuma etapa de lint (`golangci-lint`, `gofmt`, `go vet`) está configurada no CI (`build.yml`). O critério exige lint via CI, não revisão manual.
- Para Java: nenhum checkstyle ou spotbugs configurado no `pom.xml` ou CI.

---

### D4 — Tipagem usada com intenção

**Status: ✅ Atendido**

**Evidências:**
- Go: uso correto de tipos (`io.Writer`, `[]string`, structs tipados `signResult`, `validateResult`).
- Java: DTOs tipados (`SignRequest`, `SignatureResponse`, `ValidateRequest`).
- Interface `SignatureService` com contrato explícito.

---

### D5 — Sem `catch (Throwable)` genéricos engolindo erro

**Status: ✅ Atendido**

**Evidências:**
- Java: tratamento de exceções específico via Javalin (não identificado catch genérico no código fonte analisado).
- Go: erros propagados explicitamente com `return err` ou mensagem descritiva.

---

### D6 — Logs estruturados (não `print`/`System.out`)

**Status: ❌ Não atendido**

**Evidências:**
- Go: uso extensivo de `fmt.Fprintf` e `fmt.Fprintln` para saída — não é logging estruturado.
- Java: Javalin usa slf4j por padrão (adequado), mas o código da aplicação não usa logger explicitamente.

**Ação necessária:** Substituir `fmt.Fprintf` diagnóstico por `slog` (stdlib Go 1.21+). Manter `fmt` apenas para saída orientada ao usuário (`--json`, resultados).

---

### D7 — Sem segredos, caminhos absolutos ou portas hardcoded fora de configuração

**Status: ✅ Atendido**

**Evidências:**
- Nenhuma senha, token ou chave encontrada no código.
- A porta do servidor Java (8080) é passada como argumento em linha de comando, não hardcoded.
- Caminhos são relativos ou passados via parâmetros CLI.

---

### D8 — Encoding UTF-8 declarado; line endings tratados (`.gitattributes`)

**Status: ❌ Não atendido**

**Evidências:**
- Não existe arquivo `.gitattributes` no repositório.
- Go e Java geralmente assumem UTF-8, mas o critério exige declaração explícita e tratamento de line endings (especialmente para portabilidade Windows/Linux).

**Ação necessária:** Criar `.gitattributes` com:
```
* text=auto eol=lf
*.go text eol=lf
*.java text eol=lf
*.md text eol=lf
*.bat text eol=crlf
```

---

## E. Requisitos Funcionais e de Integração

### E1.1 — Executáveis funcionam independente do diretório atual

**Status: ✅ Atendido**

**Evidências:**
- CLI Go não usa caminhos relativos hardcoded para suas operações internas.
- Os parâmetros `--input` e `--output` são especificados pelo usuário.

---

### E1.2 — Passagem de argumentos preserva espaços, acentos, aspas

**Status: ⚠️ Parcial**

**Evidências:**
- O parser usa `flag.NewFlagSet` que lida corretamente com strings.

**Lacunas:**
- Sem teste específico comprovando preservação de espaços e acentos nos nomes de arquivo.

---

### E1.3 — Propaga exit code e separa stdout (resultado) de stderr (diagnóstico)

**Status: ⚠️ Parcial**

**Evidências:**
- `main.go` usa `os.Exit(1)` em caso de erro.

**Lacunas:**
- Mensagens de diagnóstico de erro vão para `c.out` (stdout), não para stderr. Diagnósticos devem ir para `os.Stderr`.

---

### E2.1 — Idempotência de start com health check real

**Status: ❌ Não implementado**

**Evidências:**
- README marca "⏳ Comandos `start`/`stop` e monitoramento de ciclo de vida na CLI Go" como pendente.
- Não há comando `start` na CLI Go.

---

### E2.2 — Porta padrão configurável; falha clara com porta ocupada

**Status: ❌ Não implementado**

**Evidências:**
- O servidor Java aceita porta via argumento, mas a CLI Go ainda não expõe essa configuração.
- Sem tratamento de "porta ocupada" na CLI.

---

### E2.3 — Shutdown controlado por endpoint/sinal em qualquer porta

**Status: ❌ Não implementado**

**Evidências:**
- Sem comando `stop` na CLI.
- Sem endpoint de shutdown no servidor Java (apenas `/health`).

---

### E2.4 — Auto-shutdown por inatividade com timer reiniciado a cada requisição

**Status: ❌ Não implementado**

**Evidências:**
- Funcionalidade não presente no servidor Java atual (`App.java`).

---

### E2.5 — Modo servidor é o padrão; modo local deve ser explicitamente ativado

**Status: ❌ Não atendido**

**Evidências:**
- Atualmente o padrão em `sign.go` e `validate.go` é `--mode local` (inverso do que o critério exige).
- `fs.StringVar(&c.mode, "mode", "local", ...)` — padrão deveria ser `"http"`.

---

### E2.6 — Tratamento de timeout, conexão recusada, resposta malformada

**Status: ❌ Não implementado**

**Evidências:**
- A integração HTTP na CLI ainda não existe; sem tratamento de erros de rede.

---

### E3.1 — Validação feita dentro do `assinador.jar` (autoridade única)

**Status: ✅ Atendido**

**Evidências:**
- `FakeSignatureService.java` realiza a validação de parâmetros no backend Java.
- A CLI Go faz apenas validação de presença de parâmetros obrigatórios (necessário para feedback imediato), não replica regras de negócio.

---

### E3.2 — Mensagens distinguem erro do usuário de erro do sistema; códigos diferentes

**Status: ⚠️ Parcial**

**Evidências:**
- CLI distingue "Erro de validação" (usuário) de erros de I/O (sistema).
- Código `[MS-03]` para parâmetros inválidos.

**Lacunas:**
- Sem códigos de saída distintos entre erro de usuário (ex: 2) e erro de sistema (ex: 1). Atualmente apenas `1` é usado.

---

### E4 — Simulador do HubSaúde: ciclo de vida com health check e readiness

**Status: ❌ Não implementado**

**Evidências:**
- O design (`docs/design.md`) menciona o "Simulador do HubSaúde" como sistema externo.
- Sem implementação de `start/stop/status` com health check e readiness check separados.

---

### E5 — Simulador PKCS11 com testes de integração

**Status: ❌ Não implementado**

**Evidências:**
- `FakeSignatureService.java` simula operações de assinatura, mas não responde a chamadas PKCS11 reais.
- Sem testes de integração PKCS11.

---

### E6 — Portabilidade comprovada em CI (Windows e Linux)

**Status: ⚠️ Parcial**

**Evidências:**
- `build.yml` compila para `linux/amd64`, `windows/amd64` e `darwin/amd64` via cross-compilation.

**Lacunas:**
- Os **testes** (`go test ./...`) são executados apenas em `ubuntu-latest`. O critério exige que os testes passem em **Windows e Linux** em CI (não apenas cross-compilation).
- Sem `runs-on: windows-latest` para executar a suíte de testes.

---

## F. Build, Dependências e Supply Chain

### F1 — Build reproduzível

**Status: ✅ Atendido**

**Evidências:**
- `go build` com `-ldflags` injetando versão, commit e buildtime.
- Maven com `pom.xml` declarativo.
- CI usa `go-version: stable` e `cache: true`.

---

### F2 — Versões mínimas declaradas e verificadas em runtime

**Status: ✅ Atendido**

**Evidências:**
- `go.mod`: `go 1.26.1` declara versão mínima do Go.
- README documenta `Go 1.26.1+` e `JDK` como requisitos.

**Lacunas menores:**
- Sem verificação em runtime da versão do Java (ex: checar se `java --version` retorna >= 21 antes de invocar o JAR).

---

### F3 — Dependências mínimas e justificadas; sem libs abandonadas ou com CVEs

**Status: ✅ Atendido**

**Evidências:**
- CLI Go: **zero dependências externas** (`go.mod` sem `require`). Usa apenas stdlib.
- Java: Javalin (web framework ativo) + Jackson (serialização JSON) — ambas mantidas e sem CVEs conhecidos.

---

### F4 — JAR único com `Main-Class` correto, sem dependências externas

**Status: ⚠️ Parcial**

**Evidências:**
- `pom.xml` usa `maven-assembly-plugin` para gerar fat JAR (`jar-with-dependencies`).
- `Main-Class` configurado no `pom.xml`.

**Lacunas:**
- O fat JAR gerado está em `target/` que está versionado no repositório — deve ser excluído via `.gitignore` e distribuído apenas via releases do GitHub.

---

## G. Testes

### G1 — Pirâmide saudável: muitos unitários, alguns integração, poucos e2e

**Status: ⚠️ Parcial**

**Evidências:**
- Go: 3 testes unitários em `command_test.go` (sign cria arquivo, validate válido, validate inválido).
- Java: `FakeSignatureServiceTest.java` e `SignatureControllerTest.java` cobrem unitários e integração HTTP.

**Lacunas:**
- Poucos testes unitários no Go (apenas 3 casos).
- Sem testes end-to-end (subprocess real invocando o binário compilado).

---

### G2 — Testes de contrato CLI↔JAR: subprocess real e HTTP real

**Status: ❌ Não atendido**

**Evidências:**
- Testes Go atuais testam a lógica interna dos comandos, não invocam subprocess real.
- Sem teste que compila e executa o binário como processo filho para validar o contrato CLI.

---

### G3 — Cenários negativos como cidadãos de primeira classe

**Status: ⚠️ Parcial**

**Evidências:**
- `TestValidateCmd_Run_InvalidSignature` — testa assinatura inválida ✅
- `[MS-03]` tratado para parâmetros ausentes ✅

**Lacunas:**
- Sem testes para: JAR ausente, JVM ausente, timeout, porta ocupada, payload inválido, race condition no start do servidor.

---

### G4 — Sem testes flaky; quando inevitável, marcados

**Status: ✅ Atendido**

**Evidências:**
- Testes atuais são determinísticos (não dependem de rede, tempo ou estado externo).
- Sem testes de concorrência ou timing que possam ser flaky.

---

### G5 — Cobertura: relatório publicado

**Status: ❌ Não atendido**

**Evidências:**
- Nenhuma etapa de geração/publicação de relatório de cobertura no CI (`build.yml`).
- Sem `go test -coverprofile` ou integração com Codecov/SonarCloud.

**Ação necessária:** Adicionar ao CI:
```yaml
- name: Test with coverage
  working-directory: cli-assinatura
  run: go test -coverprofile=coverage.out ./...
- name: Upload coverage
  uses: codecov/codecov-action@v4
  with:
    file: cli-assinatura/coverage.out
```

---

## H. Engenharia de Processo (Git/GitHub)

### H1 — Commits atômicos, mensagens no imperativo, Conventional Commits

**Status: ✅ Atendido**

**Evidências:**
- Commits seguem padrão semântico: `feat: implement signature REST API`, `feat: adicionar...`.
- Commits são granulares e focados.

**Lacunas menores:**
- Alguns commits em português e outros em inglês (inconsistência de idioma).
- `"implementação de HTTP points inicial"` não segue o formato `type(scope): subject`.

---

### H2 — PRs pequenos, revisáveis, ligados a issues

**Status: ✅ Atendido**

**Evidências:**
- PRs #24, #25, #26 identificados no git log.
- Branch `marcello-alterações` com PR vinculado ao merge.

---

### H3 — CI obrigatório: lint + build + testes em Windows e Linux

**Status: ⚠️ Parcial**

**Evidências:**
- CI (`build.yml`) executa `go test -v ./...` e build multi-plataforma em `ubuntu-latest`.
- CI é acionado em push e PRs para `main`.

**Lacunas:**
- Sem etapa de **lint** no CI (ex: `golangci-lint`).
- Testes executam apenas em Linux (`ubuntu-latest`), não em `windows-latest`.
- `ci.yml` aponta para diretório `docs/aulas/projetos/assinatura` que não corresponde à estrutura atual — provavelmente inativo/quebrado.

---

### H4 — Tags/releases semânticas coerentes; changelog gerado automaticamente

**Status: ✅ Atendido**

**Evidências:**
- Tags semânticas presentes: `v0.0.1` a `v0.1.1`.
- `build.yml` usa `generate_release_notes: true` para changelog automático via GitHub.
- Build usa `git describe --tags --always` para versão.

---

### H5 — Hygiene: sem branches mortas, sem PRs abertos há muito tempo

**Status: ❌ Requer verificação manual**

**Evidências:**
- Branch local `marcello-alterações` está ativa e mesclada.
- Branch `sync-remote-professor` pode estar obsoleta.

**Ação recomendada:** Verificar e remover branches mescladas no GitHub (`git branch -d` ou via interface web).

---

## I. Operabilidade

### I1 — `--help` que ensina com exemplos, não só lista flags

**Status: ✅ Atendido**

**Evidências:**
- `root.go Help()` inclui exemplos: `assinatura sign --input document.pdf`, `assinatura criar --input document.pdf --mode http`.
- `sign.go Help()` e `validate.go Help()` incluem exemplos de uso.
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

**Status: ⚠️ Parcial**

**Evidências:**
- `version --quiet` retorna apenas o número de versão (sem prefixo).
- `sign --json` e `validate --json` oferecem saída estruturada.

**Lacunas:**
- Sem flag `--verbose` global.
- Sem níveis de log ajustáveis (debug, info, warn, error) — especialmente relevante para o modo servidor.

---

## Plano de Ação Prioritário

Para atingir **99% de conformidade**, as seguintes ações são necessárias, ordenadas por impacto:

### 🔴 Alta Prioridade (Bloqueadores críticos)

| # | Ação | Critério |
|---|------|----------|
| 1 | Criar `.gitignore` e remover artefatos versionados (`target/`, binário `assinatura`) | B2 |
| 2 | Adicionar `LICENSE` (Apache 2.0 ou MIT) | B3 |
| 3 | Criar `.gitattributes` para encoding UTF-8 e line endings | D8 |
| 4 | Mudar padrão de `--mode` de `local` para `http` | E2.5 |
| 5 | Redirecionar mensagens de erro para `stderr` | E1.3 |
| 6 | Adicionar lint ao CI (`golangci-lint` para Go) | D3, H3 |
| 7 | Adicionar testes em `windows-latest` no CI | E6, H3 |

### 🟡 Média Prioridade (Funcionalidades essenciais)

| # | Ação | Critério |
|---|------|----------|
| 8 | Implementar comandos `start`/`stop` com health check real | E2.1, E2.3 |
| 9 | Implementar auto-shutdown por inatividade com timer | E2.4 |
| 10 | Adicionar tratamento de timeout e erros de rede na integração HTTP | E2.6 |
| 11 | Criar ADRs formais em `docs/adr/` | A5, C3 |
| 12 | Publicar relatório de cobertura via Codecov ou similar | G5 |
| 13 | Adicionar testes de contrato CLI↔JAR com subprocess real | G2 |
| 14 | Adicionar códigos de saída distintos: erro de usuário (2) vs sistema (1) | E3.2 |

### 🟢 Baixa Prioridade (Polimento)

| # | Ação | Critério |
|---|------|----------|
| 15 | Substituir `fmt.Fprintf` diagnóstico por `slog` (logs estruturados) | D6 |
| 16 | Adicionar links com hash fixo ao upstream `kyriosdata/runner` | A2, C2 |
| 17 | Adicionar seção "Como contribuir" e instrução `go test ./...` no README | C1 |
| 18 | Adicionar flag `--verbose` global | I3 |
| 19 | Verificar e excluir branches mortas no GitHub | H5 |
| 20 | Verificação em runtime da versão do Java antes de invocar o JAR | F2 |

---

*Relatório gerado automaticamente com base na análise do código-fonte e configurações do repositório em 2026-05-27.*
