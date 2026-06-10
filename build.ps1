# Script para compilar o cli-assinatura com informações de versão no Windows (PowerShell)

$VERSION = try { git describe --tags --always 2>$null } catch { "v0.0.0-dev" }
if (-not $VERSION) { $VERSION = "v0.0.0-dev" }

$COMMIT = try { git rev-parse --short HEAD 2>$null } catch { "unknown" }
if (-not $COMMIT) { $COMMIT = "unknown" }

$BUILDTIME = (Get-Date -uformat "%Y-%m-%dT%H:%M:%SZ")

$LDFLAGS = "-s -w " + `
           "-X 'github.com/Amanda23Souza/Sistema-Runner/cli-assinatura/internal/version.Version=$VERSION' " + `
           "-X 'github.com/Amanda23Souza/Sistema-Runner/cli-assinatura/internal/version.Commit=$COMMIT' " + `
           "-X 'github.com/Amanda23Souza/Sistema-Runner/cli-assinatura/internal/version.BuildTime=$BUILDTIME'"

Write-Host "Compilando cli-assinatura..." -ForegroundColor Cyan

# Entra na pasta do CLI e executa o build
Push-Location cli-assinatura
try {
    go build -ldflags="$LDFLAGS" -o assinatura.exe ./cmd/assinatura
    Write-Host "Compilação concluída com sucesso!" -ForegroundColor Green
    Write-Host "Gerado em: cli-assinatura/assinatura.exe" -ForegroundColor Green
} catch {
    Write-Host "Erro durante a compilação." -ForegroundColor Red
} finally {
    Pop-Location
}
