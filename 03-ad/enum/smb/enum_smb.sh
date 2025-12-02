#!/bin/bash
# enum_smb.sh - SMB enumeration
DOMAIN=\$1
OUTPUT_DIR="../\${DOMAIN}"
mkdir -p \$OUTPUT_DIR
smbclient -L //\$DOMAIN/ -N > \$OUTPUT_DIR/smb_shares.txt
