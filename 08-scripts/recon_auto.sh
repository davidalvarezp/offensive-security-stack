#!/bin/bash
# recon_auto.sh - Automated reconnaissance
TARGET=\$1
OUTPUT_DIR="../01-recon/\$TARGET"
mkdir -p \$OUTPUT_DIR

# Run Nmap scan
nmap -sS -sV -O -p- \$TARGET -oA \$OUTPUT_DIR/nmap_scan

# Run Amass scan
amass enum -d \$TARGET -o \$OUTPUT_DIR/amass_scan.txt

# Run DNS enumeration
dig \$TARGET ANY > \$OUTPUT_DIR/dns_enum.txt
whois \$TARGET > \$OUTPUT_DIR/whois.txt

# Run Masscan
masscan \$TARGET -p1-65535 --rate=1000 -oJ \$OUTPUT_DIR/masscan.json
