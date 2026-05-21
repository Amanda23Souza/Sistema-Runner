package com.runner.assinador;

import com.fasterxml.jackson.databind.ObjectMapper;
import com.runner.assinador.domain.SignRequest;
import com.runner.assinador.domain.SignatureResponse;
import com.runner.assinador.domain.ValidateRequest;
import io.javalin.Javalin;
import org.junit.jupiter.api.AfterAll;
import org.junit.jupiter.api.BeforeAll;
import org.junit.jupiter.api.Test;

import java.net.URI;
import java.net.http.HttpClient;
import java.net.http.HttpRequest;
import java.net.http.HttpResponse;

import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;
import static org.junit.jupiter.api.Assertions.assertFalse;

class SignatureControllerTest {

    private static Javalin app;
    private static int port;
    private static final HttpClient client = HttpClient.newHttpClient();
    private static final ObjectMapper mapper = new ObjectMapper();

    @BeforeAll
    static void setup() {
        SignatureService signatureService = new FakeSignatureService();
        SignatureController controller = new SignatureController(signatureService);

        app = Javalin.create()
            .get("/health", ctx -> ctx.status(200).json(java.util.Map.of("status", "UP")))
            .post("/sign", controller::sign)
            .post("/validate", controller::validate)
            .start(0); // Random available port
            
        port = app.port();
    }

    @AfterAll
    static void tearDown() {
        if (app != null) {
            app.stop();
        }
    }

    @Test
    void testSignEndpointSuccess() throws Exception {
        SignRequest req = new SignRequest();
        req.setContent("test content");
        
        String json = mapper.writeValueAsString(req);
        
        HttpRequest request = HttpRequest.newBuilder()
            .uri(URI.create("http://localhost:" + port + "/sign"))
            .header("Content-Type", "application/json")
            .POST(HttpRequest.BodyPublishers.ofString(json))
            .build();
            
        HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());
        
        assertEquals(200, response.statusCode());
        SignatureResponse res = mapper.readValue(response.body(), SignatureResponse.class);
        assertTrue(res.isValid());
        assertEquals("MOCKED_SIGNATURE_BASE64_==", res.getSignature());
    }

    @Test
    void testValidateEndpointSuccess() throws Exception {
        ValidateRequest req = new ValidateRequest();
        req.setContent("test content");
        req.setSignature("MOCKED_SIGNATURE_BASE64_==");
        
        String json = mapper.writeValueAsString(req);
        
        HttpRequest request = HttpRequest.newBuilder()
            .uri(URI.create("http://localhost:" + port + "/validate"))
            .header("Content-Type", "application/json")
            .POST(HttpRequest.BodyPublishers.ofString(json))
            .build();
            
        HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());
        
        assertEquals(200, response.statusCode());
        SignatureResponse res = mapper.readValue(response.body(), SignatureResponse.class);
        assertTrue(res.isValid());
        assertEquals("MOCKED_SIGNATURE_BASE64_==", res.getSignature());
    }

    @Test
    void testValidateEndpointInvalidSignature() throws Exception {
        ValidateRequest req = new ValidateRequest();
        req.setContent("test content");
        req.setSignature("WRONG_SIGNATURE");
        
        String json = mapper.writeValueAsString(req);
        
        HttpRequest request = HttpRequest.newBuilder()
            .uri(URI.create("http://localhost:" + port + "/validate"))
            .header("Content-Type", "application/json")
            .POST(HttpRequest.BodyPublishers.ofString(json))
            .build();
            
        HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());
        
        assertEquals(200, response.statusCode());
        SignatureResponse res = mapper.readValue(response.body(), SignatureResponse.class);
        assertFalse(res.isValid());
        assertEquals("Assinatura é inválida", res.getMessage());
    }

    @Test
    void testSignEndpointMissingContent() throws Exception {
        SignRequest req = new SignRequest();
        
        String json = mapper.writeValueAsString(req);
        
        HttpRequest request = HttpRequest.newBuilder()
            .uri(URI.create("http://localhost:" + port + "/sign"))
            .header("Content-Type", "application/json")
            .POST(HttpRequest.BodyPublishers.ofString(json))
            .build();
            
        HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());
        
        // As per our implementation, it returns BAD_REQUEST (400) when validation fails
        assertEquals(400, response.statusCode());
        SignatureResponse res = mapper.readValue(response.body(), SignatureResponse.class);
        assertFalse(res.isValid());
        assertEquals("Parâmetro 'content' inválido ou ausente", res.getMessage());
    }

    @Test
    void testHealthCheckSuccess() throws Exception {
        HttpRequest request = HttpRequest.newBuilder()
            .uri(URI.create("http://localhost:" + port + "/health"))
            .GET()
            .build();
            
        HttpResponse<String> response = client.send(request, HttpResponse.BodyHandlers.ofString());
        
        assertEquals(200, response.statusCode());
        assertTrue(response.body().contains("\"status\":\"UP\""));
    }
}
