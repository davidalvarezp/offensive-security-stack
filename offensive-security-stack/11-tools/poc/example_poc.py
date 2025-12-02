#!/usr/bin/env python3
# example_poc.py - Proof of Concept
import requests

def poc(url):
    try:
        response = requests.get(url)
        if response.status_code == 200:
            print("[+] Vulnerability confirmed!")
        else:
            print("[-] Vulnerability not confirmed")
    except Exception as e:
        print(f"[-] Error: {e}")

if __name__ == "__main__":
    url = input("Enter target URL: ")
    poc(url)
