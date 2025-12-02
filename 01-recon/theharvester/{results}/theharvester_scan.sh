#!/bin/bash
# Email, subdomain & employer enumeration using theHarvester

DOMAIN=$1
SOURCE="all"
OUT="theharvester-$DOMAIN.txt"

if [ -z "$DOMAIN" ]; then
    echo "Usage: ./theharvester_scan.sh <domain>"
    exit 1
fi

echo "[*] Running theHarvester on $DOMAIN"
theHarvester -d "$DOMAIN" -b "$SOURCE" -f "$OUT"

echo "[+] Output saved to $OUT"
