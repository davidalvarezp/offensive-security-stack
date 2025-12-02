#!/bin/bash
# Basic TCP port scan using nmap for lab environment

TARGET=$1

if [ -z "$TARGET" ]; then
    echo "Usage: ./port_scan.sh <IP>"
    exit 1
fi

echo "[*] Scanning all ports on $TARGET"
nmap -p- -T4 "$TARGET" -oN "scan_$TARGET.txt"

echo "[+] Scan results saved to scan_$TARGET.txt"
