# OSS Lab Guide

## Purpose

This guide explains how to safely and correctly use the Offensive Security Stack (OSS)
within **authorized, isolated laboratory environments**.

OSS is designed for **learning, experimentation, and defensive understanding** —
not for real-world exploitation.

---

## Lab Environment Requirements

All OSS usage **MUST** occur in environments that meet the following criteria:

### Mandatory
- Isolated network (no internet exposure)
- Virtual machines or containers
- Explicit authorization from system owner
- No production data
- No external routing

### Recommended
- Hypervisor-based labs (VirtualBox, VMware, Proxmox)
- Snapshot support
- Logging enabled
- Instructor-controlled access

---

## Suggested Lab Architecture

```

[ Student VM ] ──┐
├── [ Isolated Virtual Switch ]
[ Target VM ]  ──┘

```

Optional components:
- SIEM VM
- Firewall VM
- Logging server

---

## Lab Usage Workflow

1. **Preparation**
   - Instructor prepares lab images
   - Students review ethics & scope
   - Lab mode enabled in OSS

2. **Execution**
   - Students run OSS modules
   - Observe simulated findings
   - Generate reports

3. **Analysis**
   - Review outputs
   - Map findings to defenses
   - Discuss mitigation strategies

4. **Reflection**
   - What signals would defenders see?
   - How could this be prevented?
   - What controls are missing?

---

## Prohibited Actions

❌ Running OSS on:
- Public networks
- Corporate networks
- Internet-facing systems
- Systems without consent

❌ Modifying OSS to:
- Add real exploits
- Remove safeguards
- Evade detection

---

## Instructor Checklist

- [ ] Authorization confirmed
- [ ] Isolation verified
- [ ] Ethics briefing completed
- [ ] Logs enabled
- [ ] Student access controlled

---

## Student Checklist

- [ ] I understand the ethical scope
- [ ] I am working in a lab
- [ ] I have permission
- [ ] I will not attempt real exploitation

---

## Violations

Improper use may result in:
- Academic penalties
- Loss of access
- Legal consequences

Use responsibly.
