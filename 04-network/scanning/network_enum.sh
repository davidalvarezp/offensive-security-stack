#!/bin/bash
# Basic network enumeration using nmap and arp-scan (lab only)

SUBNET=$1

if [ -z "$SUBNET" ]; then
    echo "Usage: ./network_enum.sh <subnet>"
    exit 1
fi

echo "[*] Enumerating live hosts on $SUBNET"
sudo arp-scan -l

echo "[*] Performing ping sweep"
nmap -sn "$SUBNET"

echo "[+] Network enumeration completed"
