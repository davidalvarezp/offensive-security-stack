# Threat Model – Offensive Security Stack (OSS)

## Overview

This document identifies threats **to and from OSS itself**
and defines safeguards to prevent misuse.

---

## Assets

- OSS source code
- Student environments
- Instructor trust
- Institutional reputation
- Educational integrity

---

## Intended Actors

### Authorized
- Students
- Instructors
- Researchers
- Contributors

### Unauthorized / Risk Actors
- Malicious users
- Script kiddies
- Unauthorized testers
- External attackers

---

## Threat Categories

### 1. Misuse Threats

| Threat | Impact | Mitigation |
|------|-------|------------|
| Running OSS on real networks | High | Ethics docs, lab checks |
| Adding exploit code | High | Code review, policy |
| Removing safeguards | High | Guardrails |

---

### 2. Technical Threats

| Threat | Impact | Mitigation |
|------|-------|------------|
| Privilege escalation | Medium | Permission checks |
| Sandbox escape | High | Lab enforcement |
| Unsafe defaults | Medium | Secure config |

---

### 3. Social / Educational Threats

| Threat | Impact | Mitigation |
|------|-------|------------|
| Students misunderstanding scope | Medium | Docs, disclaimers |
| Instructor misuse | Medium | Governance |
| Reputational harm | High | Clear messaging |

---

## Security Controls

- Lab mode enforcement
- Explicit disclaimers
- No real exploit code
- Reporting & logging
- Modular isolation
- Code reviews

---

## Out of Scope

OSS does NOT attempt to:
- Prevent all misuse
- Replace institutional controls
- Enforce legal compliance externally

Responsibility is shared.

---

## Residual Risk

Some risk remains due to:
- Open-source nature
- Educational context
- Human behavior

These risks are accepted in exchange for learning value,
but are actively minimized.

---

## Conclusion

OSS is designed to be:
- Safe
- Ethical
- Educational
- Defensible

Threat modeling is an ongoing process and will evolve
as the project grows.
