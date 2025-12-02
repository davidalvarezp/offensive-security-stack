#!/bin/bash
# bh_collect.sh - BloodHound collection script
OUTPUT_DIR="./data"
mkdir -p \$OUTPUT_DIR
bloodhound-python -d lab.local -u user -p password -c All -ns 192.168.1.1 -o \$OUTPUT_DIR/bh_data.zip
