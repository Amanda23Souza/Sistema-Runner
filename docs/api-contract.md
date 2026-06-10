# Contrato de API — CLI ↔ assinador.jar

> **Versão:** 1.0.0
> **Data:** 2026-06-09
> **Referência à especificação:** [`kyriosdata/runner @ d3f1a9c`](https://github.com/kyriosdata/runner/blob/d3f1a9c/docs/runner.md)

---

## Visão Geral

A CLI (`assinatura`) comunica-se com o `assinador.jar` via HTTP REST. O servidor JAR escuta na porta configurável (padrão `8080`) e expõe endpoints JSON.

```
┌─────────────┐       HTTP/JSON        ┌─────────────────┐
│  CLI (Go)   │ ──────────────────────► │ assinador.jar   │
│  assinatura │ ◄────────────────────── │ (Java/Javalin)  │
└─────────────┘                        └─────────────────┘
```

---

## Endpoints HTTP

### `GET /health`

Verifica se o servidor está pronto para receber requisições.

**Request:** nenhum body.

**Response (200 OK):**
```json
{
  "status": "UP",
  "port": 8080,
  "inactivityTimeoutSeconds": 300
}
```

**Uso na CLI:** `start.go` (idempotência), `stop.go` (verificação), `status` (health check).

---

### `POST /sign`

Cria uma assinatura digital para o conteúdo fornecido.

**Request:**
```json
{
  "content": "nome-do-arquivo.pdf"
}
```

| Campo     | Tipo   | Obrigatório | Descrição                         |
|-----------|--------|:-----------:|-----------------------------------|
| `content` | string | ✅          | Identificador ou conteúdo a assinar |

**Response (200 OK):**
```json
{
  "signature": "MOCKED_SIGNATURE_BASE64_==",
  "valid": true,
  "message": "Assinatura criada com sucesso"
}
```

**Response (400 Bad Request) — parâmetro ausente:**
```json
{
  "signature": null,
  "valid": false,
  "message": "Parâmetro 'content' inválido ou ausente"
}
```

---

### `POST /validate`

Valida uma assinatura digital.

**Request:**
```json
{
  "content": "nome-do-arquivo.pdf",
  "signature": "MOCKED_SIGNATURE_BASE64_=="
}
```

| Campo       | Tipo   | Obrigatório | Descrição                       |
|-------------|--------|:-----------:|---------------------------------|
| `content`   | string | ✅          | Identificador ou conteúdo original |
| `signature` | string | ✅          | Assinatura a ser validada       |

**Response (200 OK) — assinatura válida:**
```json
{
  "signature": "MOCKED_SIGNATURE_BASE64_==",
  "valid": true,
  "message": "Assinatura é válida"
}
```

**Response (200 OK) — assinatura inválida:**
```json
{
  "signature": "assinatura-incorreta",
  "valid": false,
  "message": "Assinatura é inválida"
}
```

> **Nota:** Assinatura inválida retorna HTTP 200 (a requisição foi processada com sucesso). O campo `valid` indica o resultado da validação. HTTP 400 é reservado para erros de formato (JSON malformado, campo ausente).

---

### `POST /shutdown`

Encerra o servidor de forma controlada.

**Request:** `{}` (body vazio ou JSON vazio).

**Response (200 OK):**
```json
{
  "message": "Servidor encerrando..."
}
```

**Comportamento:** O servidor envia a resposta, aguarda 200ms, e então encerra o processo Javalin.

---

## Códigos de Saída da CLI

| Código | Significado | Exemplos |
|:------:|-------------|----------|
| `0`    | Sucesso | Operação completada sem erros |
| `1`    | Erro do sistema | Falha de I/O, rede inacessível, JVM ausente, processo não encerrado |
| `2`    | Erro do usuário | Parâmetro `--input` ausente, modo inválido, arquivo não encontrado pelo usuário |

---

## Códigos de Erro da CLI

| Código   | Contexto | Descrição |
|----------|----------|-----------|
| `[MS-03]`| Parsing de flags | Parâmetro inválido ou sintaxe incorreta |

---

## Flags da CLI

### Globais

| Flag | Descrição |
|------|-----------|
| `--help` | Exibe ajuda com exemplos |
| `--version` | Exibe versão + commit + buildtime |

### `sign` / `criar`

| Flag | Tipo | Padrão | Descrição |
|------|------|--------|-----------|
| `--input` | string | *(obrigatório)* | Caminho do arquivo a assinar |
| `--output` | string | `<input>.sig` | Caminho para salvar a assinatura |
| `--mode` | string | `http` | Modo: `http` (servidor) ou `local` (SHA-256 simulado) |
| `--port` | int | `8080` | Porta do servidor (modo http) |
| `--json` | bool | `false` | Saída em JSON estruturado |
| `--verbose` | bool | `false` | Saída detalhada com logs |

### `validate` / `validar`

| Flag | Tipo | Padrão | Descrição |
|------|------|--------|-----------|
| `--input` | string | *(obrigatório)* | Caminho do arquivo original |
| `--signature` | string | *(obrigatório)* | Caminho do arquivo de assinatura |
| `--mode` | string | `http` | Modo: `http` ou `local` |
| `--port` | int | `8080` | Porta do servidor (modo http) |
| `--json` | bool | `false` | Saída em JSON estruturado |
| `--verbose` | bool | `false` | Saída detalhada |

### `start`

| Flag | Tipo | Padrão | Descrição |
|------|------|--------|-----------|
| `--port` | int | `8080` | Porta em que o servidor escutará |
| `--jar` | string | *(auto-descoberta)* | Caminho para o `assinador.jar` |
| `--verbose` | bool | `false` | Saída detalhada |

### `stop`

| Flag | Tipo | Padrão | Descrição |
|------|------|--------|-----------|
| `--port` | int | `8080` | Porta do servidor a encerrar |

### `status`

| Flag | Tipo | Padrão | Descrição |
|------|------|--------|-----------|
| `--port` | int | `8080` | Porta do servidor a verificar |
| `--json` | bool | `false` | Saída em JSON estruturado |

### `version`

| Flag | Tipo | Padrão | Descrição |
|------|------|--------|-----------|
| `--quiet` | bool | `false` | Apenas o número de versão |
| `--json` | bool | `false` | Versão em JSON |

---

## Tratamento de Erros HTTP (CLI → Servidor)

| Cenário | Mensagem na CLI | Exit Code |
|---------|----------------|:---------:|
| Conexão recusada | `"Erro do sistema: falha ao conectar ao servidor assinador em localhost:<port>"` | 1 |
| Timeout (10s) | `"timeout após 10s aguardando resposta do servidor"` | 1 |
| HTTP 4xx/5xx | `"servidor retornou status HTTP <code> (esperado 2xx)"` | 1 |
| JSON malformado | `"resposta do servidor não é JSON válido"` | 1 |
| Resposta sem campo esperado | `"resposta malformada do servidor assinador"` | 1 |

---

## Configuração do Servidor Java

| Parâmetro | Argumento CLI | Padrão | Descrição |
|-----------|--------------|--------|-----------|
| Porta | 1º argumento posicional | `8080` | Porta HTTP |
| Timeout de inatividade | `--inactivity-timeout <segundos>` | `300` (5 min) | Auto-shutdown após N segundos sem requisições |

**Exemplo:**
```bash
java -jar assinador.jar 9090 --inactivity-timeout 600
```

---

## Auto-Shutdown por Inatividade

O servidor implementa auto-shutdown:
1. A cada requisição HTTP recebida, o timestamp de última atividade é atualizado.
2. Um watchdog thread verifica periodicamente se o tempo de inatividade excedeu o timeout configurado.
3. Se excedido, o servidor encerra automaticamente via `Javalin.stop()`.

---

*Documento gerado em 2026-06-09. Referência: código-fonte dos módulos `cli-assinatura/` e `docs/aulas/projetos/assinador-java/`.*
