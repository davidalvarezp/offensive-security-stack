#!/bin/bash
# nmap_quick.sh - Quick Nmap scan
TARGET=\$1
OUTPUT_DIR="./\${TARGET}"
mkdir -p \$OUTPUT_DIR
nmap -sS -sV -O -p- \$TARGET -oA \$OUTPUT_DIR/nmap_\${TARGET}