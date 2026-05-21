package com.runner.assinador;

import com.runner.assinador.domain.SignRequest;
import com.runner.assinador.domain.ValidateRequest;
import com.runner.assinador.domain.SignatureResponse;

public interface SignatureService {
    SignatureResponse sign(SignRequest request);
    SignatureResponse validate(ValidateRequest request);
}
