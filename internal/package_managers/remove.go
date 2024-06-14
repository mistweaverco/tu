package package_managers

import (
	"github.com/charmbracelet/log"
	"github.com/mistweaverco/tu/internal/config"
	"github.com/mistweaverco/tu/internal/package_managers/apk"
	"github.com/mistweaverco/tu/internal/package_managers/apt"
	"github.com/mistweaverco/tu/internal/package_managers/brew"
	"github.com/mistweaverco/tu/internal/package_managers/dnf"
	"github.com/mistweaverco/tu/internal/package_managers/yay"
	"github.com/mistweaverco/tu/internal/package_managers/yum"
	"github.com/mistweaverco/tu/internal/package_managers/zypper"
	"github.com/mistweaverco/tu/internal/utils"
)

func Remove(flags config.ConfigFlags, packages []string) {
	if flags.DryRun {
		log.Info("Dry run", "Removing packages", packages, "flags", flags)
		return
	}
	switch utils.GetOS() {
	case "linux":
		if flags.Brew {
			brew.Remove(flags, packages)
			return
		}
		pacman := utils.GetLinuxPackageManager()
		switch pacman {
		case "apk":
			apk.Remove(flags, packages)
		case "apt":
			apt.Remove(flags, packages)
		case "dnf":
			dnf.Remove(flags, packages)
		case "yay":
			yay.Remove(flags, packages)
		case "yum":
			yum.Remove(flags, packages)
		case "zypper":
			zypper.Remove(flags, packages)
		}
	case "darwin":
		brew.Remove(flags, packages)
	case "windows":
	default:
		log.Error("Unsupported OS")
	}
}
