## Planejamento da Iteração

## Integrantes

| Nome | GitHub |
---|---|
|  Marcello Ronald | https://github.com/mronald-js |
| Amanda Soares  | https://github.com/Amanda23Souza |  |

## Requisitos

Funcionais

| Requisito | Descrição |
|---|---|
| [RF-01](https://docs.google.com/document/d/1INuAPolIXwJDsnybkzo76w2qEoZFv5BsP6o_vdRQ974/edit?usp=drive_link) | permitir que o usuário execute comandos de criação e validação de assinatura digital diretamente pela linha de comando (CLI) |

Não funcionais

| Requisito | Descrição |
|---|---|
| [RNF-01](../requisitos/nao-funcional/RNF-01-desempenho.md) | |

### 1. Interações

Cobertura de aulas

| Sprints | Objetivo |
|---|---|
| Sprint 1 - (08/04 - 22/04) | Estrutura base do CLI + Pipeline CI/CD + Primeiro build/release + Simulador CLI |
| Sprint 2 - (29/04 - 13/05) | Assinador CLI + Assinador.jar |
| Sprint 3 - (20/05 - 03/06) | modo servidor HTTP, gestão de ciclo de vida e testes de integração principais |
| Sprint 4 - (10/06 - 24/06) | Revisar, Testes de aceitação finais, documentação e release final |

### 2. Design

| Artefato | Local |
|---|---|
| Diagrama de contexto | [contexto](../diagramas/imagens/contexto.svg) |
| Diagrama de conteiner | [conteiner](../diagramas/imagens/conteineres.svg) | 

Referência adicional: design amostral em [especificação/design.md](../docs/design.md).

### 3. Preparação de Ambiente e Convenções

| Tópico | Definição |
|---|---|
| Linguagem e versão | Go e Java 21 |  | Em andamento |
| Ferramentas de build | A definir |
| CI/CD | Pipeline com build + testes + artefatos de validação | Pendente |


#### Convenções de código

- **Commits semânticos:** use o padrão `type: subject` (ex.: `feat(cli): adicionar comando criar`).
- **Revisão por par:** PRs se abertos devem receber aprovação do outro integrante antes do merge.

#### Estratégia de branches (Git Flow)

- **Branches principais:** `main`
- **Fluxo:** Quando uma funcionalidade nova for integrada solicitar pr para `main` a partir de uma branch específica por ex:. `feat/{nome-da-feature}`, se for correção de bug fazer o mesmo processo mas como `bugfix/{bug-relacionado}`
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
