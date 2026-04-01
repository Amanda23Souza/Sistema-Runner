# Sistema-Runner
Repositório dedicado para desenvolvimento do projeto Runner, da disciplina implementação e integração de software.

## Status atual do módulo `cli-assinatura`

### ✅ Implementado
- Estrutura de projeto Go em `cli-assinatura`.
- `go.mod` configurado para `github.com/Amanda23Souza/Sistema-Runner/cli-assinatura`.
- `main.go` atualizado para usar o módulo `cmd`:
  - entrada via `cmd.Execute()`.
- Comando `version` funcional (valor padrão `v0.1.0`, overridable via `-ldflags`).
- Comando `help` funcional com lista de comandos e uso básico.
- Stub básico para `sign` e `validate` (placeholder com argumentos).
- Tests em `cmd/root_test.go`:
  - `TestExecute_Version`
  - `TestExecute_Help`
  - `TestExecute_UnknownCommand`

### ⚠️ O que falta
- Implementar parser de flags para `sign`/`validate` (ex.: `--input`, `--output`, `--sig`).
- Implementar integração com `assinador.jar` via `java -jar`.
- Implementar descoberta/caching de JDK 21 (`~/.hubsaude/jdk`).
- Implementar payload/response estruturado e validações (US-02.1/US-02.2/US-02.3).
- Implementar testes de integração real (CLI -> assinar.jar) (Sprint 2).
- `go test` não executou aqui porque ambiente não tem Go instalado (`go` missing). Precisa rodar localmente com Go 1.26+.

## Como testar localmente

1. Instalar Go 1.26+ (padrão `go` disponível no PATH):
   - `go version` deve funcionar.
   - `which go` não pode retornar vazio.
2. Executar os comandos:
   - `cd cli-assinatura`
   - `go test ./...`
   - `go run . version`
   - `go run . help`
   - `go run . sign --input foo` (output placeholder por enquanto)

## Planejamento Sprint 1 (US-01.1 / US-05.1 / US-05.2)

- [x] Inicializar `go mod`.
- [x] Estrutura de pacotes (`cmd` + entrypoint). 
- [x] Comando `version` e `help`.
- [ ] Cross-compile e pipeline CI/CD (falta `.github/workflows` em repo).
- [ ] Release SemVer + artifact naming.

## Planejamento Sprint 2 (Fluxo local)

- [ ] `sign`/`validate` com parser e chamada real a `assinador.jar` (`java -jar`).
- [ ] `FakeSignatureService` (Java). 
- [ ] Validação de parâmetros. 
- [ ] Saída legível e mensagens de erro.
- [ ] Testes de integração ponta a ponta.

## O que você precisa fazer agora

1. Garanta que Go esteja instalado no seu ambiente.
2. Rode `go test ./...` em `cli-assinatura`.
3. Prossiga com implementação de parser/execução real de `assinar.jar`.
4. Refine o backlog em `issues`/`projects` usando os US já mapeados no `plano-revisitado-v2.md`.

---

## Comandos disponíveis por enquanto

- `assinatura version`  -> Exibe versão do CLI.
- `assinatura help`    -> Exibe texto de ajuda.
- `assinatura sign ...` -> Mock placeholder (ainda em implementação).
- `assinatura validate ...` -> Mock placeholder (ainda em implementação).
