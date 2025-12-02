#!/bin/bash
# Basic DNS enumeration script

DOMAIN=$1

if [ -z "$DOMAIN" ]; then
    echo "Usage: ./dns_enum.sh <domain>"
    exit 1
fi

echo "[*] Running DNS enumeration on $DOMAIN"

dig $DOMAIN any
dig $DOMAIN mx
dig $DOMAIN ns
dig $DOMAIN txt
dig $DOMAIN soa

echo "[*] Checking zone transfer..."
for ns in $(dig $DOMAIN ns +short); do
    dig axfr $DOMAIN @$ns
done
