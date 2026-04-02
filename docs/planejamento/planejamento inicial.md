## Planejamento da Iteração

## Integrantes

| Papel | Nome | GitHub |
|---|---|---|
| - | Marcello Ronald | https://github.com/mronald-js |
| - | Amanda Soares  |  |  |

### 1. Contexto da Iteração

| Campo | Descrição |
|---|---|
| Iteração | Iteração 1 - MVP de Assinatura (foco em [US-01](requisitos/funcional/US-01-invocar-assinador-via-cli.md) e [US-02](requisitos/funcional/US-02-simular-assinatura-validacao-parametros.md)) |
| Estimativa (início/fim) | 18/03/2026 a - (a definir) |
| Objetivo principal | Facilitar a execução de aplicações Java via CLI, ocultando complexidade de configuração do ambiente Java, com foco na integração entre assinatura e assinador.jar em modo local e HTTP (ver [Catálogo de Requisitos](requisitos/README.md) e [US-01](requisitos/funcional/US-01-invocar-assinador-via-cli.md)). |
| Entregas esperadas | 1 - CLI assinatura com comandos criar/validar ([US-01](requisitos/funcional/US-01-invocar-assinador-via-cli.md)); <br> 2 - assinador.jar com validação de parâmetros e simulação ([US-02](requisitos/funcional/US-02-simular-assinatura-validacao-parametros.md)); <br> 3 - integração assinatura e assinador.jar ([US-01](requisitos/funcional/US-01-invocar-assinador-via-cli.md)); <br> 4 - testes iniciais e documentação ([Catálogo de Requisitos](requisitos/README.md)). |

Referências complementares: [US-03](requisitos/funcional/US-03-gerenciar-simulador-hubsaude.md), [US-04](requisitos/funcional/US-04-provisionar-jdk-automaticamente.md), [US-05](requisitos/funcional/US-05-disponibilizar-binarios-multiplataforma.md), [Regras de negócio gerais](requisitos/funcional/regras-negocio-gerais.md), [Requisitos não funcionais](requisitos/nao-funcional/RNF-01-desempenho.md).

### 2. Design

Abaixo está a lista de artefatos de design detalhado e seus locais no repositório. Cada item deve conter modelos e descrições suficientes para guiar implementação e testes.

