#!/usr/bin/env python3
"""
Simple lab report generator
Collects logs from various folders and creates a summary report.
"""

import os
from datetime import datetime

# Directories to include in report
LOG_DIRS = [
    "./09-logs",
    "./04-network/logs",
    "./03-ad/logs",
    "./06-osint/logs",
    "./02-web/logs"
]

REPORT_FILE = f"10-reports/summary_report_{datetime.now().strftime('%Y%m%d_%H%M%S')}.txt"

os.makedirs("10-reports", exist_ok=True)

with open(REPORT_FILE, "w") as report:
    report.write(f"Lab Report generated on {datetime.now()}\n")
    report.write("="*50 + "\n\n")
    for log_dir in LOG_DIRS:
        report.write(f"Logs from {log_dir}:\n")
        if os.path.exists(log_dir):
            for file in os.listdir(log_dir):
                report.write(f" - {file}\n")
        else:
            report.write(" (No logs found)\n")
        report.write("\n")

print(f"[+] Report generated: {REPORT_FILE}")
