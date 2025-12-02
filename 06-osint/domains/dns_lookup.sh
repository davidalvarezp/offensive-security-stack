#!/bin/bash
# DNS info gathering (legal OSINT)

DOMAIN=$1

if [ -z "$DOMAIN" ]; then
    echo "Usage: ./dns_lookup.sh <domain>"
    exit 1
fi

echo "[*] DNS Lookup for $DOMAIN"
dig "$DOMAIN" ANY
dig "$DOMAIN" TXT
nslookup "$DOMAIN"
