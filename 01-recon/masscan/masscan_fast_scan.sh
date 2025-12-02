#!/bin/bash
# Fast Internet-range port scan using Masscan

TARGET=$1
RATE=1000
OUT="masscan-$TARGET.json"

if [ -z "$TARGET" ]; then
    echo "Usage: ./masscan_fast_scan.sh <IP/CIDR>"
    exit 1
fi

echo "[*] Running Masscan on $TARGET"
sudo masscan -p1-65535 "$TARGET" --rate=$RATE -oJ "$OUT"

echo "[+] Results saved to $OUT"
