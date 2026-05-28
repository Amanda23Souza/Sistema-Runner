# ADR-002 — Modo Servidor HTTP como Padrão

**Data:** 2026-04-22
**Status:** Aceito
**Autores:** Amanda Soares, Marcello Ronald
**Referência à especificação:** [`kyriosdata/runner @ d3f1a9c`](https://github.com/kyriosdata/runner/blob/d3f1a9c/docs/runner.md)

---

## Contexto

A especificação prevê dois modos de operação para o assinador:
1. **Modo local**: o CLI invoca `java -jar assinador.jar` diretamente a cada chamada (cold start ~500ms–2s).
2. **Modo HTTP**: o CLI comunica-se com uma instância do `assinador.jar` já em execução via HTTP (warm start ~10ms).

Precisamos definir qual modo é o **padrão** quando o usuário não especifica `--mode`.

## Decisão

O **modo HTTP é o padrão** (`--mode http`). O modo local deve ser ativado explicitamente com `--mode local`.

## Racional

- A especificação indica que o modo servidor é o cenário de uso principal (item E2.5 dos critérios).
- O warm start elimina a latência de JVM em operações repetidas (ex: assinar um lote de documentos).
- Usuários avançados que precisam do modo local podem fazê-lo explicitamente.
- Mantém consistência com ferramentas similares (ex: Language Server Protocol — servidor em background).

## Consequências

- **Positivas:** experiência de usuário mais rápida no caso de uso dominante; incentiva o uso do servidor.
- **Negativas:** o usuário precisa executar `assinatura start` antes de `assinatura sign`, o que adiciona um passo inicial; o modo local continua disponível para casos de uso simples.

## Porta padrão

A porta padrão é **8080**, configurável via `--port`. Esta decisão foi tomada por ser a porta HTTP não-privilegiada mais reconhecida, minimizando a curva de configuração.
