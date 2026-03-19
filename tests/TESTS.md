# Planejamento de Testes (detalhes)

Objetivo: usar Cucumber/Gherkin para testes de aceitação e contrato (BDD), mantendo testes unitários com JUnit/Mockito.

- **Estrutura recomendada:**
	- `src/test/java/.../steps/` → step definitions (Java).
	- `src/test/resources/features/` → arquivos `.feature` em Gherkin.
	- `tests/fixtures/` → payloads JSON de entrada/saída.

- **Divisão de tarefas (paralela):**
	- Integrante A: escrever features e steps para fluxos CLI (criar/validar).
	- Integrante B: escrever features e steps para integração HTTP e cenários de erro.

Unitários
- CLI: JUnit 5 + Mockito para parsing, validação e formatação (responsabilidade do Integrante A).
- `assinador.jar`: JUnit 5 para validação de parâmetros e regras (responsabilidade do Integrante B).

BDD / Cucumber (aceitação)
- Escopo: cenários de aceitação dos US-01 e US-02 descritos em Gherkin; cada cenário mapeia passos que podem invocar a CLI (processo) ou enviar requisições HTTP.
- Exemplo de feature (`src/test/resources/features/criar_assinatura.feature`):

```gherkin
Feature: Criar assinatura
    Scenario: Criar assinatura válida
        Given um pedido de assinatura válido em "tests/fixtures/pedido.json"
        When eu executo o comando "assinatura create --input tests/fixtures/pedido.json"
        Then a saída contém "signature"
```

- Step definitions podem usar ProcessBuilder para executar o CLI (`java -jar`), ou HttpClient para chamadas HTTP; use assertions do JUnit.

Execução
- Com Maven: `mvn test` (configurar `cucumber-junit`/`cucumber-junit-platform-engine`).
- Com Gradle: `./gradlew test` (configurar plugin Cucumber).

Integração local / scripts
- Mantenha um script `scripts/test-e2e.sh` que:
	1. inicia `assinador.jar` em modo servidor (se necessário),
	2. executa `mvn test` (ou `./gradlew test`) para rodar features Cucumber,
	3. encerra o servidor e retorna código de erro apropriado.

Notas práticas
- Prefira pequenas fixtures e cenários rápidos para CI; mova casos pesados para um job separado.
- Configure o CI para rodar os testes BDD usando um runner com Java 21 e as dependências Cucumber.

- Integrante A: escrever testes para o parser (`args -> DTO`) e validação de parâmetros (cenários válidos/inválidos).
- Integrante B: escrever testes para a formatação de saída e tratamento de erros (mocks para invocação do `assinador`).

Integração HTTP
- Escopo: iniciar `assinador.jar` em modo servidor e enviar requisições HTTP (criar/validar).
- Ferramentas: `curl` para testes manuais; Rest-assured ou HttpClient para testes automatizados.

Aceitação / Testes de contrato
- Mapear casos de aceitação dos US-01 e US-02 em testes automatizados ou scripts de verificação (checklist executável).
- Exemplo de checagem simples (pass/fail):
	- Invocar `assinatura create` com input válido → saída contém campo `signature`.
	- Invocar `assinatura validate` com assinatura simulada → retorno `valid: true`.

Notas práticas finais
- Use fixtures (`tests/fixtures/`) com exemplos pequenos e casos de erro para acelerar execução.
- Integração: inclua um `scripts/test-e2e.sh` que roda os cenários de integração local e retorna código de saída não-zero em falha.
- CI: configure job que execute `mvn test`/`./gradlew test` e o script de integração em runners com Java 21.
