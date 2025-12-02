#!/bin/bash
# user_variations.sh - Generate user variations
NAMES_FILE="names_lab.txt"
OUTPUT_FILE="user_variations.txt"

while read name; do
  echo "\$name"
  echo "\${name}1"
  echo "\${name}2"
  echo "\${name}!"
  echo "\${name}.\${name}"
done < \$NAMES_FILE > \$OUTPUT_FILE
