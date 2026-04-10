# US-01 - Invocar Assinador via CLI

## Objetivo

Permitir que o usuario execute comandos de criacao e validacao de assinatura por linha de comando sem conhecer detalhes de configuracao Java.

## Atores

- Usuario
- CLI `assinatura`
- `assinador.jar`

## Pre-condicoes

- CLI `assinatura` disponivel no sistema.
- `assinador.jar` acessivel localmente ou em modo servidor HTTP.

## Pos-condicoes

- Operacao concluida com resposta legivel para o usuario.
- Codigo de saida retornado conforme resultado.

## Fluxo Principal (Caminho Feliz)

1. Usuario executa comando de assinatura (criar ou validar) no CLI.
2. CLI valida sintaxe e parametros minimos.
3. CLI seleciona modo de invocacao (local ou HTTP).
4. CLI envia requisicao ao `assinador.jar`.
5. `assinador.jar` processa e retorna resultado simulado.
6. CLI formata e exibe o resultado.
7. CLI finaliza com codigo `0`.

## Fluxos Alternativos

### FA-01 - Parametro obrigatorio ausente

1. No passo 2, CLI identifica parametro ausente.
2. CLI exibe mensagem de uso e erro padronizado.
3. CLI finaliza com codigo diferente de `0`.

### FA-02 - Falha de comunicacao com assinador.jar

1. No passo 4, comunicacao falha (processo nao inicia, timeout ou HTTP indisponivel).
2. CLI exibe erro de conectividade e sugestao de correcao.
3. CLI finaliza com codigo diferente de `0`.

### FA-03 - Resposta invalida do assinador.jar

1. No passo 5, resposta retorna formato inesperado.
2. CLI sinaliza erro de protocolo/contrato.
3. CLI finaliza com codigo diferente de `0`.

## Regras de Negocio

- RN-US01-01: O CLI deve suportar explicitamente os modos `local` e `http`.
- RN-US01-02: O resultado deve ser apresentado em formato legivel para humano.
- RN-US01-03: O CLI deve aplicar `RNP-01`, `RNP-02` e `RNP-03`.

## Criterios de Aceitacao

- Comandos de criar e validar sao aceitos.
- Invocacao local e HTTP funcionam com mesmo resultado logico.
- Mensagens de sucesso e erro sao claras.

## Rastreabilidade

- Origem: Especificacao, secao 5, US-01.
- Relacao: Integracao entre aplicacoes (secao 6).
