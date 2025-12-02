#!/bin/bash
# Kerberos basic enumeration (without exploitation)

DOMAIN=$1
DC_IP=$2

if [ -z "$DOMAIN" ] || [ -z "$DC_IP" ]; then
    echo "Usage: ./enum_kerberos.sh <domain> <domain controller IP>"
    exit 1
fi

echo "[*] Kerberos Enumeration on $DOMAIN ($DC_IP)"

nmap --script krb5-enum-users --script-args krb5-enum-users.realm=$DOMAIN -p 88 $DC_IP

echo "[+] Kerberos enumeration completed."
