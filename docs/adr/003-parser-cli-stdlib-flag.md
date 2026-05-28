# ADR-003 — Parser de CLI: stdlib `flag` vs frameworks externos

**Data:** 2026-04-08
**Status:** Aceito
**Autores:** Amanda Soares, Marcello Ronald

---

## Contexto

A CLI precisa de um parser de argumentos. O ecossistema Go oferece:
- `flag` (stdlib): zero dependências, suporte a flags GNU-style limitado.
- `cobra` + `pflag`: padrão de facto para CLIs Go complexas, com subcomandos, completions e help automático.
- `urfave/cli`: alternativa popular com sintaxe mais simples.

## Decisão

Adotar a **stdlib `flag`** para o parser de argumentos.

## Racional

- A especificação requer um conjunto fixo e pequeno de subcomandos (`sign`, `validate`, `start`, `stop`, `status`, `version`) — não justifica a complexidade de cobra.
- Zero dependências externas para a CLI mantém o binário menor e a supply chain mais simples (critério F3).
- A implementação manual do dispatcher de subcomandos em `root.go` é trivial e dá controle total sobre mensagens de erro e help.

## Consequências

- **Positivas:** sem dependências externas, binário menor, sem overhead de inicialização de frameworks.
- **Negativas:** sem geração automática de completions de shell; help formatting menos sofisticado; se o número de comandos crescer significativamente, migração para cobra pode ser necessária.

## Revisão futura

Se o projeto crescer para mais de 15 subcomandos ou precisar de shell completions automáticas, reavaliar migração para `cobra`.
