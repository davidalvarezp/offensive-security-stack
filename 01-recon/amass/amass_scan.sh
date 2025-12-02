#!/bin/bash
# amass_scan.sh - Run Amass scan
TARGET=\$1
OUTPUT_DIR="./\${TARGET}"
mkdir -p \$OUTPUT_DIR
amass enum -d \$TARGET -o \$OUTPUT_DIR/amass_enum.txt
amass viz -d \$TARGET -o \$OUTPUT_DIR/amass_viz.png