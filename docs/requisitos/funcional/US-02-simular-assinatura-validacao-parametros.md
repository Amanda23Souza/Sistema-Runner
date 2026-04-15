# US-02 - Simular Assinatura Digital com Validacao de Parametros

## Objetivo

Garantir validacao rigorosa de parametros antes de simular criacao ou validacao de assinatura digital.

## Atores

- Usuario
- CLI `assinatura`
- `assinador.jar`
- Dispositivo criptografico (token/smart card via PKCS#11)

## Pre-condicoes

- Requisicao recebida pelo `assinador.jar`.
- Parametros de entrada fornecidos pelo usuario.

## Pos-condicoes

- Em caso valido: resposta simulada retornada.
- Em caso invalido: erro padronizado retornado.

## Fluxo Principal (Caminho Feliz)

1. `assinador.jar` recebe operacao (`criar` ou `validar`) e parametros.
2. Sistema valida obrigatoriedade, formato e consistencia dos parametros.
3. Para `criar`, sistema produz assinatura simulada pre-construida.
4. Para `validar`, sistema produz resultado pre-determinado conforme regra definida.
5. Sistema retorna resposta estruturada ao CLI.

## Fluxos Alternativos

### FA-01 - Parametro invalido

1. No passo 2, validacao falha.
2. Sistema interrompe processamento.
3. Sistema retorna erro detalhando parametro e motivo.

### FA-02 - Falha na interacao PKCS#11

1. Durante processamento, biblioteca/dispositivo PKCS#11 nao responde.
2. Sistema retorna erro tecnico controlado.
3. Nao deve haver travamento do processo.

### FA-03 - Operacao desconhecida

1. No passo 1, operacao diferente de `criar`/`validar`.
2. Sistema retorna erro de operacao nao suportada.

## Regras de Negocio

- RN-US02-01: Nenhuma simulacao deve ocorrer sem validacao previa bem-sucedida.
- RN-US02-02: Erros devem indicar claramente qual campo falhou e regra violada.
- RN-US02-03: Integracao PKCS#11 deve ser opcional no ambiente de simulacao, mas com tratamento de falhas.
- RN-US02-04: O sistema deve aplicar `RNP-01`, `RNP-02` e `RNP-04`.

## Criterios de Aceitacao

- Todos os parametros definidos sao validados.
- Criacao de assinatura retorna payload simulado quando valido.
- Validacao retorna resultado pre-determinado quando valido.
- Erros invalidos sao claros e verificaveis.

## Rastreabilidade

- Origem: Especificacao, secao 5, US-02.
- Relacao: Qualidade de validacao e tratamento de erros (secao 6.3 e 8.2).
