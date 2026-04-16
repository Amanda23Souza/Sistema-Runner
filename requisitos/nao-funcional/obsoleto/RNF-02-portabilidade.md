# RNF-02 - Portabilidade

## Objetivo

Garantir execucao equivalente do Runner em Windows, Linux e macOS (amd64).

## Escopo

- CLI `assinatura`
- CLI `simulador`
- Provisionamento de JDK

## Cenario Principal de Conformidade

1. Dado build multiplataforma publicado.
2. Quando usuario executa os comandos principais em cada SO suportado.
3. Entao comportamento funcional deve ser equivalente.

## Cenarios Alternativos

### CA-01 - Diferenca de permissao/path

- Sistema adapta paths e permissoes por SO sem alterar contrato funcional.

### CA-02 - Ferramenta nativa indisponivel

- Sistema informa dependencia ausente com orientacao de resolucao.

## Regras de Negocio

- RNF02-01: Toda release deve conter artefatos para Windows, Linux e macOS.
- RNF02-02: Scripts/execucao nao devem depender de caminho absoluto fixo.
- RNF02-03: Diferencas de shell nao devem alterar a semantica dos comandos.

## Criterios de Aceitacao

- Testes de smoke executam com sucesso nas tres plataformas.
- Provisionamento de JDK funciona nas tres plataformas.
- Mensagens de erro permanecem coerentes em todos os SOs.

## Rastreabilidade

- Origem: US-04 e US-05 + atributo de portabilidade (ISO 25010).
