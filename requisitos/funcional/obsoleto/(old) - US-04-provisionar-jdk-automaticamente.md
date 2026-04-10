# US-04 - Provisionar JDK Automaticamente

## Objetivo

Detectar ausencia/incompatibilidade de JDK e realizar provisionamento automatico para uso pelo `assinador.jar` e `simulador.jar`.

## Atores

- Usuario
- CLI do Runner
- Repositorio de distribuicao de JDK

## Pre-condicoes

- Runner executando em Windows, Linux ou macOS.

## Pos-condicoes

- JDK requerido disponivel localmente para o Runner.

## Fluxo Principal (Caminho Feliz)

1. CLI identifica versao de JDK requerida pelo sistema.
2. CLI verifica disponibilidade e versao do JDK local.
3. Se ausente/incompativel, CLI baixa distribuicao compativel com SO/arquitetura.
4. CLI instala/configura JDK no diretorio do Runner.
5. CLI registra caminho do JDK provisionado para uso interno.
6. CLI prossegue com execucao normal de assinador/simulador.

## Fluxos Alternativos

### FA-01 - JDK compativel ja presente

1. No passo 2, JDK compativel e encontrado.
2. CLI reutiliza instalacao existente.
3. Nao realiza download.

### FA-02 - Falha de download

1. No passo 3, download falha.
2. CLI retorna erro com diagnostico e orientacao de retentativa.

### FA-03 - Pacote invalido/corrompido

1. No passo 4, pacote nao pode ser instalado.
2. CLI interrompe provisionamento e registra erro.

## Regras de Negocio

- RN-US04-01: A deteccao de versao de JDK e obrigatoria antes do download.
- RN-US04-02: O JDK provisionado e de uso interno do Runner, sem exigir configuracao manual global.
- RN-US04-03: Provisionamento deve suportar as tres plataformas alvo (Windows, Linux, macOS).
- RN-US04-04: O sistema deve aplicar `RNP-04` e `RNP-06`.

## Criterios de Aceitacao

- Sistema detecta presenca e versao do JDK.
- Sistema baixa e configura JDK quando necessario.
- Assinador e simulador usam o JDK provisionado.
- Comportamento e consistente nas tres plataformas.

## Rastreabilidade

- Origem: Especificacao, secao 5, US-04.
