set shell := ["bash", "-cu"]

bin := "alfred-aws-icons"
dist_file := "aws-icons.alfredworkflow"

# Run all asset preprocessing steps
setup-assets: fix-spaces remove-unused resize

# Fix filenames containing spaces
fix-spaces:
    #!/usr/bin/env bash
    set -euo pipefail
    find assets -name "* *" -print0 | while IFS= read -r -d '' f; do
        dir="$(dirname "$f")"
        base="$(basename "$f")"
        newbase="$(echo "$base" | sed 's/ //g')"
        if [ "$base" != "$newbase" ]; then
            echo "Renaming: $f -> $dir/$newbase"
            mv "$f" "$dir/$newbase"
        fi
    done

# Remove unused sizes (16,32,48) and @5x files
remove-unused:
    rm -rf assets/Architecture-Service-Icons/Arch_*/{16,32,48}
    find assets/Architecture-Service-Icons -name "*@5x.png" -delete

# Resize Resource-Icons SVGs from 48 to 64
resize:
    #!/usr/bin/env bash
    set -euo pipefail
    # Standard Resource-Icons
    for src in assets/Resource-Icons/Res_*/*_48.svg; do
        [ -f "$src" ] || continue
        dst="${src/_48.svg/_64.svg}"
        rsvg-convert --format svg -h 64 -w 64 -o "$dst" "$src"
        rm "$src"
    done
    # General-Icons (Dark/Light)
    for variant in Dark Light; do
        for src in assets/Resource-Icons/Res_General-Icons/Res_48_${variant}/*_48_${variant}.svg; do
            [ -f "$src" ] || continue
            dst="${src/_48_${variant}.svg/_64_${variant}.svg}"
            rsvg-convert --format svg -h 64 -w 64 -o "$dst" "$src"
            rm "$src"
        done
    done

# Build binary
build:
    go build -o {{ bin }} ./main.go

# Package as .alfredworkflow
dist: build
    zip -r {{ dist_file }} {{ bin }} info.plist icon.png assets

# Clean build artifacts
clean:
    rm -f {{ bin }} {{ dist_file }}
