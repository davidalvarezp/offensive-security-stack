#!/bin/bash
# sqlmap_scan.sh - SQLMap scan
TARGET=\$1
OUTPUT_DIR="./\${TARGET}"
mkdir -p \$OUTPUT_DIR
sqlmap -u \$TARGET --batch --dump-all -o \$OUTPUT_DIR/sqlmap_\${TARGET}.txt
