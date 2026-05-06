# Assinatura CLI

Ferramenta CLI desenvolvida em Go para realizar assinatura digital simulada e validação de artefatos.

## Releases e Downloads

O CLI está disponível para download e é empacotado individualmente para os sistemas `windows`, `linux` e `darwin` na página de [Releases](../../releases) deste repositório.

## Como Gerar Novos Artefatos / Releases

A geração de novos executáveis é 100% automatizada pelo pipeline de CI/CD via GitHub Actions. Para gerar e publicar uma nova versão:
1. Edite o arquivo `cmd/version.go` e altere a variável `var version = "x.y.z"`.
2. Faça o *commit* e envie (*push*) para a branch `main`.
3. O pipeline compilará as 3 plataformas automaticamente, anexará os `checksums` SHA256, assinará o pacote usando o *Cosign* e fará o deploy completo no [GitHub Releases](../../releases).

## Verificando Autenticidade com Cosign (Supply Chain Security)

Todas as nossas releases são assinadas digitalmente via **Sigstore / Cosign** usando a modalidade *keyless* (sem chaves longas) atrelada via certificado OIDC do GitHub Actions.
Para verificar os executáveis que você baixou na página de Releases, utilize as ferramentas de linha de comando exigidas.

**1. Verifique a Hash SHA256 (Opcional):**
```bash
sha256sum -c assinatura-vX.Y.Z-<os>-<arch>.sha256
```

**2. Verifique o Certificado Sigstore (Cosign):**
```bash
cosign verify-blob \
  --certificate assinatura-vX.Y.Z-<os>-<arch>.pem \
  --signature assinatura-vX.Y.Z-<os>-<arch>.sig \
  --certificate-identity "https://github.com/kyriosdata/runner/.github/workflows/ci.yml@refs/heads/main" \
  --certificate-oidc-issuer "https://token.actions.githubusercontent.com" \
  assinatura-vX.Y.Z-<os>-<arch>.exe
```

Se a saída for **"Verified OK"**, significa que o pacote veio infalivelmente do nosso pipeline do GitHub e não sofreu alterações no caminho.
