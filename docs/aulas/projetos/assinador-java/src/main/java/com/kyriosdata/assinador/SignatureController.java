package com.kyriosdata.assinador;

import com.kyriosdata.assinador.domain.SignRequest;
import com.kyriosdata.assinador.domain.SignatureResponse;
import com.kyriosdata.assinador.domain.ValidateRequest;
import io.javalin.http.Context;
import io.javalin.http.HttpStatus;

public class SignatureController {

    private final SignatureService signatureService;

    public SignatureController(SignatureService signatureService) {
        this.signatureService = signatureService;
    }

    public void sign(Context ctx) {
        try {
            SignRequest request = ctx.bodyAsClass(SignRequest.class);
            SignatureResponse response = signatureService.sign(request);
            if (response.isValid()) {
                ctx.status(HttpStatus.OK).json(response);
            } else {
                ctx.status(HttpStatus.BAD_REQUEST).json(response);
            }
        } catch (Exception e) {
            ctx.status(HttpStatus.BAD_REQUEST).json(new SignatureResponse(null, false, "Requisição inválida: formato incorreto"));
        }
    }

    public void validate(Context ctx) {
        try {
            ValidateRequest request = ctx.bodyAsClass(ValidateRequest.class);
            SignatureResponse response = signatureService.validate(request);
            if (response.isValid()) {
                ctx.status(HttpStatus.OK).json(response);
            } else {
                // We return OK or BAD_REQUEST for invalid signatures?
                // The task description says "Respostas HTTP seguem estrutura consistente (sucesso e erro)"
                // Often a validation returning "invalid" is still a successful HTTP 200 response
                // containing a JSON with `valid=false`. 
                // But if parameters are invalid, we can return 400.
                // Looking at FakeSignatureService, it returns valid=false when parameters are invalid AND when signature is invalid.
                // We will just return 200 with the response, and 400 when the JSON is malformed.
                // Wait, if content/signature is empty, FakeSignatureService returns valid=false.
                ctx.status(HttpStatus.OK).json(response);
            }
        } catch (Exception e) {
            ctx.status(HttpStatus.BAD_REQUEST).json(new SignatureResponse(null, false, "Requisição inválida: formato incorreto"));
        }
    }
}
