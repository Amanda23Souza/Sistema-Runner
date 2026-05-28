# ADR-001 — Escolha de Go para a CLI

**Data:** 2026-04-08
**Status:** Aceito
**Autores:** Amanda Soares, Marcello Ronald
**Referência à especificação:** [`kyriosdata/runner @ d3f1a9c`](https://github.com/kyriosdata/runner/blob/d3f1a9c/docs/runner.md)

---

## Contexto

O projeto Runner exige uma CLI multiplataforma (Linux, Windows, macOS) que:
- Invoque o `assinador.jar` tanto diretamente (subprocesso) quanto via HTTP.
- Produza binários autônomos distribuíveis sem dependências de runtime.
- Suporte cross-compilation a partir de um único ambiente de CI.

## Decisão

Adotar **Go 1.26+** para a implementação da CLI.

## Alternativas Consideradas

| Alternativa | Motivo de descarte |
|-------------|-------------------|
| Python 3.x | Requer runtime instalado no cliente; distribuição via PyInstaller gera binários pesados e com falsos positivos em antivírus |
| Rust | Curva de aprendizado elevada para a equipe; ecossistema de CLI menos maduro que Go |
| Java (mesmo stack que o JAR) | Inicia JVM a cada invocação CLI — latência inaceitável para uso interativo |
| Node.js | Requer runtime ou bundler; menos adequado para ferramentas de sistema |

## Consequências

- **Positivas:** binário único sem dependências, cross-compilation nativa, startup < 10ms, stdlib rica para HTTP/subprocess/sinais.
- **Negativas:** time precisa aprender Go; interoperabilidade com código Java exige protocolo HTTP ou subprocesso.
