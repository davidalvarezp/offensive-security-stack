#!/bin/bash
# SNMP enumeration in lab environments

TARGET=$1
COMMUNITY=$2

if [ -z "$TARGET" ] || [ -z "$COMMUNITY" ]; then
    echo "Usage: ./snmp_enum.sh <IP> <community>"
    exit 1
fi

echo "[*] Enumerating SNMP on $TARGET"
snmpwalk -v2c -c $COMMUNITY $TARGET

echo "[+] SNMP enumeration completed."
