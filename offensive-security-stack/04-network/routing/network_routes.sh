#!/bin/bash
# network_routes.sh - Configure network routes
TARGET_SUBNET=\$1
GATEWAY_IP=\$2
INTERFACE=\$3

# Add route to target subnet
ip route add \$TARGET_SUBNET via \$GATEWAY_IP dev \$INTERFACE

# Enable IP forwarding
echo 1 > /proc/sys/net/ipv4/ip_forward
