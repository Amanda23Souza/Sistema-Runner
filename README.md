# Sistema-Runner

Repositório dedicado para desenvolvimento do **projeto Runner**, da disciplina **Implementação e Integração de Software** (UFG).

O Sistema Runner é uma aplicação CLI (Command-Line Interface) em Go que facilita operações de assinatura digital através da linha de comando, sem necessidade de conhecimento aprofundado em configurações Java.

### Arquivos designados pelo professor:

- [Especificação](docs/aulas/especificacao.md)
- [Design](docs/design.md)
- [Plano de implementação](docs/aulas/plano-revisitado-v2.md)
- [Sprint 1](docs/aulas/sprint-1-tasks.md)

### [📌 Nosso planejamento](https://github.com/Amanda23Souza/Sistema-Runner/blob/main/docs/planejamento/nossoPlanejamento.md)


---

## 📋 Requisitos

- **Go 1.26.1+** (ou superior)
- **JDK** instalado (para operações de assinatura digital com `assinador.jar`)

---

## 📁 Estrutura do Projeto

O repositório é composto por dois subprojetos principais:

1. **`cli-assinatura/`**: A interface de linha de comandos (CLI) escrita em Go (1.25).
2. **`docs/aulas/projetos/assinador-java/`**: O backend do assinador executável escrito em Java 21, responsável pelas validações de parâmetros e simulação de assinaturas, agora integrável via HTTP e local.

```
.
├── cli-assinatura/              ← Subprojeto CLI (Go)
│   ├── cmd/assinatura/
│   │   └── main.go              ← Ponto de entrada (executável)
│   ├── internal/
│   │   ├── command/
│   │   │   ├── root.go           ← Parser e orquestrador
│   │   │   ├── version.go        ← Comando de versão
│   │   │   ├── sign.go           ← Comando de assinatura (stub)
│   │   │   └── validate.go       ← Comando de validação (stub)
│   │   └── version/
│   └── go.mod                   ← Definição de módulo Go
└── docs/aulas/projetos/
    └── assinador-java/          ← Subprojeto Assinador (Java)
        ├── src/
        │   ├── main/java/com/runner/assinador/
        │   │   ├── App.java                   ← Classe Main (servidor Javalin)
        │   │   ├── SignatureController.java    ← Rotas /sign, /validate e /health
        │   │   ├── SignatureService.java       ← Interface do serviço
        │   │   ├── FakeSignatureService.java   ← Simulação e validação
        │   │   └── domain/                    ← DTOs de Request/Response
        │   └── test/                          ← Testes unitários e de integração
        └── pom.xml                            ← Definição Maven (Javalin, Jackson)
```

---

## 🚀 Como Compilar e Executar

### 1. CLI (Go)
```bash
cd cli-assinatura
go build -o assinatura ./cmd/assinatura
```

### 2. Assinador (Java)
Para compilar e empacotar o executável do servidor Java (gerando o *fat JAR* com as dependências do Javalin embutidas):
```bash
cd docs/aulas/projetos/assinador-java
mvn clean package
```

Para rodar o servidor em uma porta customizada (ex: `8080`):
```bash
java -jar target/assinador-java-1.0.0-SNAPSHOT-jar-with-dependencies.jar 8080
```

---

## 💻 Como Usar

### CLI (Go)
O CLI atual aceita os comandos principais `sign` e `validate`, e também os aliases `criar` e `validar`.

```bash
# Modo de ajuda geral
./assinatura --help

# Exibir versão do CLI
./assinatura version
./assinatura version --json
```

### Assinador HTTP Server (Java)
O servidor HTTP aceita requisições JSON em três endpoints configurados:

*   **`GET /health`**: Endpoint de monitoramento de saúde do processo.
    ```bash
    curl -i http://localhost:8080/health
    # Retorna: {"status":"UP"} com HTTP 200
    ```
*   **`POST /sign`**: Criação de assinatura simulada.
    ```bash
    curl -X POST http://localhost:8080/sign \
      -H "Content-Type: application/json" \
      -d '{"content": "meu_conteudo"}'
    ```
*   **`POST /validate`**: Validação de assinatura.
    ```bash
    curl -X POST http://localhost:8080/validate \
      -H "Content-Type: application/json" \
      -d '{"content": "meu_conteudo", "signature": "MOCKED_SIGNATURE_BASE64_=="}'
    ```

---

## 📊 Status de Implementação

### ✅ Fase 1 - Estrutura Base (Completo)
- ✓ Estrutura de pacotes conforme padrão Go (cmd/, internal/)
- ✓ Interface `Command` padronizada para todos os comandos
- ✓ Ponto de entrada minimalista em `cmd/assinatura/main.go`
- ✓ Gerenciamento centralizado de versão

### ✅ Fase 2 - Comando Version (Completo)
- ✓ Comando `version` totalmente funcional
- ✓ Flags: `--quiet`, `--json`, `--help`
- ✓ Exit codes corretos (0 para sucesso, 1 para erro)

### ✅ Fase 3 - Backend Java Base & Validações (Completo)
- ✓ Criação do subprojeto `assinador-java`
- ✓ Definição do `SignatureService` e `FakeSignatureService`
- ✓ Validação de parâmetros de assinatura no backend

### ⏳ Fase 4 - Modo Servidor HTTP & Integração (Em Progresso)
- ✓ Servidor web leve baseado em **Javalin** e **Jackson** no backend Java (Completo)
- ✓ Endpoints de API `/sign` e `/validate` robustos com tratamento de erro (Completo)
- ✓ Endpoint de monitoramento de saúde `/health` funcional (Completo)
- ✓ Suíte de testes de integração `SignatureControllerTest` validando todas as rotas (Completo)
- ⏳ Integração real CLI-HTTP (Comunicação da CLI com a API Java)
- ⏳ Comandos `start`/`stop` e monitoramento de ciclo de vida na CLI Go

---

## 📖 Documentação

- **[User Story 01 - Invocar Assinador via CLI](./docs/US-01%20-%20Invocar%20Assinador%20via%20CLI.md)** — Requisitos funcionais e critérios de aceitação
- **[Planejamento do Projeto](./docs/planejamento/nossoPlanejamento.md)** — Visão geral e timeline
- **[Design de Arquitetura](./docs/design.md)** — Padrões e decisões técnicas

---

## 🔧 Próximos Passos

1. **Implementar comando `sign` / `validate` funcionais na CLI Go**:
   - Chamar `java -jar assinador.jar` no modo local (*cold start*).
   - Enviar requisições HTTP para a API Java no modo HTTP (*warm start*).

2. **Implementar comandos `start`/`stop` na CLI Go**:
   - Iniciar servidor HTTP do `assinador.jar` em background.
   - Encerrar o processo e monitorar status ativo via `/health`.

3. **Melhorias Técnicas**:
   - Provisionamento automático do JDK.
   - Setup de `Makefile` para build multi-plataforma e pipelines avançadas.
