#!/bin/bash
# generate_password_list.sh - Generate password list
OUTPUT_FILE="custom_passwords.txt"
touch \$OUTPUT_FILE

# Add common passwords
echo -e "password\n123456\nadmin\nletmein\nqwerty" >> \$OUTPUT_FILE

# Add dictionary words
cat /usr/share/dict/words | grep -E '^[a-zA-Z0-9]{8,12}$' >> \$OUTPUT_FILE

# Add common patterns
for i in {1..10}; do
  echo "password\$i"
  echo "admin\$i"
  echo "letmein\$i"
done >> \$OUTPUT_FILE
