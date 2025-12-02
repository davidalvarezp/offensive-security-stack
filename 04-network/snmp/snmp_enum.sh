#!/bin/bash
# snmp_enum.sh - SNMP enumeration
TARGET=\$1
OUTPUT_DIR="./\${TARGET}"
mkdir -p \$OUTPUT_DIR
for community in \$(cat ../snmp_default_communities.txt); do
    snmpwalk -v 2c -c \$community \$TARGET > \$OUTPUT_DIR/snmp_\${community}.txt
done
