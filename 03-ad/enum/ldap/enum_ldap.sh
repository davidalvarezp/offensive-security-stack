#!/bin/bash
# enum_ldap.sh - LDAP enumeration
DOMAIN=\$1
OUTPUT_DIR="../\${DOMAIN}"
mkdir -p \$OUTPUT_DIR
ldapsearch -x -H ldap://192.168.1.1 -D "cn=admin,cn=Users,dc=\$(echo \$DOMAIN | tr '.' ',dc=')" -w password -b "dc=\$(echo \$DOMAIN | tr '.' ',dc=')" > \$OUTPUT_DIR/ldap_dump.ldif
