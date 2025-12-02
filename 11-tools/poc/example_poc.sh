#!/bin/bash
# example_poc.sh - Shell script PoC
TARGET=\$1

curl -s \$TARGET | grep -q "vulnerable_string"
if [ \$? -eq 0 ]; then
    echo "[+] Vulnerability confirmed!"
else
    echo "[-] Vulnerability not confirmed"
fi
