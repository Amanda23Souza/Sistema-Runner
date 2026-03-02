#!/bin/bash

# Trabalha apenas no diretório onde o script está (e subdiretórios)
SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

FILE="/tmp/plantuml-1.2025.10.jar"
URL="https://github.com/plantuml/plantuml/releases/download/v1.2025.10/plantuml-1.2025.10.jar"

# Verifica se o arquivo existe
if [ -f "$FILE" ]; then
    echo "Plantuml disponível."
else
    echo "O arquivo $FILE não existe. Baixando..."
    http --download "$URL"
    mv plantuml-1.2025.10.jar /tmp
    echo "Download concluído."
fi

echo "Produzindo arquivos SVG a partir de arquivos PlantUML (recursivamente). Aguarde ..."
mkdir -p ./imagens
echo "Diretórios pesquisados (apenas a partir de $SCRIPT_DIR):"
find . -type f -name "*.puml" -not -path "*/imagens/*" -print0 | \
    xargs -0 -n1 dirname | sort -u

# Processa cada arquivo .puml individualmente
# -tsvg gera SVG (determinístico entre plataformas)
# -nometadata remove timestamps para garantir arquivos idênticos
# --output-dir usa caminho absoluto (evita criação de diretórios duplicados)
# Só regenera se o .puml for mais recente que o .svg correspondente
regenerated=0
skipped=0
while IFS= read -r -d '' puml; do
    dir=$(dirname "$puml")
    outdir=$(realpath "$dir/imagens")
    mkdir -p "$outdir"
    
    # Nome do arquivo SVG correspondente
    basename_puml=$(basename "$puml" .puml)
    svg_file="$outdir/$basename_puml.svg"
    
    # Verifica se precisa regenerar (SVG não existe ou PUML mais recente)
    if [ ! -f "$svg_file" ] || [ "$puml" -nt "$svg_file" ]; then
        java -jar "$FILE" -quiet -tsvg -nometadata --output-dir "$outdir" "$puml"
        echo "  [GERADO] $puml"
        ((regenerated++))
    else
        echo "  [SKIP]   $puml (sem alterações)"
        ((skipped++))
    fi
done < <(find . -type f -name "*.puml" -not -path "*/imagens/*" -print0)
echo "Geração concluída: $regenerated arquivo(s) gerado(s), $skipped ignorado(s)."
