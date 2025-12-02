
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
## System Requirements & Preparation

Go to [README_system-preparation.md](README_system-preparation.md)

---

## Goals of This Project

- Provide a **clean, organized filesystem structure** for offensive operations.
- Install a **minimal but powerful set of recon, web, AD, network and OSINT tools**.
- Harden the VPS and apply correct OPSEC principles.
- Prepare Docker for **isolated C2 frameworks** and AD visualization tools.
- Offer a standardized layout similar to what real Red Teams use.
- Support training for OSCP, PNPT, CRTO, and internal labs.

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

