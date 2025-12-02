#!/bin/bash
# AD Enumeration Script (Lab Use Only)
# Requires: enum_smb.sh, enum_ldap.sh, enum_rpc.sh, enum_kerberos.sh from 03-ad/enum

TARGET=$1
DOMAIN=$2
DC_IP=$3

if [ -z "$TARGET" ] || [ -z "$DOMAIN" ] || [ -z "$DC_IP" ]; then
    echo "Usage: ./ad_enum.sh <target> <domain> <dc_ip>"
    exit 1
fi

echo "[*] Starting Active Directory enumeration for $TARGET"

bash ./03-ad/enum/enum_smb.sh "$TARGET"
bash ./03-ad/enum/enum_ldap.sh "$TARGET"
bash ./03-ad/enum/enum_rpc.sh "$TARGET"
bash ./03-ad/enum/enum_kerberos.sh "$DOMAIN" "$DC_IP"

echo "[+] AD enumeration completed."
