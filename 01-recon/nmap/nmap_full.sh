#!/bin/bash
# Full TCP scan with version detection

TARGET=$1
OUT="nmap_full-$TARGET.txt"

if [ -z "$TARGET" ]; then
    echo "Usage: ./nmap_full.sh <IP>"
    exit 1
fi

echo "[*] Running full Nmap scan on $TARGET"
nmap -p- -sV -sC -O -T4 -oN "$OUT" "$TARGET"

echo "[+] Output saved to $OUT"
