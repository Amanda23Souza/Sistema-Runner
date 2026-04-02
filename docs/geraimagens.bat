@echo off
chcp 65001

:: Trabalha apenas no diretório onde o script está (e subdiretórios)
pushd "%~dp0"

set FILE=%TEMP%\plantuml-1.2025.10.jar
set URL=https://github.com/plantuml/plantuml/releases/download/v1.2025.10/plantuml-1.2025.10.jar

:: Opção --force para regenerar todos os arquivos
set "FORCE=0"
if "%1"=="--force" set "FORCE=1"
if "%1"=="-f" set "FORCE=1"

:: Verifica se o arquivo existe
if exist "%FILE%" (
    echo Plantuml disponível.
) else (
    echo O arquivo %FILE% não existe. Baixando...
    powershell -Command "Invoke-WebRequest -Uri %URL% -OutFile '%FILE%'"
    echo Download concluído.
)

if "%FORCE%"=="1" (
    echo Produzindo TODOS os arquivos SVG a partir de arquivos PlantUML --force. Aguarde ...
) else (
    echo Produzindo arquivos SVG apenas para puml modificados. Use --force para regenerar todos.
)

setlocal enabledelayedexpansion
set "LASTDIR="
for /r %%f in (*.puml) do (
    rem Ignora qualquer arquivo dentro de diretórios chamados "imagens"
    set "FILEPATH=%%~f"
    echo !FILEPATH! | findstr /I /L /C:"\imagens\" >nul
    if errorlevel 1 (
        set "CURRDIR=%%~dpf"
        rem Diretório de saída é imagens/ dentro do diretório do arquivo
        set "OUTDIR=%%~dpfimagens"
        if not exist "!OUTDIR!" mkdir "!OUTDIR!"
        
        if /I not "!CURRDIR!"=="!LASTDIR!" (
            echo Diretório: !CURRDIR!
            set "LASTDIR=!CURRDIR!"
        )
        rem Verifica se SVG precisa ser regenerado
        set "SVGFILE=!OUTDIR!\%%~nf.svg"
        set "NEEDSGEN=1"
        set "PUMLFILE=%%f"
        if "!FORCE!"=="0" (
            if exist "!SVGFILE!" (
                rem Compara timestamps usando PowerShell
                for /f %%r in ('powershell -NoProfile -Command "if ((Get-Item -LiteralPath '!PUMLFILE!').LastWriteTime -le (Get-Item -LiteralPath '!SVGFILE!').LastWriteTime) { 'skip' } else { 'gen' }"') do (
                    if "%%r"=="skip" set "NEEDSGEN=0"
                )
            )
        )
        if "!NEEDSGEN!"=="1" (
            echo   Gerando: %%~nf.svg
            java -jar "%FILE%" -tsvg -nometadata -quiet -o "!OUTDIR!" "%%f"
        ) else (
            echo   Pulando: %%~nf.svg [sem alterações]
        )
    )
)
endlocal
echo Geração de arquivos concluída.
