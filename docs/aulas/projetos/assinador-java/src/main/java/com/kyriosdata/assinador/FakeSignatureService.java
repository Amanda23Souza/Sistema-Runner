package com.kyriosdata.assinador;

import com.kyriosdata.assinador.domain.SignRequest;
import com.kyriosdata.assinador.domain.ValidateRequest;
import com.kyriosdata.assinador.domain.SignatureResponse;

public class FakeSignatureService implements SignatureService {

    private static final String FAKE_SIGNATURE = "MOCKED_SIGNATURE_BASE64_==";

    @Override
    public SignatureResponse sign(SignRequest request) {
        if (request == null || request.getContent() == null || request.getContent().isEmpty()) {
            return new SignatureResponse(null, false, "Parâmetro 'content' inválido ou ausente");
        }
        return new SignatureResponse(FAKE_SIGNATURE, true, "Assinatura criada com sucesso");
    }

    @Override
    public SignatureResponse validate(ValidateRequest request) {
        if (request == null || request.getContent() == null || request.getContent().isEmpty()) {
            return new SignatureResponse(null, false, "Parâmetro 'content' inválido ou ausente");
        }
        if (request.getSignature() == null || request.getSignature().isEmpty()) {
            return new SignatureResponse(null, false, "Parâmetro 'signature' inválido ou ausente");
        }
        
        boolean isValid = FAKE_SIGNATURE.equals(request.getSignature());
        return new SignatureResponse(request.getSignature(), isValid, isValid ? "Assinatura é válida" : "Assinatura é inválida");
    }
}
