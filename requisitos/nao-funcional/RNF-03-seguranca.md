# RNF-03 - Seguranca e Integridade de Artefatos

## Objetivo

Assegurar autenticidade e integridade dos artefatos distribuidos e reduzir riscos de supply chain.

## Escopo

- Publicacao de releases
- Verificacao de artefatos

## Cenario Principal de Conformidade

1. Dado pipeline de release habilitado.
2. Quando nova versao e publicada.
3. Entao cada artefato deve possuir assinatura Cosign e certificado associado.

## Cenarios Alternativos

### CA-01 - Falha na assinatura

- Release deve falhar e nao publicar artefatos sem assinatura valida.

### CA-02 - Verificacao local pelo usuario

- Usuario consegue validar assinatura com `cosign verify-blob` usando `.sig` e `.pem`.

## Regras de Negocio

- RNF03-01: Todo artefato deve ter arquivos correspondentes `.sig` e `.pem`.
- RNF03-02: Assinatura deve usar identidade OIDC e transparency log do Sigstore.
- RNF03-03: Pipeline nao pode promover release com etapa de assinatura falha.

## Criterios de Aceitacao

- 100% dos artefatos de release com assinatura e certificado.
- Verificacao Cosign bem-sucedida para amostras de cada plataforma.
- Evidencias de assinatura registradas no pipeline.

## Rastreabilidade

- Origem: Especificacao, secao 9 (integridade e assinatura de artefatos).
