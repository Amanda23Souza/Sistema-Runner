# Regras de Negocio Gerais (Padrao)

Estas regras se aplicam a todos os requisitos funcionais, salvo quando houver regra especifica mais restritiva.

## RNP-01 - Validacao de entrada

Todo comando CLI deve validar obrigatoriedade, tipo, formato e consistencia dos parametros antes de executar acao de negocio.

## RNP-02 - Mensageria de erro

Erros devem ser retornados em formato padrao contendo:

- codigo de erro
- mensagem clara para usuario
- detalhe tecnico opcional para diagnostico

## RNP-03 - Codigos de saida

Comandos CLI devem retornar:

- `0` para sucesso
- diferente de `0` para erro

## RNP-04 - Rastreabilidade minima

Operacoes devem registrar, no minimo:

- timestamp
- comando executado
- modo de execucao
- resultado (sucesso/erro)

## RNP-05 - Idempotencia de operacoes administrativas

Comandos administrativos (iniciar/parar/status/download) devem lidar com repeticao sem causar estado inconsistente.

## RNP-06 - Compatibilidade de plataforma

Comportamentos funcionais devem ser equivalentes em Windows, Linux e macOS, respeitando diferencas de ambiente.
