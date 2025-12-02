#!/bin/bash
# nmap_full.sh - Full Nmap scan
TARGET=\$1
OUTPUT_DIR="./\${TARGET}"
mkdir -p \$OUTPUT_DIR
nmap -sS -sU -sV -O -p- -A \$TARGET -oA \$OUTPUT_DIR/nmap_full_\${TARGET}
