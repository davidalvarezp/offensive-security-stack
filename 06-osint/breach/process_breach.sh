#!/bin/bash
# Parses breach CSV files (lab only)
# No real breach data should ever be added to this repository.

FILE=$1

if [ -z "$FILE" ]; then
    echo "Usage: ./process_breach.sh <breach_file.csv>"
    exit 1
fi

echo "[*] Counting unique emails..."
cut -d',' -f1 "$FILE" | sort -u | wc -l

echo "[*] Top repeated passwords (lab simulation)"
cut -d',' -f2 "$FILE" | sort | uniq -c | sort -nr | head

echo "[+] Breach stats generated."
