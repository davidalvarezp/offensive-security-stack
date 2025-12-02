#!/bin/bash
# Generates username variants (lab only)

NAME=$1
if [ -z "$NAME" ]; then
    echo "Usage: ./user_variations.sh <name>"
    exit 1
fi

FILE="$NAME-variants.txt"

echo "[*] Generating username variants..."

echo "$NAME"           >  "$FILE"
echo "${NAME}1"       >> "$FILE"
echo "${NAME}123"     >> "$FILE"
echo "${NAME}_lab"    >> "$FILE"
echo "${NAME}.user"   >> "$FILE"
echo "${NAME}2024"    >> "$FILE"

echo "[+] Saved to $FILE"
