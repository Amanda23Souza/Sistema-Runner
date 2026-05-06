package com.kyriosdata.assinador;

import com.kyriosdata.assinador.domain.SignRequest;
import com.kyriosdata.assinador.domain.ValidateRequest;
import com.kyriosdata.assinador.domain.SignatureResponse;
import org.junit.jupiter.api.Test;

import static org.junit.jupiter.api.Assertions.*;

class FakeSignatureServiceTest {

    private final SignatureService service = new FakeSignatureService();

    @Test
    void shouldReturnSimulatedSignatureForValidSignRequest() {
        SignRequest request = new SignRequest();
        request.setContent("sample content");
        
        SignatureResponse response = service.sign(request);
        
        assertNotNull(response);
        assertTrue(response.isValid());
        assertEquals("MOCKED_SIGNATURE_BASE64_==", response.getSignature());
        assertEquals("Assinatura criada com sucesso", response.getMessage());
    }

    @Test
    void shouldReturnErrorForInvalidSignRequest() {
        SignRequest request = new SignRequest();
        request.setContent(null); // Invalid
        
        SignatureResponse response = service.sign(request);
        
        assertNotNull(response);
        assertFalse(response.isValid());
        assertNull(response.getSignature());
        assertEquals("Parâmetro 'content' inválido ou ausente", response.getMessage());
    }

    @Test
    void shouldValidateSimulatedSignatureCorrectly() {
        ValidateRequest request = new ValidateRequest();
        request.setContent("sample content");
        request.setSignature("MOCKED_SIGNATURE_BASE64_==");
        
        SignatureResponse response = service.validate(request);
        
        assertNotNull(response);
        assertTrue(response.isValid());
        assertEquals("Assinatura é válida", response.getMessage());
    }

    @Test
    void shouldReturnErrorForInvalidSimulatedSignature() {
        ValidateRequest request = new ValidateRequest();
        request.setContent("sample content");
        request.setSignature("BAD_SIGNATURE_==");
        
        SignatureResponse response = service.validate(request);
        
        assertNotNull(response);
        assertFalse(response.isValid());
        assertEquals("Assinatura é inválida", response.getMessage());
    }
}
