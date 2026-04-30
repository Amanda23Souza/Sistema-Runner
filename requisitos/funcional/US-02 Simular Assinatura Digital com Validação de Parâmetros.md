---

**Versão do documento**: 1.0

**Autor**: Marcello Ronald \- 08/04/2026  
**Revisado por:** 

---

**Como:** Um usuário autorizado a assinar documentos no sistema.  
**Quero:** Testar a validação da assinatura digital de um documento, sem efetivamente assiná-lo.  
**Para que:** Garantir que as regras e a configuração do sistema para assinatura digital estão corretas antes de liberar a funcionalidade completa.

---

# **US002 \-  Simular Assinatura Digital com Validação de Parâmetros**

## **Objetivo**

Garantir validação rigorosa de parâmetros antes de simular criação ou validação de assinatura digital.

## **Atores Envolvidos**

| Ator | Descrição |
| :---- | :---- |
| CLI Assinatura | A Interface de Linha de Comando, responsável por receber, validar e rotear a requisição. Implementada referencialmente em Go. |
| Assinador.jar | O componente de backend Java que, de fato, realiza as operações criptográficas de criação e validação de assinatura. |
| Dispositivo criptográfico (token/smart card via PKCS\#11) | Dispositivo físico de segurança que armazena certificados digitais e chaves privadas, utilizando o padrão PKCS\#11 para interface com aplicações de assinatura. |

## 

## **Pré-condições**

Requisição recebida pelo \`assinador.jar\`.  
Parâmetros de entrada fornecidos pelo usuário.

## **Pós-condições**

Em caso válido: Exibir **MS-01**.  
Em caso invalido: Exibir **MS-02**.

## **Fluxo Principal**

**1\. Entrada de Dados e Operação**

* O módulo `Assinador.jar` é invocado via CLI (Command Line Interface), recebendo a operação desejada (`criar` ou `validar`) e os respectivos parâmetros necessários.

**2\. Validação dos Parâmetros**

* O sistema realiza a validação rigorosa dos parâmetros de entrada, verificando:  
  * **Obrigatoriedade:** Se todos os parâmetros requeridos pela operação foram fornecidos.  
  * **Formato:** Se os parâmetros estão no formato esperado (ex: datas, caminhos de arquivo, tipos de dados).  
  * **Consistência:** Se os valores fornecidos são logicamente coerentes entre si e com a operação.

**3\. Execução da Operação**

* **SE a operação for `criar`:**  
  * O sistema gera uma **assinatura simulada** (pré-construída/mock) para fins de teste.  
  * O processo de assinatura real é abstraído ou simulado.  
* **SE a operação for `validar`:**  
  * O sistema executa a regra de validação pré-determinada, verificando se o objeto de entrada atende aos critérios esperados.  
  * O sistema produz um **resultado pré-determinado** (sucesso/falha) baseado nessa regra.

**4\. Retorno da Resposta**  
O sistema constrói e retorna uma **resposta estruturada** (ex: JSON, XML) ao CLI, contendo o resultado da operação executada (assinatura simulada ou resultado da validação).

## **Fluxos Alternativos**

### **FA-01 \- Parametro invalido**

### 1\. No passo 2, a validação falha. 2\. Sistema interrompe processamento. 3\. Sistema retorna erro detalhando parâmetro e motivo.

### **FA-02 \- Falha na interação PKCS\#11**

1\. Sistema retorna erro técnico controlado.  
2\. Não deve haver travamento do processo.

### **FA-03 \- Operação desconhecida**

### 1\. No passo 1, operação diferente de criar/validar. 2\. Sistema retorna erro de operação não suportada.

## 

## **Regras de Negócio (RN)**

| ID | Regra |
| :---- | :---- |
| RN-US002-01 | Nenhuma simulação deve ocorrer sem validação prévia bem-sucedida. |
| RN-US002-02 | Erros devem indicar claramente qual campo falhou e regra violada. |
| RN-US002-03 | Integração PKCS\#11 deve ser opcional no ambiente de simulação, mas com tratamento de falhas. |

##  **Critérios de Aceitação**

\- Todos os parâmetros definidos são validados.  
\- Criação de assinatura retorna payload simulado quando válido.  
\- Validação retorna resultado pré-determinado quando válido.  
\- Erros inválidos são claros e verificáveis.

## **Mensagens Padronizadas**

| Código | Descrição |
| :---- | :---- |
| **MS-01** | Operação concluída com sucesso. |
| **MS-02** | Operação concluída com erro.  |
| **MS-03** | Criação de assinatura simulada concluída com sucesso. |
| **MS-04** | Validação simulada concluída e resultado retornado. |

