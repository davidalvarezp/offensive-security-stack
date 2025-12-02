#!/bin/bash
echo "[*] Applying secure SSH configuration..."

sudo cp secure_sshd_config /etc/ssh/sshd_config
sudo systemctl restart ssh

echo "[+] SSH configuration hardened."
