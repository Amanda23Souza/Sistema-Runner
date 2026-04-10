# **Versão de documento: 1.2**

# **Autor:**  Marcello Ronald \- 02/04/2026 **Revisado** por:

---

**Como** usuário do Sistema Runner    
**Quero** executar comandos de assinatura digital através da linha de comandos    
**Para que** eu possa invocar a aplicação **assinador.jar** (doravante, Assinador) sem conhecer os detalhes técnicos de configuração Java

---

# **US-01 \- Invocar Assinador via CLI**

## **Objetivo**

O objetivo principal é **permitir que o usuário execute comandos de criação e validação de assinatura digital diretamente pela linha de comando (CLI)**.

Dessa forma, o usuário pode interagir com a funcionalidade de assinatura de forma simplificada, sem a necessidade de conhecimento aprofundado em detalhes de configuração ou execução de ambientes Java.

## **Atores Envolvidos** 

| Ator | Descrição |
| :---- | :---- |
| **Usuário** | Entidade que executa o comando na linha de comando. |
| **CLI Assinatura** | A Interface de Linha de Comando, responsável por receber, validar e rotear a requisição. Implementada referencialmente em Go. |
| **`assinador.jar`** | O componente de backend Java que, de fato, realiza as operações criptográficas de criação e validação de assinatura. |

## **Pré-condições**

Para que a US seja executada com sucesso, as seguintes condições devem ser atendidas:

* **CLI de Assinatura Disponível:** O executável do CLI de assinatura deve estar instalado e acessível no `PATH` do sistema operacional do usuário.  
* **`assinador.jar` Acessível:** O binário `assinador.jar` deve estar disponível para o CLI, seja:  
  * **Localmente:** No sistema de arquivos do usuário.  
  * **Em Modo Servidor HTTP:** Rodando como um serviço acessível por rede (localhost ou remoto).

## **Pós-condições**

Após a conclusão do fluxo, o sistema deve garantir:

* **Resposta Legível:** A operação deve ser concluída com uma resposta formatada e clara para o usuário   
  * Em caso de sucesso, exibir **MS-01**  
  * Em caso de erro, exibir **MS-02**

## **Fluxo Principal**

