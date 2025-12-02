
## System **Minimun** Requirements

- **Debian 13**
- **1GB RAM** (ZRAM optimization included)
- **1 vCPU**
- **10GB SSD**
- SSH key-based access
- Basic Linux knowledge

---

## System Preparation

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
