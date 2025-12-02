#!/bin/bash
# WHOIS lookup script (legal OSINT)

DOMAIN=$1

if [ -z "$DOMAIN" ]; then
    echo "Usage: ./whois_lookup.sh <domain>"
    exit 1
fi

whois "$DOMAIN"
