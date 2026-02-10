# Module Development Guide

## Purpose

This guide defines **how to design, implement, and review OSS modules**
while preserving safety, ethics, and educational value.

---

## Core Design Principles

Every module MUST:

- Be **non-destructive**
- Be **simulation-based**
- Teach a concept, not automate abuse
- Include defensive context
- Respect lab-only usage

---

## Module Structure

Each module should follow this pattern:

```

modules/<category>/<module_name>/
├── *.go
└── README.md (optional)

```

---

## Required Interfaces

Each module MUST implement:

- `Name() string`
- `Description() string`
- `Run(ctx *core.AppContext) error`

---

## Logging Requirements

Modules must:
- Log actions clearly
- Avoid sensitive data
- Indicate simulation status

Example:
```

"Simulated SQL injection pattern analysis completed"

```

---

## Reporting Requirements

Modules must:
- Use `ReportManager`
- Produce human-readable output
- Avoid raw payloads or exploits

---

## What Modules MUST NOT Do

❌ No real attacks  
❌ No exploit payloads  
❌ No shell execution  
❌ No privilege escalation  
❌ No obfuscation  
❌ No evasion logic  

---

## Acceptable Module Types

✅ Enumeration simulations  
✅ Misconfiguration analysis  
✅ Pattern detection  
✅ Defensive mapping  
✅ Risk scoring  
✅ Educational visualizations  

---

## Review Checklist

Before submission:

- [ ] Educational value clear
- [ ] No real exploitation
- [ ] Safe defaults
- [ ] Logs & reports present
- [ ] Ethics respected

---

## Rejected Contributions

Submissions will be rejected if they:
- Enable real-world attacks
- Bypass safeguards
- Misrepresent OSS purpose
- Encourage misuse

---

## Final Note

OSS modules are **teaching tools**.
If it would be unsafe to run on a real system, it does not belong here.