#!/bin/bash
# process_breach.sh - Process breach data
BREACH_FILE=\$1
OUTPUT_DIR="./processed/\$(basename \$BREACH_FILE .csv)"
mkdir -p \$OUTPUT_DIR

# Extract unique emails
cut -d',' -f1 \$BREACH_FILE | sort -u > \$OUTPUT_DIR/emails.txt

# Extract unique domains
cut -d',' -f2 \$BREACH_FILE | sort -u > \$OUTPUT_DIR/domains.txt

# Extract passwords
cut -d',' -f3 \$BREACH_FILE | sort -u > \$OUTPUT_DIR/passwords.txt
