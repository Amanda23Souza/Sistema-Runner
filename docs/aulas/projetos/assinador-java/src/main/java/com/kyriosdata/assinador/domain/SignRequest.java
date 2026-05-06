package com.kyriosdata.assinador.domain;

/**
 * Representa os dados da requisição para operação real de criação de assinatura digital.
 * 
 * <p>Esta classe atua como um DTO para transportar informações de criação de assinatura.</p>
 */
public class SignRequest {
    
    /**
     * O conteúdo original que se deseja assinar.
     */
    private String content;

    /**
     * Token de autenticação, credencial ou PIN opcional.
     */
    private String token;

    public SignRequest() {}

    public String getContent() {
        return content;
    }

    public void setContent(String content) {
        this.content = content;
    }

    public String getToken() {
        return token;
    }

    public void setToken(String token) {
        this.token = token;
    }
}
