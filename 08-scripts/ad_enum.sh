#!/bin/bash
# ad_enum.sh - Active Directory enumeration
DOMAIN=\$1
OUTPUT_DIR="../03-ad/enum/\$DOMAIN"
mkdir -p \$OUTPUT_DIR

# Run Kerberos enumeration
GetUserSPNs.py -request \$DOMAIN/ -dc-ip 192.168.1.1 > \$OUTPUT_DIR/kerberos_enum.txt

# Run LDAP enumeration
ldapsearch -x -H ldap://192.168.1.1 -D "cn=admin,cn=Users,dc=\$(echo \$DOMAIN | tr '.' ',dc=')" -w password -b "dc=\$(echo \$DOMAIN | tr '.' ',dc=')" > \$OUTPUT_DIR/ldap_dump.ldif

# Run RPC enumeration
rpcclient -U "" -N \$DOMAIN > \$OUTPUT_DIR/rpc_enum.txt

# Run SMB enumeration
smbclient -L //\$DOMAIN/ -N > \$OUTPUT_DIR/smb_enum.txt
