# Case Study: Insecure Web Application Design

## Classification
- Domain: Web Applications
- Focus: Secure Design Failures
- Context: Academic Demonstration Application
- Nature: Educational Case Study

---

## Scenario Overview

This case study examines a **fictional web application** designed for
instructional purposes. The application functions correctly from a user
perspective but contains **design weaknesses that increase security risk**.

The objective is to analyze **why the design failed**, not how to attack it.

---

## Application Characteristics

- User authentication via sessions
- Server-side rendering
- Database-backed user data
- Minimal security controls beyond basic functionality

---

## Observed Design Weaknesses (Conceptual)

The following issues are identified during review:

- Trust in client-provided input
- Inconsistent validation boundaries
- Security controls added late in development
- Lack of centralized security logic

These weaknesses are discussed **without exploit mechanics**.

---

## Failure Chain (Abstract)

The case demonstrates how insecurity emerges when:

1. Input validation is treated as optional
2. Business logic and security logic are intertwined
3. Trust assumptions are undocumented
4. Defensive testing is absent

---

## Root Causes

| Area | Issue |
|----|------|
| Architecture | No threat modeling |
| Development | Security added post-hoc |
| Process | Lack of secure coding standards |
| Testing | No adversarial review |

---

## Defensive Detection Opportunities

Defenders could identify issues through:

- Code review focused on trust boundaries
- Application logging and anomaly detection
- Security testing during development
- Monitoring unexpected application states

---

## Secure Design Lessons

- Separate data from logic
- Validate input at all boundaries
- Treat the client as untrusted
- Design security *before* implementation

---

## Teaching Objectives

Students should understand:

- Most web vulnerabilities originate in **design**, not syntax
- Prevention is more effective than detection
- Secure architecture reduces entire classes of risk

---

## Ethics Reminder

This case study exists to improve secure development practices.
It must not be used to target real systems.
