# HTTP Reverse Shell Demonstration (Educational)

## Purpose

This document describes the **conceptual flow** of HTTP-based command channels
to help students understand why **normal-looking traffic can be risky**.

No shell code or commands are provided.

---

## Concept Overview

An HTTP-based reverse shell:
- Uses outbound web requests
- Blends with normal traffic
- Relies on application-layer protocols

---

## Abstract Communication Flow

1. Endpoint initiates outbound request
2. Server responds with instructions
3. Endpoint executes logic
4. Response is sent back

---

## Why HTTP Is Chosen

- Often allowed through firewalls
- Familiar protocol
- Proxy-friendly

---

## Defensive Indicators

| Indicator | Meaning |
|--------|---------|
| Abnormal request frequency | Automation |
| Unusual headers | Non-browser behavior |
| Uniform payload sizes | Encoded instructions |

---

## Teaching Focus

Understanding **why protocol allowlists are insufficient alone**.
