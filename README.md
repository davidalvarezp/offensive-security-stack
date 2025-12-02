
# Offensive Security Stack

A fully-organized, secure and lightweight offensive security environment
designed for **ethical hacking**, **penetration testing labs**,
**red team training**, and **cybersecurity research**.  

This repository documents how to prepare a minimal VPS (1GB RAM / 1 vCPU / 10GB SSD)
running **Debian 13** and turn it into a **professional, modular, well-structured offensive platform**.

---

## ⚠️ Legal & Ethical Disclaimer

This project, offensive-security-stack, is intended strictly for:
- Educational purposes
- Ethical hacking training
- Authorized penetration testing
- Controlled laboratory environments

Any misuse of the tools, techniques, or workflows documented in this repository
against systems without explicit permission is illegal.

The author assumes no responsibility for any damage caused by improper or
unauthorized use of this project.


---

## Goals of This Project

- Provide a **clean, organized filesystem structure** for offensive operations.
- Install a **minimal but powerful set of recon, web, AD, network and OSINT tools**.
- Harden the VPS and apply correct OPSEC principles.
- Prepare Docker for **isolated C2 frameworks** and AD visualization tools.
- Offer a standardized layout similar to what real Red Teams use.
- Support training for OSCP, PNPT, CRTO, and internal labs.

---

## System Min. Requirements

- **Debian 13**
- **1GB RAM** (ZRAM optimization included)
- **1 vCPU**
- **10GB SSD**
- SSH key-based access
- Basic Linux knowledge

---

## Server Preparation

All commands assume a **fresh Debian 13 VPS**.

### 1. Update System

```bash
sudo apt update && sudo apt upgrade -y
```

### 2. Create Operational User

```bash
adduser pentest
usermod -aG sudo pentest
```

Re-login:

```bash
su - pentest
```

### 3. Secure SSH

```bash
sudo nano /etc/ssh/sshd_config
```

Add/ensure:

```bash
PermitRootLogin no
PasswordAuthentication no
PubkeyAuthentication yes
```

Restart:

```bash
sudo systemctl restart ssh
```

### 4. Firewall & Intrusion Prevention

```bash
sudo apt install ufw fail2ban -y

sudo ufw default deny incoming
sudo ufw default allow outgoing
sudo ufw allow 22/tcp
sudo ufw enable
```

### 5. RAM Optimization for Small VPS

```bash
sudo apt install zram-tools -y
sudo systemctl enable zramswap --now
```

### 6. Base Tools

```bash
sudo apt install -y \
build-essential git curl wget jq tmux htop lsof \
python3 python3-pip python3-venv \
golang ruby nodejs npm \
netcat-traditional tcpdump
```

### 7. Recon Tools

```bash
sudo apt install -y \
nmap masscan whatweb dnsutils whois
```

### 8. Web Pentest Tools

```bash
sudo apt install -y \
ffuf gobuster nikto sqlmap
```

### 9. Active Directory (Lab)

```bash
sudo apt install -y \
smbclient ldap-utils rpcclient krb5-user enum4linux-ng
```

### 10. Internal Network Tools

```bash
sudo apt install -y \
onesixtyone snmp snmp-mibs-downloader
```

### 11. OSINT Tools

```bash
sudo apt install -y \
theharvester amass
```

### 12. Wordlists

```bash
sudo apt install -y wordlists seclists
```

### 13. Install Docker

```bash
sudo apt install -y docker.io docker-compose
sudo systemctl enable docker --now
sudo usermod -aG docker pentest
```


Reconnect afterwards.

---

## Folder Structure

Generate the complete professional structure:

```bash
git pull
```


### Link Wordlists

```bash
ln -s /usr/share/wordlists ~/VPS-OFFENSIVE/07-wordlists/system
ln -s /usr/share/seclists ~/VPS-OFFENSIVE/07-wordlists/web
```


### Initialize Scripts & Logs

```bash
touch ~/VPS-OFFENSIVE/08-scripts/{recon_auto.sh,web_enum.sh,ad_enum.sh,network_enum.sh}
touch ~/VPS-OFFENSIVE/08-scripts/report_gen.py
chmod +x ~/VPS-OFFENSIVE/08-scripts/*.sh

touch ~/VPS-OFFENSIVE/09-logs/{recon.log,web.log,ad.log,network.log,c2.log,system.log}
```

---

## Architecture Overview

```bash
offensive-security-stack/
├── 00-opsec/        → Security, backups, firewall, VPN
├── 01-recon/        → Enumeration tools
├── 02-web/          → Web pentesting
├── 03-ad/           → AD lab tools
├── 04-network/      → Internal network ops
├── 05-c2/           → Dockerized C2 (lab use only)
├── 06-osint/        → OSINT operations
├── 07-wordlists/    → Dictionaries & wordlists
├── 08-scripts/      → Automation scripts
├── 09-logs/         → Central logging
├── 10-reports/      → Structured reporting
├── 11-tools/        → External tools & PoCs
└── 99-archive/      → Archived projects
```

---

## Ethical Notice

This repository:

- Does not include malware, payloads, exploits, C2 binaries or attack scripts.
- Is only intended for legal, ethical and controlled environments.
- Must never be used on systems you do not own or lack explicit written authorization to test.

**Unauthorized access** is **illegal** and punishable by **law**.

---

## Contributing

Contributions for improvements, documentation, structure, and automation are welcome.
Content promoting illegal activity will not be accepted.

---

## License

MIT License — Free to use for education and research.

---

## Support

If this project helps you organize your offensive environment, a star is appreciated!
