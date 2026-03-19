# Sistema-Runner
Repositório dedicado para desenvolvimento do projeto Runner, da disciplina implementação e integração de software.

## Planejamento da Iteração

## Integrantes

| Papel | Nome | GitHub |
|---|---|---|
| - | Marcello Ronald | https://github.com/mronald-js |
| - | Amanda Soares  |  |  |

### 1. Contexto da Iteração

| Campo | Preenchimento |
|---|---|
| Iteração | Iteração 1 - MVP de Assinatura (foco em [US-01](requisitos/funcional/US-01-invocar-assinador-via-cli.md) e [US-02](requisitos/funcional/US-02-simular-assinatura-validacao-parametros.md)) |
| Estimativa (início/fim) | 18/03/2026 a - (a definir) |
| Objetivo principal | Facilitar a execução de aplicações Java via CLI, ocultando complexidade de configuração do ambiente Java, com foco na integração entre assinatura e assinador.jar em modo local e HTTP (ver [Catálogo de Requisitos](requisitos/README.md) e [US-01](requisitos/funcional/US-01-invocar-assinador-via-cli.md)). |
| Entregas esperadas | 1 - CLI assinatura com comandos criar/validar ([US-01](requisitos/funcional/US-01-invocar-assinador-via-cli.md)); <br> 2 - assinador.jar com validação de parâmetros e simulação ([US-02](requisitos/funcional/US-02-simular-assinatura-validacao-parametros.md)); <br> 3 - integração assinatura e assinador.jar ([US-01](requisitos/funcional/US-01-invocar-assinador-via-cli.md)); <br> 4 - testes iniciais e documentação ([Catálogo de Requisitos](requisitos/README.md)). |

Referências complementares: [US-03](requisitos/funcional/US-03-gerenciar-simulador-hubsaude.md), [US-04](requisitos/funcional/US-04-provisionar-jdk-automaticamente.md), [US-05](requisitos/funcional/US-05-disponibilizar-binarios-multiplataforma.md), [Regras de negócio gerais](requisitos/funcional/regras-negocio-gerais.md), [Requisitos não funcionais](requisitos/nao-funcional/RNF-01-desempenho.md).

### 2. Design

| Item | Definição | Referência | Responsável | Status |
|---|---|---|---|---|
| Modo de invocação do assinador | Definir e documentar os dois modos exigidos: invocação direta (CLI) e invocação via HTTP (servidor) | [US-01](requisitos/funcional/US-01-invocar-assinador-via-cli.md) | a definir | Em andamento |
| Fluxo de criação de assinatura | Especificar o fluxo ponta a ponta de criação (usuário -> assinatura -> assinador.jar -> assinatura -> usuário) | [US-01](requisitos/funcional/US-01-invocar-assinador-via-cli.md) | a definir | Em andamento |
| Fluxo de validação de assinatura | Especificar o fluxo ponta a ponta de validação (usuário -> assinatura -> assinador.jar -> assinatura -> usuário) | [US-02](requisitos/funcional/US-02-simular-assinatura-validacao-parametros.md) | a definir | Em andamento |
| Validação e simulação no assinador | Definir validações de parâmetros e comportamento de simulação para criar/validar | [US-02](requisitos/funcional/US-02-simular-assinatura-validacao-parametros.md) | a definir | Em andamento |
| Tratamento de erros | Definir erros estruturados e mensagens claras em qualquer ponto do fluxo | [Regras de negócio gerais](requisitos/funcional/regras-negocio-gerais.md) | a definir | Pendente |

### 3. Preparação de Ambiente e Convenções

| Tópico | Definição | Responsável | Status |
|---|---|---|---|
| Linguagem e versão | CLI: A definir. <br>Assinador: Java 21 | a definir | Em andamento |
| Ferramentas de build | Java: `mvn test package` (ou `gradle`, confirmar na iteração) | a definir | Pendente |
| Convenções de código | Commits semânticos, lint obrigatório e revisão do outro integrante antes de merge | a definir | Em andamento |
| Estratégia de branches | `main` protegida + branches `feature/*` + PR obrigatório | a definir | Em andamento |
| CI/CD (se aplicavel) inicial | Pipeline com build + testes + artefatos de validação | a definir | Pendente |

### 4. Backlog da Iteração

será criado em projects no github

### 5. Planejamento de Testes

| Tipo | Escopo | Evidência esperada | Responsável | Status |
|---|---|---|---|---|
| Unitário CLI | Parsing de comandos e validação inicial | Relatório de testes + cobertura mínima | a definir | Pendente |
| Unitário assinador | Validação de parâmetros e regras de erro | Casos válidos/inválidos automatizados | a definir | Pendente |
| Integração local | CLI -> `java -jar assinador.jar` | Execução ponta a ponta em ambiente local | a definir | Pendente |
| Integração HTTP | CLI -> endpoint do assinador em servidor | Execução ponta a ponta com warm start | a definir | Pendente |
| Aceitação | Critérios de US-01 e US-02 | Checklist de aceite por requisito | a definir | Pendente |

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