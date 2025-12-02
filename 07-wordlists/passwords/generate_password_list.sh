#!/bin/bash
# Generates simple password variations for lab use.

BASE=$1

if [ -z "$BASE" ]; then
    echo "Usage: ./generate_password_list.sh <baseword>"
    exit 1
fi

echo "[*] Generating password variations for: $BASE"

echo "$BASE"              >  "$BASE-wordlist.txt"
echo "${BASE}123"        >> "$BASE-wordlist.txt"
echo "${BASE}2024"       >> "$BASE-wordlist.txt"
echo "${BASE}!"          >> "$BASE-wordlist.txt"
echo "${BASE}@123"       >> "$BASE-wordlist.txt"
echo "${BASE}_lab"       >> "$BASE-wordlist.txt"

echo "[+] Generated: $BASE-wordlist.txt"
