package utils

import "os"

type PackageManager struct {
	Name string
}

func findAvailablePackageManager(osName string) string {
	switch osName {
	case "macOS":
		return "brew"
	case "Linux":
		if _, err := os.Stat("/etc/arch-release"); err == nil {
			return "pacman"
		} else if _, err := os.Stat("/etc/debian_version"); err == nil {
			return "apt"
		} else if _, err := os.Stat("/etc/fedora-release"); err == nil {
			return "dnf"
		} else if _, err := os.Stat("/etc/SuSE-release"); err == nil {
			return "zypper"
		} else if _, err := os.Stat("/etc/redhat-release"); err == nil {
			return "yum"
		} else if _, err := os.Stat("/etc/alpine-release"); err == nil {
			return "apk"
		} else {
			return "unknown"
		}
	case "Windows":
		return "choco"
	default:
		return "unknown"
	}
}

func GetPackageManager(goos string) PackageManager {
	os := GetOperatingSystem(goos)
	p := PackageManager{
		Name: findAvailablePackageManager(os),
	}
	return p
}

func GetOperatingSystem(goos string) string {
	switch goos {
	case "darwin":
		return "macOS"
	case "linux":
		return "Linux"
	case "windows":
		return "Windows"
	default:
		return "Other"
	}
}
