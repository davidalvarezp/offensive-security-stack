# TCP Reverse Shell Demonstration (Educational)

## Purpose

Introduce the **concept of direct TCP-based control channels** without technical detail.

---

## Concept Summary

A TCP reverse shell:
- Initiates a connection outward
- Maintains a persistent session
- Enables bidirectional communication

---

## Conceptual Risks

- Long-lived connections
- Rare destination ports
- Unusual traffic persistence

---

## Defensive Controls

- Egress filtering
- Connection duration analysis
- Application-aware firewalls

---

## Educational Outcome

Understand **why outbound control is often more dangerous than inbound attacks**.
