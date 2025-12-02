#!/bin/bash
# theharvester_scan.sh - TheHarvester scan
TARGET=\$1
OUTPUT_DIR="./\${TARGET}"
mkdir -p \$OUTPUT_DIR
theharvester -d \$TARGET -b google,bing,linkedin -l 500 > \$OUTPUT_DIR/theharvester_\${TARGET}.txt
