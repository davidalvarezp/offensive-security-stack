#!/bin/bash
# Basic directory fuzzing using ffuf

TARGET=$1
WORDLIST=$2
OUTPUT="ffuf-$TARGET.txt"

if [ -z "$TARGET" ] || [ -z "$WORDLIST" ]; then
    echo "Usage: ./ffuf_scan.sh <url> <wordlist>"
    exit 1
fi

echo "[*] Running ffuf on $TARGET"
ffuf -w "$WORDLIST" -u "$TARGET/FUZZ" -o "$OUTPUT"

echo "[+] Results saved to $OUTPUT"
