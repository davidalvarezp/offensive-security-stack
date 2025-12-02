#!/bin/bash
# Network enumeration script (Lab Use Only)
# Requires: scanning/port_scan.sh, scanning/network_enum.sh

SUBNET=$1
TARGET=$2

if [ -z "$SUBNET" ] || [ -z "$TARGET" ]; then
    echo "Usage: ./network_enum.sh <subnet> <target>"
    exit 1
fi

echo "[*] Running network enumeration for subnet $SUBNET and host $TARGET"

bash ./04-network/scanning/network_enum.sh "$SUBNET"
bash ./04-network/scanning/port_scan.sh "$TARGET"

echo "[+] Network enumeration completed."
