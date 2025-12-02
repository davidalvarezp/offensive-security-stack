#!/bin/bash
# Network routing visualization and basic checks

echo "[*] Display current routes"
ip route show

echo "[*] Ping default gateway"
GATEWAY=$(ip route | grep default | awk '{print $3}')
ping -c 3 $GATEWAY

echo "[*] Trace route to lab target (example: 10.10.10.10)"
# traceroute 10.10.10.10
