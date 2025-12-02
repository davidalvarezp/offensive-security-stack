#!/bin/bash
# BloodHound data collection (Safe & Legal Labs Only)

TARGET_DC=$1

if [ -z "$TARGET_DC" ]; then
    echo "Usage: ./bh_collect.sh <DOMAIN_CONTROLLER_IP>"
    exit 1
fi

echo "[*] Collecting BloodHound data..."

# SharpHound collection example (manually upload binary in lab contexts)
# ./SharpHound.exe -c All -d domain.local -dc $TARGET_DC -o output.zip

echo "[!] NOTE: This is only a template. 
SharpHound binaries are NOT included for security and legal reasons."
