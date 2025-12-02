#!/bin/bash
# example_tool.sh - Custom tool template
TARGET=\$1
OUTPUT_DIR="../tools_output/\$TARGET"
mkdir -p \$OUTPUT_DIR

echo "[+] Running custom tool against \$TARGET"
# Add your custom functionality here