1\. O usuário executa comando de assinatura (criar ou validar) no CLI.  
2\. A CLI valida sintaxe e parâmetros mínimos.  
3\. A CLI seleciona modo de invocação (local ou HTTP).  
4\. A CLI envia requisição ao \`assinador.jar\`.  
5\. O \`assinador.jar\` processa e retorna resultado simulado.  
6\. A CLI formata e exibe o resultado.  
7\. A CLI finaliza com código \`0\` e exibe mensagem

## 

## **Fluxos Alternativos**

### FA-01 \- Parâmetro Osbrigatório Ausente

| Passo | Ação |
| :---- | :---- |
| **1** | No Passo 2 do Fluxo Principal, o CLI identifica a omissão de um parâmetro crucial. |
| **2** | O CLI exibe uma mensagem de erro padronizada e o guia de uso do comando. |
| **3** | O CLI finaliza com um código de saída diferente de `0` (ex: `1`). |

### FA-02 \- Falha de Comunicação com assinador.jar

| Passo | Ação |
| :---- | :---- |
| **1** | No Passo 4, o CLI tenta a comunicação, mas falha. |
| **2** | O CLI exibe um erro específico de conectividade e fornece sugestões de correção (ex: "Verifique se o JDK está instalado", "Verifique o endereço/porta do servidor"). |
| **3** | O CLI finaliza com um código de saída diferente de `0` (ex: `2`). |

### FA-03 \- Resposta Inválida do `assinador.jar`

| Passo | Ação |
| :---- | :---- |
| **1** | No Passo 5, a resposta recebida do `assinador.jar` está em um formato inesperado (JSON malformado, cabeçalhos incorretos, etc.). |
| **2** | O CLI sinaliza um erro de protocolo ou de contrato de comunicação (impedindo que a resposta inválida seja repassada ao usuário). |
| **3** | O CLI finaliza com um código de saída diferente de `0` (ex: `3`). |

## **Regras de Negócio (RN)**

| ID | Regra |
| :---- | :---- |
| **RN-US01-01** | O CLI deve suportar explicitamente os modos de invocação local (via java \-jar) e http (requisição de rede). |
| **RN-US01-02** | O resultado final (sucesso ou erro) deve ser apresentado em um formato legível para consumo humano (ex: JSON indentado, ou texto informativo). |
| **RN-US01-03** | O CLI deve aplicar as regras de negócio de política de assinatura: RNP-01, RNP-02 e RNP-03 (detalhadas em documento de Regras de Negócio de Políticas). |

## **Critérios de Aceitação**

| Critério | Descrição |
| ----- | ----- |
| **CLI e Versão** | O projeto Go está inicializado, o CLI é compilável e o comando `assinatura version` funciona. |
| **Parsing de Comandos** | O *parsing* dos comandos principais (`sign`, `validate`) e suas *flags* (`--input`, `--output`, `--mode`, etc.) é suportado, incluindo a exibição de ajuda com `--help`. |
| **Invocação Local** | A **Invocação Local** via `java -jar assinador.jar ...` funciona corretamente e trata erros comuns de ambiente (JDK ausente, arquivo jar não encontrado). |
| **Formato da Saída** | O resultado das operações de assinatura e validação são apresentados de forma **legível e informativa**. |
| **Inicialização do Servidor** | O *start* do servidor HTTP do `assinador.jar` é suportado, seja na porta padrão ou em uma porta customizada (ex: `assinatura start --port 8080`). |
| **Invocação HTTP e Fallback** | A **Invocação via HTTP** para *endpoints* `/sign` e `/validate` é funcional e inclui um mecanismo de *fallback* para o modo local, caso o modo HTTP não esteja ativo. |
| **Detecção de Instância** | Existe um mecanismo de **detecção de instância ativa** do `assinador.jar` para evitar a criação de múltiplos processos desnecessários. |
| **Comando de Parada** | O comando de parada (`assinatura stop`) encerra o processo objetivo do `assinador.jar` e atualiza o registro de estado. |
| **Timeout de Inatividade** | É possível configurar um **Timeout de Inatividade** (ex: `assinatura start --timeout 5`) que encerra o servidor automaticamente após o período especificado (em minutos). |
| **Mensagens de Sucesso/Erro** | As mensagens de sucesso e de erro são claras, concisas e fornecem informações úteis para correção de problemas pelo usuário. |

## **Mensagens padronizadas**

| Código | Descrição  |
| :---- | :---- |
| `MS-01` | Operação concluída com sucesso. |
| `MS-02` | Operação concluída com erro. |
| MS-03 | Falha: Parâmetro obrigatório ausente ou sintaxe de comando inválida. |
| MS-04 | Falha: Erro de comunicação ou conectividade com o assinador.jar (Local ou HTTP). Sugere verificar JDK/servidor. (Baseado em FA-02) |
| MS-05 | Falha: Resposta do assinador.jar em formato inválido ou inesperado (Protocolo de comunicação quebrado). (Baseado em FA-03) |
| MS-06 | Falha: JDK (Java Development Kit) não encontrado ou inacessível para execução do modo local. (Baseado no Critério de Invocação Local) |
| MS-07 | Falha: O binário assinador.jar não foi encontrado no local especificado ou acessível. (Baseado no Critério de Invocação Local) |
| MS-08 | Falha: Ocorreu um erro interno no processo de assinatura ou validação do assinador.jar. |
| MS-09 | Falha: Timeout de inatividade atingido. O servidor HTTP foi encerrado automaticamente. (Baseado no Critério de Timeout de Inatividade) |

