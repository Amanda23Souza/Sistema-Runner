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

## [Backlog da Iteração](https://github.com/users/Amanda23Souza/projects/1/views/1?verticalGroupedBy[columnId]=271373991)

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



### 4. Planejamento de Testes

O projeto usará BDD com Cucumber para testes de aceitação e JUnit/Mockito para testes unitários.


### 5. Planejamento de Documentação

| Entregável | Conteúdo mínimo | Referência | Responsável | Status |
|---|---|---|---|---|
| Manual de uso CLI | Instalação, comandos, exemplos e erros comuns | [US-01](requisitos/funcional/US-01-invocar-assinador-via-cli.md) | a definir | Pendente |
| Doc técnica integração | Fluxos local/HTTP, contrato de entrada/saída | [US-01](requisitos/funcional/US-01-invocar-assinador-via-cli.md) | a definir | Pendente |
| Guia de testes | Como executar suíte unitária e integração | [Catálogo de Requisitos](requisitos/README.md) | a definir | Pendente |

### 6. Planejamento de Release e Distribuição

| Item | Definição | Referência | Status |
|---|---|---|---|
| SemVer | Tag no formato `vMAJOR.MINOR.PATCH` | [US-05](requisitos/funcional/US-05-disponibilizar-binarios-multiplataforma.md) | Pendente |
| Artefatos multiplataforma | Build para Windows, Linux e macOS (amd64) | [US-05](requisitos/funcional/US-05-disponibilizar-binarios-multiplataforma.md) | Pendente |
| Integridade | Publicação de checksums SHA256 | [US-05](requisitos/funcional/US-05-disponibilizar-binarios-multiplataforma.md) | Pendente |

### 7. Planejamento de Segurança de Artefatos (Cosign)

| Item | Definição | Referência | Status |
|---|---|---|---|
| Assinatura automática | CI/CD deve assinar cada artefato com Cosign | [RNF-03](requisitos/nao-funcional/RNF-03-seguranca.md) | Pendente |
| Arquivos obrigatórios | Publicar `<artefato>`, `<artefato>.sig`, `<artefato>.pem` | [RNF-03](requisitos/nao-funcional/RNF-03-seguranca.md) | Pendente |
| Verificação | Instrução com `cosign verify-blob` na release | [RNF-03](requisitos/nao-funcional/RNF-03-seguranca.md) | Pendente |