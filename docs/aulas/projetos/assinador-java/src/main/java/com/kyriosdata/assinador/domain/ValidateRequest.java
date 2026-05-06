package com.kyriosdata.assinador.domain;

/**
 * Representa os dados da requisição para operação de validação de assinatura digital.
 * 
 * <p>Esta classe atua como um DTO para transportar informações de validação de assinatura.</p>
 */
public class ValidateRequest {
    
    /**
     * O conteúdo original da assinatura que será validada.
     */
    private String content;

    /**
     * A assinatura digital pré-existente (ex. em Base64).
     */
    private String signature;

    public ValidateRequest() {}

    public String getContent() {
        return content;
    }

    public void setContent(String content) {
        this.content = content;
    }

    public String getSignature() {
        return signature;
    }

    public void setSignature(String signature) {
        this.signature = signature;
    }
}
