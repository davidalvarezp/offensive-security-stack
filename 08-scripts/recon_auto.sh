#!/bin/bash
# Automated Recon Script (Lab Use Only)
# Combines tools from 01-recon and 04-network/scanning

TARGET=$1
SUBNET=$2

if [ -z "$TARGET" ] || [ -z "$SUBNET" ]; then
    echo "Usage: ./recon_auto.sh <target> <subnet>"
    exit 1
fi

echo "[*] Starting automated recon for $TARGET"

# Run recon tools
echo "[*] Running Nmap scan"
nmap -sC -sV -oN recon_nmap_$TARGET.txt "$TARGET"

echo "[*] Running DNS enumeration"
bash ./01-recon/dns/dns_enum.sh "$TARGET"

echo "[*] Running masscan (lab only)"
masscan "$SUBNET" -p1-65535 --rate=1000 -oG masscan_$SUBNET.txt

echo "[+] Automated recon completed. Check output files."
