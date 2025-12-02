#!/bin/bash
# network_enum.sh - Network enumeration
TARGET=\$1
OUTPUT_DIR="../04-network/scanning/\$TARGET"
mkdir -p \$OUTPUT_DIR

# Run Nmap ping scan
nmap -sP \$TARGET -oG \$OUTPUT_DIR/nmap_ping_scan.txt

# Run Nmap port scan
nmap -sS -sU -p- \$TARGET -oA \$OUTPUT_DIR/nmap_port_scan

# Run Masscan
masscan \$TARGET -p1-65535 --rate=1000 -oJ \$OUTPUT_DIR/masscan_scan.json
