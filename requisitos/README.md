# Catálogo de Requisitos

Este diretório organiza os requisitos do Sistema Runner em duas categorias:

- `funcional/`: requisitos funcionais derivados das histórias de usuário US-01 a US-05.
- `nao-funcional/`: requisitos de qualidade propostos para complementar a especificação.

## Estrutura

| Categoria | ID | Arquivo | Descricao |
|---|---|---|---|
| Funcional | RNP | funcional/regras-negocio-gerais.md | Regras padrao aplicadas a todos os requisitos funcionais |
| Funcional | US-01 | funcional/US-01-invocar-assinador-via-cli.md | Invocacao do assinador por linha de comando |
| Funcional | US-02 | funcional/US-02-simular-assinatura-validacao-parametros.md | Validacao de parametros e simulacao de assinatura |
| Funcional | US-03 | funcional/US-03-gerenciar-simulador-hubsaude.md | Gestao de ciclo de vida do simulador |
| Funcional | US-04 | funcional/US-04-provisionar-jdk-automaticamente.md | Provisionamento automatico de JDK |
| Funcional | US-05 | funcional/US-05-disponibilizar-binarios-multiplataforma.md | Publicacao de binarios multiplataforma |
| Nao funcional | RNF-01 | nao-funcional/RNF-01-desempenho.md | Metas de desempenho e tempo de resposta |
| Nao funcional | RNF-02 | nao-funcional/RNF-02-portabilidade.md | Compatibilidade entre Windows, Linux e macOS |
| Nao funcional | RNF-03 | nao-funcional/RNF-03-seguranca.md | Integridade de artefatos e assinatura Cosign |
| Nao funcional | RNF-04 | nao-funcional/RNF-04-confiabilidade-observabilidade.md | Tratamento de falhas, logs e observabilidade |
| Nao funcional | RNF-05 | nao-funcional/RNF-05-manutenibilidade-testabilidade.md | Evolucao segura com testes e padroes |

## Padrão adotado

Cada requisito, sempre que aplicável, inclui:

- objetivo
- atores
- pre-condicoes e pos-condicoes
- fluxo principal (caminho feliz)
- fluxos alternativos
- regras de negocio padronizadas
- criterios de aceitacao
- rastreabilidade