| Artefato | Local | Descrição | Status |
|---|---|---|---|
| Modelos de comportamento (diagramas de sequência / casos de uso) | [design/comportamento](design/comportamento/README.md) | Sequências para fluxos principais (criar/validar assinatura, modos local e HTTP) e cenários de erro. | Em andamento |
| Estrutura (componentes / deployment) | [design/estrutura](design/estrutura/README.md) | Componentes, responsabilidades e layout de deployment (CLI, `assinador.jar`, PKCS#11, JDK). | Em andamento |
| Contratos / API (OpenAPI, exemplos) | [design/contratos](design/contratos/README.md) | Especificação dos endpoints HTTP, payloads e exemplos de request/response; formato de erros. | Em andamento |
| Modelo de dados (schemas) | [design/dados](design/dados/README.md) | Schemas para entradas/saídas (JSON Schema), mapeamento de campos e exemplos. | Em andamento |

Referência adicional: design amostral em [especificação/design.md](../docs/design.md).
### 3. Preparação de Ambiente e Convenções

| Tópico | Definição | Status |
|---|---|---|
| Linguagem e versão | CLI: Java 21 |  | Em andamento |
| Ferramentas de build | Java: `mvn test package` (ou `gradle`, confirmar na iteração) | a definir | Pendente |
| CI/CD | Pipeline com build + testes + artefatos de validação | Pendente |


#### Convenções de código

- **Commits semânticos:** use o padrão `type(scope): subject` (ex.: `feat(cli): adicionar comando criar`).
- **Revisão por par:** todo PR deve ter pelo menos uma revisão aprovada antes do merge.
- **Testes obrigatórios:** PRs que alterem lógica devem incluir testes unitários relevantes.

#### Estratégia de branches (Git Flow)

- **Branches principais:** `main` (produção), `dev` (integração contínua).
- **Fluxo:** Quando uma funcionalidade nova for integrada subir commit para `dev` a partir de uma branch `feat/{nome-da-feature}`, se for correção de bug fazer o mesmo processo mas como `bugfix/{bug-relacionado}`
- **Referência:** Inspirado e adaptado de [Git Flow — Atlassian](https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow)

### 4. Backlog da Iteração

será criado em projects no github

### 5. Planejamento de Testes

O projeto usará BDD com Cucumber para testes de aceitação e JUnit/Mockito para testes unitários.


### 6. Planejamento de Documentação

| Entregável | Conteúdo mínimo | Referência | Responsável | Status |
|---|---|---|---|---|
| Manual de uso CLI | Instalação, comandos, exemplos e erros comuns | [US-01](requisitos/funcional/US-01-invocar-assinador-via-cli.md) | a definir | Pendente |
| Doc técnica integração | Fluxos local/HTTP, contrato de entrada/saída | [US-01](requisitos/funcional/US-01-invocar-assinador-via-cli.md) | a definir | Pendente |
| Guia de testes | Como executar suíte unitária e integração | [Catálogo de Requisitos](requisitos/README.md) | a definir | Pendente |

### 7. Planejamento de Release e Distribuição

| Item | Definição | Referência | Status |
|---|---|---|---|
| SemVer | Tag no formato `vMAJOR.MINOR.PATCH` | [US-05](requisitos/funcional/US-05-disponibilizar-binarios-multiplataforma.md) | Pendente |
| Artefatos multiplataforma | Build para Windows, Linux e macOS (amd64) | [US-05](requisitos/funcional/US-05-disponibilizar-binarios-multiplataforma.md) | Pendente |
| Integridade | Publicação de checksums SHA256 | [US-05](requisitos/funcional/US-05-disponibilizar-binarios-multiplataforma.md) | Pendente |

### 8. Planejamento de Segurança de Artefatos (Cosign)

| Item | Definição | Referência | Status |
|---|---|---|---|
| Assinatura automática | CI/CD deve assinar cada artefato com Cosign | [RNF-03](requisitos/nao-funcional/RNF-03-seguranca.md) | Pendente |
| Arquivos obrigatórios | Publicar `<artefato>`, `<artefato>.sig`, `<artefato>.pem` | [RNF-03](requisitos/nao-funcional/RNF-03-seguranca.md) | Pendente |
| Verificação | Instrução com `cosign verify-blob` na release | [RNF-03](requisitos/nao-funcional/RNF-03-seguranca.md) | Pendente |

### 9. Próximas Iterações

| Iteração | Escopo | Requisitos alvo | Critério de conclusão |
|---|---|---|---|
| Iteração ? | Gestão do simulador e download condicional | [US-03](requisitos/funcional/US-03-gerenciar-simulador-hubsaude.md) | Comandos start/stop/status e controle de versão do simulador |
| Iteração ? | Provisionamento automático de JDK | [US-04](requisitos/funcional/US-04-provisionar-jdk-automaticamente.md) | Detecção, download e uso interno do JDK em 3 plataformas |
| Iteração ? | Release multiplataforma e assinatura | [US-05](requisitos/funcional/US-05-disponibilizar-binarios-multiplataforma.md), [RNF-03](requisitos/nao-funcional/RNF-03-seguranca.md) | Release com binários, checksums e assinatura válida |

### 10. Acompanhamento de RNFs

| RNF | Meta da iteração | Evidência | Status |
|---|---|---|---|
| [RNF-01](requisitos/nao-funcional/RNF-01-desempenho.md) | Medir baseline de tempo para fluxo local e HTTP | Registro de benchmark simples | Pendente |
| [RNF-04](requisitos/nao-funcional/RNF-04-confiabilidade-observabilidade.md) | Padronizar formato de erro e logs mínimos | Exemplos de log e erros em execução | Pendente |
| [RNF-05](requisitos/nao-funcional/RNF-05-manutenibilidade-testabilidade.md) | Garantir testes para US-01 e US-02 | Pipeline com testes automatizados | Pendente |
