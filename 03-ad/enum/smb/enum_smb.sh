#!/bin/bash
# Enumerates SMB shares and server info
TARGET=$1

if [ -z "$TARGET" ]; then
    echo "Usage: ./enum_smb.sh <IP>"
    exit 1
fi

echo "[*] SMB Enumeration on $TARGET"

smbclient -L "$TARGET" -N
rpcclient -U "" "$TARGET" -c "enumdomusers"
enum4linux-ng -A "$TARGET"

echo "[+] SMB enumeration completed."
