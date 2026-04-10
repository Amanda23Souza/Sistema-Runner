# RNF-01 - Desempenho

## Objetivo

Definir metas minimas de tempo de resposta para operacoes do Runner em modos local e HTTP.

## Escopo

- Comandos CLI de assinatura e validacao
- Operacoes de start/status/stop do simulador

## Cenario Principal de Conformidade

1. Dado ambiente estavel e artefatos disponiveis.
2. Quando usuario executa operacao de assinatura/validacao via HTTP (warm start).
3. Entao o tempo medio de resposta deve atender ao SLO definido.

## Cenarios Alternativos

### CA-01 - Primeira execucao (cold start)

- Tempo pode ser maior, mas deve permanecer dentro do limite maximo estabelecido.

### CA-02 - Degradacao por indisponibilidade externa

- Sistema deve falhar rapido com mensagem clara, sem travamento.

## Regras de Negocio

- RNF01-01: Operacoes HTTP em warm start devem ser mais rapidas que cold start equivalente.
- RNF01-02: Comando `status` do simulador deve responder rapidamente em estado normal.
- RNF01-03: Limites de desempenho devem ser mensuraveis em testes automatizados.

## Criterios de Aceitacao (SLO inicial)

- p95 de assinatura/validacao em modo HTTP <= 1500 ms.
- p95 de `status` do simulador <= 500 ms.
- Falhas externas retornam em ate 5 s com erro controlado.

## Rastreabilidade

- Origem: Lacuna de qualidade identificada no planejamento (ISO 25010).
