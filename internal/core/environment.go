package core

import (
	"fmt"
	"net"
	"os"

	"github.com/davidalvarezp/offensive-security-stack/internal/config"
)

// Environment holds runtime information about the lab/system
type Environment struct {
	IsLab       bool
	HostIP      string
	OS          string
	Architecture string
	Config      *config.Config
}

// NewEnvironment initializes environment details
func NewEnvironment(cfg *config.Config) *Environment {
	hostIP := detectHostIP()
	return &Environment{
		IsLab:        cfg.Lab.Enabled,
		HostIP:       hostIP,
		OS:           os.Getenv("OSTYPE"), // fallback for cross-platform
		Architecture: getArch(),
		Config:       cfg,
	}
}

// detectHostIP attempts to find a non-loopback IPv4 address
func detectHostIP() string {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "unknown"
	}
	for _, iface := range ifaces {
		addrs, err := iface.Addrs()
		if err != nil {
			continue
		}
		for _, a := range addrs {
			switch v := a.(type) {
			case *net.IPNet:
				ip := v.IP
				if !ip.IsLoopback() && ip.To4() != nil {
					return ip.String()
				}
			}
		}
	}
	return "unknown"
}

// getArch returns the runtime architecture
func getArch() string {
	return fmt.Sprintf("%s/%s", os.Getenv("GOOS"), os.Getenv("GOARCH"))
}
