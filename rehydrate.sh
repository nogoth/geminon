#!/bin/bash
# Rehydrate helper script
# Usage: ./rehydrate.sh <keyword>
# Matches the keyword against memory filenames AND file contents
# (frontmatter tags and body text are searchable).

MEMORIES_DIR="$(dirname "$0")/memories"

if [ -z "$1" ]; then
    echo "Usage: ./rehydrate.sh <keyword>"
    echo "Available memories:"
    ls -1 "$MEMORIES_DIR" | sed 's/^/  - /'
    exit 1
fi

KEYWORD=$1

MATCHES=$(
    {
        ls -1 "$MEMORIES_DIR" | grep -i -- "$KEYWORD"
        grep -ril -- "$KEYWORD" "$MEMORIES_DIR" | xargs -r -n1 basename
    } | sort -u
)

if [ -z "$MATCHES" ]; then
    echo "No memories found for '$KEYWORD'."
    exit 1
fi

COUNT=$(echo "$MATCHES" | wc -l)

if [ "$COUNT" -eq 1 ]; then
    cat "$MEMORIES_DIR/$MATCHES"
else
    echo "Multiple matches found. Please be more specific:"
    while IFS= read -r name; do
        printf '  - %s\n' "$name"
        grep -im1 -- "$KEYWORD" "$MEMORIES_DIR/$name" | sed 's/^/      /'
    done <<< "$MATCHES"
fi
