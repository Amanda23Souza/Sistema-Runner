# US-03 - Gerenciar Ciclo de Vida do Simulador do HubSaude

## Objetivo

Permitir iniciar, parar e consultar status do `simulador.jar` via CLI, incluindo obtencao automatica da versao mais recente quando necessario.

## Atores

- Usuario
- CLI `simulador`
- `simulador.jar` (externo ao desenvolvimento)
- GitHub Releases da disciplina

## Pre-condicoes

- CLI `simulador` instalado.
- Conectividade com internet para verificacao/download (quando necessario).

## Pos-condicoes

- Simulador iniciado, parado ou consultado com retorno claro.
- Artefato local atualizado somente quando necessario.

## Fluxo Principal (Caminho Feliz)

1. Usuario executa comando `simulador start`.
2. CLI verifica portas necessarias.
3. CLI verifica se a versao mais recente do `simulador.jar` ja existe localmente.
4. Se nao existir, CLI baixa a versao mais recente do release.
5. CLI inicia o simulador e registra identificador do processo.
6. Usuario pode executar `simulador status` para verificar estado.
7. Usuario executa `simulador stop` para encerrar com seguranca.

## Fluxos Alternativos

### FA-01 - Porta indisponivel

1. No passo 2, porta requerida esta em uso.
2. CLI aborta inicio com mensagem orientativa.
3. CLI sugere porta alternativa ou liberacao da porta.

### FA-02 - Falha no download do simulador

1. No passo 4, ocorre falha de rede ou release indisponivel.
2. CLI retorna erro de download e nao inicia processo.

### FA-03 - Simulador ja em execucao

1. No passo 1, CLI detecta processo ativo.
2. CLI informa estado atual e evita inicializacao duplicada.

### FA-04 - Stop sem processo ativo

1. Usuario executa `simulador stop` sem processo em execucao.
2. CLI responde de forma idempotente, sem erro critico.

## Regras de Negocio

- RN-US03-01: O `simulador.jar` nao e desenvolvido no escopo do projeto, apenas orquestrado.
- RN-US03-02: Download deve ocorrer somente se a versao mais recente nao estiver disponivel localmente.
- RN-US03-03: Verificacao de portas e obrigatoria antes do start.
- RN-US03-04: O sistema deve aplicar `RNP-02`, `RNP-05` e `RNP-06`.

## Criterios de Aceitacao

- Start, stop e status funcionam por CLI.
- Portas sao verificadas antes de iniciar.
- Download e condicional a ausencia de versao atual.
- Estado do simulador e exibido corretamente.

## Rastreabilidade

- Origem: Especificacao, secao 5, US-03.
- Relacao: Objetivo de gerenciamento de ciclo de vida do simulador.
