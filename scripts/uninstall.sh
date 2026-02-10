#!/bin/bash
set -e

BIN_PATH="/usr/local/bin/oss"
CONFIG_PATH="/etc/oss"
LOG_PATH="/var/log/oss"

echo "[OSS] Starting uninstallation..."

# Check root
if [ "$EUID" -ne 0 ]; then
    echo "[OSS] Please run as root"
    exit 1
fi

# Remove binary
if [ -f "${BIN_PATH}" ]; then
    rm -f "${BIN_PATH}"
    echo "[OSS] Removed binary"
fi

# Ask before removing config
read -p "[OSS] Remove configuration files in ${CONFIG_PATH}? [y/N]: " CONFIRM
if [ "$CONFIRM" = "y" ] || [ "$CONFIRM" = "Y" ]; then
    rm -rf "${CONFIG_PATH}"
    echo "[OSS] Configuration removed"
else
    echo "[OSS] Configuration preserved"
fi

# Ask before removing logs
read -p "[OSS] Remove log files in ${LOG_PATH}? [y/N]: " CONFIRM_LOG
if [ "$CONFIRM_LOG" = "y" ] || [ "$CONFIRM_LOG" = "Y" ]; then
    rm -rf "${LOG_PATH}"
    echo "[OSS] Logs removed"
else
    echo "[OSS] Logs preserved"
fi

echo "[OSS] Uninstallation completed."
