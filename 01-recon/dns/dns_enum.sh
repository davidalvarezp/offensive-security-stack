#!/bin/bash
# dns_enum.sh - DNS enumeration script
TARGET=\$1
OUTPUT_DIR="./\${TARGET}"
mkdir -p \$OUTPUT_DIR
dig \$TARGET ANY > \$OUTPUT_DIR/dig_\${TARGET}.txt
whois \$TARGET > \$OUTPUT_DIR/whois_\${TARGET}.txt