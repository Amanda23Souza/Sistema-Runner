# RNF-05 - Manutenibilidade e Testabilidade

## Objetivo

Promover evolucao segura do sistema por meio de modularidade, cobertura de testes e padroes de codigo.

## Escopo

- Codigo do CLI
- Codigo do `assinador.jar`
- Pipeline de testes

## Cenario Principal de Conformidade

1. Dado nova funcionalidade implementada.
2. Quando alteracao e submetida para revisao.
3. Entao deve haver testes automatizados e evidencias de qualidade.

## Cenarios Alternativos

### CA-01 - Alteracao sem testes

- Merge deve ser bloqueado ate inclusao de testes minimos.

### CA-02 - Regressao detectada

- Alteracao retorna para correcao antes de release.

## Regras de Negocio

- RNF05-01: Cada requisito funcional deve possuir ao menos um teste associado.
- RNF05-02: Codigo novo deve seguir convencoes definidas pela equipe.
- RNF05-03: Revisao por par (dupla) e obrigatoria para merge em branch principal.

## Criterios de Aceitacao

- Suite de testes unitarios e de integracao executa no CI.
- Requisitos US-01 a US-05 possuem rastreio para casos de teste.
- Evidencias de revisao e refatoracao registradas no planejamento.

## Rastreabilidade

- Origem: Especificacao, secao 7 (testes), secao 8.2 (qualidade) e planejamento.
