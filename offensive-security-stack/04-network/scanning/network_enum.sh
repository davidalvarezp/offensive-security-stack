#!/bin/bash
# network_enum.sh - Network enumeration
SUBNET=\$1
OUTPUT_DIR="./\${SUBNET//\//_}"
mkdir -p \$OUTPUT_DIR
nmap -sP \$SUBNET -oG \$OUTPUT_DIR/nmap_ping_scan.txt
