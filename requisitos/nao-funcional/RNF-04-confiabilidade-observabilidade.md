# RNF-04 - Confiabilidade e Observabilidade

## Objetivo

Garantir comportamento previsivel em falhas e disponibilidade de informacoes para diagnostico.

## Escopo

- Operacoes do CLI e interacoes com assinador/simulador
- Logs de execucao

## Cenario Principal de Conformidade

1. Dado comando executado com sucesso.
2. Quando operacao termina.
3. Entao sistema registra evento minimo de auditoria e retorna codigo `0`.

## Cenarios Alternativos

### CA-01 - Erro de validacao

- Sistema retorna erro estruturado, sem stack trace bruto para usuario final.

### CA-02 - Erro de dependencia externa

- Sistema registra detalhe tecnico e apresenta mensagem orientativa.

## Regras de Negocio

- RNF04-01: Erros devem seguir formato padrao (codigo, mensagem, detalhe opcional).
- RNF04-02: Logs devem conter timestamp, comando, modo e resultado.
- RNF04-03: Estado do simulador deve ser consultavel por comando dedicado (`status`).

## Criterios de Aceitacao

- 100% dos comandos retornam codigo de saida consistente.
- Logs minimos presentes em sucesso e falha.
- Erros reproduziveis possuem informacao suficiente para suporte.

## Rastreabilidade

- Origem: Especificacao, secao 6.3 e 8.2.
