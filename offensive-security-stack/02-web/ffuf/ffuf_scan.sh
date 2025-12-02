#!/bin/bash
# ffuf_scan.sh - FFUF fuzzing scan
TARGET=\$1
WORDLIST=\$2
OUTPUT_DIR="./\${TARGET}"
mkdir -p \$OUTPUT_DIR
ffuf -w \$WORDLIST -u \$TARGET/FUZZ -mc all -fc 404 -o \$OUTPUT_DIR/ffuf_\${TARGET}.txt
