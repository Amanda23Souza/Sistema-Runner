package com.kyriosdata.assinador.domain;

public class SignatureResponse {
    private String signature;
    private boolean valid;
    private String message;

    public SignatureResponse() {}

    public SignatureResponse(String signature, boolean valid, String message) {
        this.signature = signature;
        this.valid = valid;
        this.message = message;
    }

    public String getSignature() {
        return signature;
    }

    public void setSignature(String signature) {
        this.signature = signature;
    }

    public boolean isValid() {
        return valid;
    }

    public void setValid(boolean valid) {
        this.valid = valid;
    }

    public String getMessage() {
        return message;
    }

    public void setMessage(String message) {
        this.message = message;
    }
}
