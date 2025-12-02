#!/bin/bash
# web_enum.sh - Web application enumeration
TARGET=\$1
OUTPUT_DIR="../02-web/logs/\$TARGET"
mkdir -p \$OUTPUT_DIR

# Run Nikto scan
nikto -h \$TARGET -o \$OUTPUT_DIR/nikto_scan.txt

# Run Gobuster scan
gobuster dir -u \$TARGET -w ../../07-wordlists/web/common_paths.txt -o \$OUTPUT_DIR/gobuster_scan.txt

# Run SQLMap scan
sqlmap -u \$TARGET --batch --dump-all -o \$OUTPUT_DIR/sqlmap_scan.txt
