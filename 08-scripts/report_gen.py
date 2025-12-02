#!/usr/bin/env python3
# report_gen.py - Generate assessment report
import argparse
import datetime
import os

def generate_report(report_type, output_dir):
    timestamp = datetime.datetime.now().strftime("%Y-%m-%d_%H-%M-%S")
    report_path = os.path.join(output_dir, f"{report_type}_{timestamp}.md")
    
    with open(report_path, 'w') as f:
        f.write(f"# {report_type.upper()} Assessment Report\n")
        f.write(f"## Date: {datetime.date.today()}\n")
        f.write("## Executive Summary:\n")
        f.write("## Findings:\n")
        f.write("## Recommendations:\n")
        f.write("## Appendices:\n")
    
    print(f"Report generated at {report_path}")

if __name__ == "__main__":
    parser = argparse.ArgumentParser(description='Generate assessment report')
    parser.add_argument('report_type', choices=['recon', 'web', 'ad', 'network', 'c2'], help='Type of report to generate')
    parser.add_argument('--output-dir', default='./reports', help='Output directory for report')
    args = parser.parse_args()
    
    generate_report(args.report_type, args.output_dir)
