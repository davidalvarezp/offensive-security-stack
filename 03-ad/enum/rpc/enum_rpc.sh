#!/bin/bash
# enum_rpc.sh - RPC enumeration
DOMAIN=\$1
OUTPUT_DIR="../\${DOMAIN}"
mkdir -p \$OUTPUT_DIR
rpcclient -U "" -N \$DOMAIN > \$OUTPUT_DIR/rpc_services.txt
E