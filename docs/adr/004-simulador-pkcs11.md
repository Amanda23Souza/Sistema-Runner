# ADR-004 — Simulador PKCS#11 via FakeSignatureService

**Data:** 2026-06-09
**Status:** Aceito
**Autores:** Amanda Soares, Marcello Ronald
**Referência à especificação:** [`kyriosdata/runner @ d3f1a9c`](https://github.com/kyriosdata/runner/blob/d3f1a9c/docs/runner.md)

---

## Contexto

A especificação do Runner prevê integração com dispositivos criptográficos via PKCS#11 (tokens USB, smart cards). O critério E5 exige um "Simulador PKCS11 com testes de integração".

No contexto acadêmico deste projeto, a integração real com hardware criptográfico apresenta desafios:
1. **Hardware indisponível:** Membros da equipe não possuem tokens USB de assinatura digital.
2. **Complexidade:** A API PKCS#11 é de baixo nível (interface C com bindings JNI/IAIK) e requer bibliotecas nativas específicas por plataforma.
3. **SoftHSM:** Alternativas de software (SoftHSM2) requerem instalação e configuração complexa no CI.

## Decisão

Adotar o `FakeSignatureService` como **simulador de assinatura** que substitui a camada PKCS#11 durante o desenvolvimento e testes.

O simulador:
- Aceita qualquer conteúdo e retorna uma assinatura constante (`MOCKED_SIGNATURE_BASE64_==`).
- Valida assinaturas comparando com o valor fixo esperado.
- Implementa a mesma interface (`SignatureService`) que uma eventual implementação real usaria.

## Alternativas Consideradas

| Alternativa | Motivo de descarte |
|-------------|-------------------|
| SoftHSM2 + sun.security.pkcs11 | Instalação complexa no CI; binários nativos por plataforma; overhead de configuração desproporcional ao escopo acadêmico |
| Mock completo da API PKCS#11 (JNI) | Requer binding nativo; altamente complexo; não agrega valor significativo à validação da integração CLI↔JAR |
| Certificado autoassinado via JCA | Mais realista, mas o foco do projeto é a integração CLI↔JAR, não a criptografia em si |

## Consequências

- **Positivas:** Testes determinísticos e rápidos; zero dependências nativas; foco mantido na integração CLI↔HTTP↔JAR que é o objetivo principal.
- **Negativas:** Não exercita criptografia real; migração para PKCS#11 real requer implementação de `SignatureService` com bindings PKCS#11 e configuração de provedor de segurança Java.

## Caminho para PKCS#11 Real

Quando o hardware estiver disponível, a migração consiste em:
1. Implementar uma classe `Pkcs11SignatureService implements SignatureService` usando `sun.security.pkcs11.SunPKCS11`.
2. Configurar o provedor PKCS#11 via arquivo `pkcs11.cfg` apontando para a biblioteca nativa do token.
3. Injetar a implementação no `App.java` via configuração ou parâmetro CLI.
4. Manter `FakeSignatureService` como fallback para testes sem hardware.
