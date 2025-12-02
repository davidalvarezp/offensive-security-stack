#!/bin/bash
echo "[*] Applying UFW rules..."

sudo ufw disable
sudo ufw --force reset

# Load rules
while read rule; do
    [[ $rule =~ ^#.*$ || -z $rule ]] && continue
    echo "Applying: $rule"
    sudo ufw $rule
done < ufw-rules.txt

sudo ufw enable

echo "[+] Firewall rules applied successfully."
