#!/bin/bash
# nikto_scan.sh - Nikto scan
TARGET=\$1
OUTPUT_DIR="./\${TARGET}"
mkdir -p \$OUTPUT_DIR
nikto -h \$TARGET -o \$OUTPUT_DIR/nikto_\${TARGET}.txt
