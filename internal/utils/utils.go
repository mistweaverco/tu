package utils

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/charmbracelet/log"
)

type PackageManager struct {
	Name string
}

func OpenBrowser(url string) {
	var cmd string
	if IsMac() {
		cmd = "open"
	} else if IsWindows() {
		cmd = "start"
	} else if IsLinux() {
		cmd = "xdg-open"
	}
	c := exec.Command(cmd, url)
	c.Stdout = nil
	c.Stderr = nil
	err := c.Run()
	if err != nil {
		log.Error("Error running command", "error", err)
	}
}

// IsLinux returns true if the current operating system is Linux
func IsLinux() bool {
	return runtime.GOOS == "linux"
}

// IsWindows returns true if the current operating system is Windows
func IsWindows() bool {
	return runtime.GOOS == "windows"
}

// IsMac returns true if the current operating system is macOS
func IsMac() bool {
	return runtime.GOOS == "darwin"
}

// GetOS returns the current operating system
func GetOS() string {
	return runtime.GOOS
}

// GetLinuxPackageManager returns the package manager for the current Linux distribution
func GetLinuxPackageManager() string {
	if _, err := os.Stat("/etc/arch-release"); err == nil {
		// Arch Linux, we use yay - but need to make sure it is installed
		return "yay"
	} else if _, err := os.Stat("/etc/debian_version"); err == nil {
		// Debian-based distros, we use apt
		return "apt"
	} else if _, err := os.Stat("/etc/fedora-release"); err == nil {
		// Fedora, we use dnf
		return "dnf"
	} else if _, err := os.Stat("/etc/SuSE-release"); err == nil {
		// OpenSUSE, we use zypper
		return "zypper"
	} else if _, err := os.Stat("/etc/redhat-release"); err == nil {
		// Red Hat, we use yum
		return "yum"
	} else if _, err := os.Stat("/etc/alpine-release"); err == nil {
		// Alpine, we use apk
		return "apk"
	} else {
		// if we can't determine the package manager, return unknown
		return "unknown"
	}
}
