package com.kyriosdata.assinador;

import io.javalin.Javalin;

public class App {

    public static void main(String[] args) {
        int port = 8080;
        if (args.length > 0) {
            try {
                port = Integer.parseInt(args[0]);
            } catch (NumberFormatException e) {
                System.err.println("Porta inválida: " + args[0] + ". Usando porta padrão 8080.");
            }
        }

        SignatureService signatureService = new FakeSignatureService();
        SignatureController controller = new SignatureController(signatureService);

        Javalin app = Javalin.create()
            .post("/sign", controller::sign)
            .post("/validate", controller::validate)
            .start(port);
            
        System.out.println("Assinador iniciado na porta " + port);
    }
}
