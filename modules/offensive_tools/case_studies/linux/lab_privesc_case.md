# Case Study: Linux Privilege Escalation in a Lab Environment

## Classification
- Domain: Operating Systems (Linux)
- Focus: Privilege Escalation (Conceptual)
- Context: Isolated Academic Lab
- Nature: Educational Case Study

---

## Scenario Overview

In this case study, a student analyst is presented with a **Linux lab system**
where a non-privileged user account has access to the system for routine tasks.

The goal is **not** to exploit the system, but to **analyze how configuration
decisions can unintentionally increase privilege escalation risk**.

---

## Initial Conditions

- Multi-user Linux system
- Standard user permissions
- Several administrative conveniences enabled for usability
- No malicious activity initially present

---

## Observed Conditions (High-Level)

During analysis, the following conditions are identified:

- Multiple system utilities available to non-admin users
- Legacy configurations preserved for compatibility
- Administrative delegation configured for convenience
- Inconsistent permission hygiene across system components

These observations **do not imply exploitation**, but highlight **risk factors**.

---

## Conceptual Escalation Path (Abstract)

This case study explores how escalation *could* occur if:

1. A trusted component operates with elevated privileges
2. A lower-privileged user can influence that component’s behavior
3. Boundaries between trust levels are not strictly enforced

No technical steps or mechanisms are described.

---

## Root Causes

| Category | Description |
|--------|-------------|
| Configuration | Overly permissive defaults |
| Governance | Lack of periodic privilege review |
| Design | Convenience prioritized over security |
| Visibility | Insufficient auditing |

---

## Defensive Detection Opportunities

Blue teams could detect risk by monitoring:

- Unexpected privilege boundary crossings
- Administrative utilities used outside expected roles
- Permission changes over time
- Inconsistent ownership or access rights

---

## Preventive Controls

- Principle of least privilege
- Regular permission audits
- Clear separation of administrative duties
- Documentation of trust assumptions

---

## Teaching Objectives

Students should learn:

- Privilege escalation is often a **design issue**, not a bug
- Small configuration decisions compound over time
- Defensive thinking begins with understanding trust boundaries

---

## Ethics Reminder

This case study is **analytical only**.
No attempt should be made to reproduce escalation behavior outside authorized labs.
