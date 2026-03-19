# Sistema-Runner
Repositório dedicado para desenvolvimento do projeto Runner, da disciplina implementação e integração de software.

## Planejamento da Iteração

### 5. Planejamento de Testes

O projeto usará BDD com Cucumber para testes de aceitação e JUnit/Mockito para testes unitários.

### 6. Planejamento de Documentação

| Entregável | Conteúdo mínimo |  Responsável | Status |
|---|---|---|---|
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