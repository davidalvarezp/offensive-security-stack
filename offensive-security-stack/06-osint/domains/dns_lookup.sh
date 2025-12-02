#!/bin/bash
# dns_lookup.sh - Perform DNS lookup
DOMAIN=\$1
OUTPUT_DIR="./\${DOMAIN}"
mkdir -p \$OUTPUT_DIR
dig \$DOMAIN ANY > \$OUTPUT_DIR/dns_\${DOMAIN}.txt
whois \$DOMAIN > \$OUTPUT_DIR/whois_\${DOMAIN}.txt
