#!/bin/bash
# Passive/Active subdomain enumeration using amass

DOMAIN=$1
OUTPUT="results-$DOMAIN.txt"

if [ -z "$DOMAIN" ]; then
    echo "Usage: ./amass_scan.sh <domain>"
    exit 1
fi

echo "[*] Running Amass enumeration on $DOMAIN"
amass enum -d "$DOMAIN" -o "$OUTPUT"

echo "[+] Output saved to $OUTPUT"
