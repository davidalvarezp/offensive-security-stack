#!/bin/bash
set -e

BIN_NAME="oss"
INSTALL_PATH="/usr/local/bin"
CONFIG_PATH="/etc/oss"
LOG_PATH="/var/log/oss"

echo "[OSS] Starting installation..."

# Check root
if [ "$EUID" -ne 0 ]; then
    echo "[OSS] Please run as root"
    exit 1
fi

# Install binary
if [ ! -f "./${BIN_NAME}" ]; then
    echo "[OSS] Binary '${BIN_NAME}' not found in current directory"
    exit 1
fi

cp "./${BIN_NAME}" "${INSTALL_PATH}/${BIN_NAME}"
chmod 755 "${INSTALL_PATH}/${BIN_NAME}"

# Create directories
mkdir -p "${CONFIG_PATH}"
mkdir -p "${LOG_PATH}"

chmod 755 "${CONFIG_PATH}"
chmod 755 "${LOG_PATH}"

echo "[OSS] Installed binary to ${INSTALL_PATH}/${BIN_NAME}"
echo "[OSS] Configuration directory: ${CONFIG_PATH}"
echo "[OSS] Log directory: ${LOG_PATH}"

echo "[OSS] Installation completed successfully."
