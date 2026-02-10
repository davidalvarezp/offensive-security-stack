module github.com/davidalvarezp/offensive-security-stack

// Go version policy
// -----------------
// - Go 1.21+ is recommended for:
//   * Improved crypto primitives
//   * Better net/http performance
//   * Mature generics support
// - Debian 12, Kali Linux, Ubuntu 22.04+ all support this
go 1.21

// Toolchain directive (optional but recommended for reproducibility)
// Comment this out if you want system Go instead
toolchain go1.21.6

require (
	// =========================
	// CLI & UX
	// =========================

	// Cobra is the de-facto standard for professional Go CLIs
	// Used for:
	// - Root command
	// - Subcommands
	// - Flags
	github.com/spf13/cobra v1.8.0

	// Viper is used for configuration management
	// Supports:
	// - YAML / JSON / TOML
	// - Environment variables
	// - Profiles
	github.com/spf13/viper v1.18.2

	// =========================
	// Terminal UI / Output
	// =========================

	// For colored output, tables, progress bars
	github.com/fatih/color v1.16.0
	github.com/olekukonko/tablewriter v0.0.5

	// Optional: nicer interactive menus (if you use them)
	github.com/manifoldco/promptui v0.9.0

	// =========================
	// Networking & Protocols
	// =========================

	// Packet capture & analysis (pcap reader only, no injection required)
	github.com/google/gopacket v1.1.19

	// DNS utilities
	github.com/miekg/dns v1.1.58

	// =========================
	// Security / Crypto
	// =========================

	// Extended cryptographic helpers (hashing, key derivation)
	golang.org/x/crypto v0.21.0

	// Secure random, system calls, OS-level utilities
	golang.org/x/sys v0.18.0

	// =========================
	// Parsing & Formats
	// =========================

	// YAML support (configs, labs, datasets)
	gopkg.in/yaml.v3 v3.0.1

	// =========================
	// Logging & Observability
	// =========================

	// Structured logging (audit-friendly)
	go.uber.org/zap v1.27.0

	// =========================
	// Reporting & Templates
	// =========================

	// HTML / Markdown reporting
	github.com/Masterminds/sprig/v3 v3.2.3

	// =========================
	// Validation & Utilities
	// =========================

	// Input validation for configs, labs, and modules
	github.com/go-playground/validator/v10 v10.18.0
)

require (
	// =========================
	// Indirect dependencies
	// (Managed automatically by Go)
	// =========================

	github.com/fsnotify/fsnotify v1.7.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/spf13/afero v1.11.0 // indirect
	github.com/spf13/cast v1.6.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/subosito/gotenv v1.6.0 // indirect
	golang.org/x/net v0.22.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)
