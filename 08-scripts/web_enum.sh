#!/bin/bash
# Web enumeration script (Lab Use Only)
# Requires: ffuf, gobuster, nikto, sqlmap

TARGET=$1

if [ -z "$TARGET" ]; then
    echo "Usage: ./web_enum.sh <target_url>"
    exit 1
fi

echo "[*] Starting web enumeration for $TARGET"

echo "[*] Running Gobuster"
gobuster dir -u "$TARGET" -w ./07-wordlists/web/common_paths.txt -o gobuster_$TARGET.txt

echo "[*] Running FFUF"
ffuf -w ./07-wordlists/web/fuzz_template.txt -u "$TARGET/FUZZ" -o ffuf_$TARGET.json

echo "[*] Running Nikto scan"
nikto -h "$TARGET" -o nikto_$TARGET.txt

echo "[*] Web enumeration completed."
