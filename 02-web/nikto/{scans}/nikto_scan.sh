#!/bin/bash
# Web server vulnerability scanning using Nikto

TARGET=$1
OUTPUT="nikto-$TARGET.txt"

if [ -z "$TARGET" ]; then
    echo "Usage: ./nikto_scan.sh <url>"
    exit 1
fi

echo "[*] Running Nikto on $TARGET"
nikto -h "$TARGET" -o "$OUTPUT" -Format txt

echo "[+] Results saved to $OUTPUT"
