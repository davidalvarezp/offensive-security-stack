#!/bin/bash
# Directory brute-forcing using Gobuster

TARGET=$1
WORDLIST=$2
OUTPUT="gobuster-$TARGET.txt"

if [ -z "$TARGET" ] || [ -z "$WORDLIST" ]; then
    echo "Usage: ./dir_scan.sh <url> <wordlist>"
    exit 1
fi

echo "[*] Running Gobuster on $TARGET"
gobuster dir -u "$TARGET" -w "$WORDLIST" -o "$OUTPUT"

echo "[+] Results saved to $OUTPUT"
