package lab

import "fmt"

// PrintDisclaimer displays the OSS educational and ethical disclaimer.
func PrintDisclaimer() {
	fmt.Println("============================================================")
	fmt.Println("⚠️  Offensive Security Stack (OSS) - Educational Use Only ⚠️")
	fmt.Println("============================================================")
	fmt.Println()
	fmt.Println("This software is intended to be used strictly in controlled, isolated lab environments.")
	fmt.Println("Do NOT attempt to use OSS against systems you do not own or have explicit permission to test.")
	fmt.Println()
	fmt.Println("All modules, exercises, and demo exploits are for teaching purposes only.")
	fmt.Println("Unauthorized use of these techniques outside the lab is illegal and prohibited.")
	fmt.Println("By proceeding, you acknowledge that you understand and accept these rules.")
	fmt.Println()
}
