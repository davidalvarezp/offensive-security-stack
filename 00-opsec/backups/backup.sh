#!/bin/bash
# Simple secure backup script for the Offensive VPS environment
# Stores compressed backups in this directory

DATE=$(date +"%Y-%m-%d_%H-%M-%S")
BACKUP_FILE="backup-$DATE.tar.gz"
TARGET_DIR="$HOME/offensive-security-stack"

echo "[*] Creating backup for: $TARGET_DIR"
tar -czf "$BACKUP_FILE" "$TARGET_DIR"

echo "[+] Backup created: $BACKUP_FILE"
