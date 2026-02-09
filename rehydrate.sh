#!/bin/bash
# Rehydrate helper script
# Usage: ./rehydrate.sh <keyword>

MEMORIES_DIR="/opt/memories"

if [ -z "$1" ]; then
    echo "Usage: ./rehydrate.sh <keyword>"
    echo "Available memories:"
    ls -1 "$MEMORIES_DIR" | sed 's/^/  - /'
    exit 1
fi

KEYWORD=$1
MATCHES=$(ls "$MEMORIES_DIR" | grep -i "$KEYWORD")

if [ -z "$MATCHES" ]; then
    echo "No memories found for '$KEYWORD'."
    exit 1
fi

COUNT=$(echo "$MATCHES" | wc -l)

if [ "$COUNT" -eq 1 ]; then
    cat "$MEMORIES_DIR/$MATCHES"
else
    echo "Multiple matches found. Please be more specific:"
    echo "$MATCHES" | sed 's/^/  - /'
fi
