#!/bin/bash
# Quick scan for faster results

TARGET=$1
OUT="nmap_quick-$TARGET.txt"

if [ -z "$TARGET" ]; then
    echo "Usage: ./nmap_quick.sh <IP>"
    exit 1
fi

echo "[*] Running quick scan..."
nmap -T4 -F -oN "$OUT" "$TARGET"

echo "[+] Output saved to $OUT"
