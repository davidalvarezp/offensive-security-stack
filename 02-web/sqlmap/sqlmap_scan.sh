#!/bin/bash
# SQL Injection testing template using sqlmap

TARGET=$1
OUTPUT="sqlmap-$TARGET"

if [ -z "$TARGET" ]; then
    echo "Usage: ./sqlmap_scan.sh <url>"
    exit 1
fi

echo "[*] Running sqlmap on $TARGET"
sqlmap -u "$TARGET" --batch --output-dir="$OUTPUT"

echo "[+] Output saved to $OUTPUT"
