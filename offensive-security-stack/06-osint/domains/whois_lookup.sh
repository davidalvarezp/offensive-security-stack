#!/bin/bash
# whois_lookup.sh - Perform WHOIS lookup
DOMAIN=\$1
OUTPUT_DIR="./\${DOMAIN}"
mkdir -p \$OUTPUT_DIR
whois \$DOMAIN > \$OUTPUT_DIR/whois_\${DOMAIN}.txt
