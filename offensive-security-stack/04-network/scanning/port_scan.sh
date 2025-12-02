#!/bin/bash
# port_scan.sh - Port scanning
TARGET=\$1
OUTPUT_DIR="./\${TARGET}"
mkdir -p \$OUTPUT_DIR
nmap -sS -sU -p- \$TARGET -oA \$OUTPUT_DIR/nmap_port_scan
masscan \$TARGET -p1-65535 --rate=1000 -oJ \$OUTPUT_DIR/masscan_port_scan.json
