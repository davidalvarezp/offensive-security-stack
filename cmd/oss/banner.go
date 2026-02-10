package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/fatih/color"
)

// Banner metadata.
// These values should be kept in sync with version.go.
const (
	projectName = "Offensive Security Stack"
	projectAcr  = "OSS"
	projectDesc = "Offensive Security Stack Framework (Lab-Only)"
)

// PrintBanner displays the project banner and legal notice.
// It should be called once at application startup.
func PrintBanner(version string) {
	now := time.Now().Format("2006-01-02 15:04:05")

	header := color.New(color.FgRed, color.Bold)
	sub := color.New(color.FgWhite)
	meta := color.New(color.FgHiBlack)
	warn := color.New(color.FgYellow, color.Bold)

	header.Println(`
   ____  _____ _____
  / __ \/ ___// ___/
 / / / /\__ \ \__ \
/ /_/ /___/ /___/ /
\____//____//____/

`)

	sub.Printf("%s (%s)\n", projectName, projectAcr)
	sub.Println(projectDesc)
	fmt.Println()

	meta.Printf("Version      : %s\n", version)
	meta.Printf("Go Runtime   : %s\n", runtime.Version())
	meta.Printf("OS / Arch    : %s / %s\n", runtime.GOOS, runtime.GOARCH)
	meta.Printf("Started At  : %s\n", now)
	fmt.Println()

	warn.Println("⚠️  WARNING: EDUCATIONAL USE ONLY")
	warn.Println("This software is intended for:")
	fmt.Println(" - Controlled laboratory environments")
	fmt.Println(" - Isolated networks (LAN only)")
	fmt.Println(" - Academic training and research")
	fmt.Println()
	fmt.Println("Unauthorized use against systems you do not own or")
	fmt.Println("explicitly have permission to test is strictly prohibited.")
	fmt.Println()
}
