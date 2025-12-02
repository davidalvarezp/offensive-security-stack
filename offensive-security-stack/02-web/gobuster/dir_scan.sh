#!/bin/bash
# dir_scan.sh - Gobuster directory scan
TARGET=\$1
WORDLIST=\$2
OUTPUT_DIR="./\${TARGET}"
mkdir -p \$OUTPUT_DIR
gobuster dir -u \$TARGET -w \$WORDLIST -o \$OUTPUT_DIR/gobuster_\${TARGET}.txt
