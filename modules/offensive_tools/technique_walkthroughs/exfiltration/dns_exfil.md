# DNS Exfiltration (Educational Walkthrough)

## Purpose
Explain how DNS can be abused conceptually for data movement.

---

## Concept Overview
DNS exfiltration leverages:
- The ubiquity of DNS
- Small, frequent queries
- Encoded data in request fields

---

## Abstract Flow
1. Data is transformed
2. Data is embedded in DNS requests
3. Requests are resolved externally

---

## Defensive Indicators
- High-entropy domain labels
- Unusual query frequency
- Rare or algorithmic domains

---

## Defensive Controls
- DNS logging and analytics
- Egress filtering
- Known-good resolver enforcement

---

## Learning Outcome
Understand **why DNS monitoring is critical**.
