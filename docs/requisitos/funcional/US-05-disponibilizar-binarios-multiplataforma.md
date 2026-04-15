# US-05 - Disponibilizar Binarios Multiplataforma

## Objetivo

Publicar binarios pre-compilados do CLI para Windows, Linux e macOS via GitHub Releases com integridade verificavel.

## Atores

- Equipe de desenvolvimento
- Pipeline CI/CD
- Usuario final

## Pre-condicoes

- Build e testes da versao aprovados.
- Pipeline de release configurado.

## Pos-condicoes

- Release publicada com artefatos por plataforma, checksums e versao SemVer.

## Fluxo Principal (Caminho Feliz)

1. Equipe define nova versao seguindo SemVer.
2. Pipeline gera binarios para Windows, Linux e macOS (amd64).
3. Pipeline calcula e publica checksums SHA256.
4. Pipeline publica artefatos na release do GitHub.
5. Pipeline finaliza release com metadados de versao e evidencias.

## Fluxos Alternativos

### FA-01 - Falha de build em uma plataforma

1. No passo 2, build falha em plataforma especifica.
2. Pipeline interrompe release e reporta logs.

### FA-02 - Checksum nao gerado

1. No passo 3, geracao de checksum falha.
2. Release nao deve ser publicada parcialmente.

### FA-03 - Publicacao interrompida

1. No passo 4, erro de autenticacao/permissao no GitHub.
2. Pipeline marca release como falha e evita estado inconsistente.

## Regras de Negocio

- RN-US05-01: Toda release deve usar versionamento semantico.
- RN-US05-02: Toda release deve conter binarios para as tres plataformas alvo.
- RN-US05-03: Toda release deve incluir checksums SHA256.
- RN-US05-04: Artefatos devem ser assinados conforme requisito de integridade (Cosign/Sigstore).

## Criterios de Aceitacao

- Binarios para Windows/Linux/macOS publicados.
- Checksums SHA256 disponiveis.
- SemVer aplicado em tags e artefatos.
- Release validada no GitHub.

## Rastreabilidade

- Origem: Especificacao, secao 5, US-05 e secao 9 (integridade e assinatura).
