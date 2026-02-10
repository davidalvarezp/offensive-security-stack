# Offensive Security Stack (OSS)

**Offensive Security Stack (OSS)** is an open-source, Go-based **educational security framework** designed for teaching and learning **offensive security concepts** in **controlled, isolated laboratory environments**.

OSS is **not** a hacking toolset.  
It is a **modular learning platform** that demonstrates *how* offensive techniques work, *why* systems fail, and *how* defenders can detect and mitigate them.

---

## ⚠️ Important Notice (Read First)

> **OSS is intended for educational use only.**

This project **must only be used**:
- In **controlled laboratory environments**
- On **isolated networks (LAN / lab VMs)**
- On systems you **own** or have **explicit written permission** to test

Any use outside these conditions is **strictly prohibited** and may be illegal.

The authors and contributors **do not accept responsibility** for misuse.

See: [`docs/ETHICS.md`](docs/ETHICS.md)

---

## Project Goals

OSS was created to:

- Teach **offensive security fundamentals** without weaponized exploits
- Provide a **single, extensible master tool**
- Encourage understanding over automation
- Bridge **red team and blue team** perspectives
- Be safe to distribute as open source

This project is especially suited for:
- Universities & technical schools
- Cybersecurity courses
- Internal training labs
- CTF preparation (foundational skills)
- Defensive security education

---

## Architecture Overview

OSS is built around a **modular architecture**:

- A single **master CLI binary**
- Independent **modules** grouped by domain
- Centralized **guardrails** and lab enforcement
- Clear separation between:
  - Reconnaissance
  - Analysis
  - Simulation
  - Defensive mapping

```

oss
├── recon
├── network
├── web
├── auth
├── crypto
├── malware (simulation)
├── post-exploitation (analysis)
├── blue-team
└── offensive-tools (non-runnable demos)

````

Each module is designed to be:
- Self-contained
- Auditable
- Extendable
- Safe by default

---

## What OSS Includes

### ✅ Included
- Network discovery & analysis
- Service and configuration inspection
- Web security analysis (non-exploitative)
- Authentication & authorization audits
- Cryptography & entropy labs
- Malware lifecycle **simulations**
- Post-exploitation **analysis models**
- Blue team detection & mapping tools
- **Markdown-only demo exploit walkthroughs**

### ❌ Not Included
- Zero-day exploits
- Weaponized payloads
- Auto-exploitation frameworks
- Obfuscation or AV evasion engines
- Botnets or live C2 infrastructure

---

## Offensive Tools (Educational Demos)

The `modules/offensive_tools` section contains:

- 📄 **Markdown-only demonstrations**
- 🧠 Conceptual walkthroughs
- 📐 Pseudo-code (non-runnable)
- 🗺 MITRE ATT&CK mappings
- 🔍 Detection and mitigation notes

These demos exist to **explain**, not to execute.

---

## Safety & Guardrails

OSS enforces safety by design:

- Lab-mode enforcement at startup
- Environment validation (LAN-only)
- Explicit user confirmations
- No default dangerous behavior
- Centralized permission model
- Full audit logging support

If lab conditions are not met, OSS **will refuse to run**.

---

## Supported Platforms

OSS is designed for:

- Debian GNU/Linux (primary)
- Kali Linux
- Ubuntu LTS
- Parrot OS

### Requirements
- Go **1.21+**
- Linux (x86_64 / arm64)

---

## Installation

### From source

```bash
git clone https://github.com/davidalvarezp/offensive-security-stack.git
cd offensive-security-stack
go build -o oss ./cmd/oss
````

### Run (lab mode required)

```bash
export OSS_LAB_MODE=true
./oss
```

---

## Extending OSS (Modules)

OSS is built to be extended.

To add a new module:

1. Create a new folder under `modules/`
2. Implement the module interface
3. Register it in the module registry
4. Add documentation

See: [`docs/MODULE_DEV_GUIDE.md`](docs/MODULE_DEV_GUIDE.md)

---

## Documentation

* `docs/PHILOSOPHY.md` – Project mindset & scope
* `docs/ETHICS.md` – Usage rules & responsibilities
* `docs/THREAT_MODEL.md` – Assumptions & boundaries
* `docs/LAB_GUIDE.md` – Example lab setups
* `docs/MODULE_DEV_GUIDE.md` – Writing new modules

---

## Contributing

Contributions are welcome, especially:

* New educational modules
* Improved documentation
* Lab scenarios
* Blue team tooling
* Code quality improvements

Before contributing, please read:

* [`CONTRIBUTING.md`](.github/CONTRIBUTING.md)
* [`CODE_OF_CONDUCT.md`](CODE_OF_CONDUCT.md)

---

## License

This project is released under an open-source license.
See [`LICENSE`](LICENSE) for details.

---

## Final Words

> **OSS is about understanding systems, not breaking them blindly.**

If you teach security, OSS is built for you.

---

**Offensive Security Stack**
Educational • Modular • Open Source • Lab-Only
