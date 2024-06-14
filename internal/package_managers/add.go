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

func Add(flags config.ConfigFlags, packages []string) {
	if flags.DryRun {
		log.Info("Dry run", "Adding packages", packages, "flags", flags)
		return
	}
	switch utils.GetOS() {
	case "linux":
		if flags.Brew {
			brew.Add(flags, packages)
			return
		}
		pacman := utils.GetLinuxPackageManager()
		switch pacman {
		case "apk":
			if flags.Sync {
				apk.Sync()
			}
			apk.Add(flags, packages)
		case "apt":
			if flags.Sync {
				apt.Sync()
			}
			apt.Add(flags, packages)
		case "dnf":
			if flags.Sync {
				dnf.Sync()
			}
			dnf.Add(flags, packages)
		case "yay":
			if flags.Sync {
				yay.Sync()
			}
			yay.Add(flags, packages)
		case "yum":
			if flags.Sync {
				yum.Sync()
			}
			yum.Add(flags, packages)
		case "zypper":
			if flags.Sync {
				zypper.Sync()
			}
			zypper.Add(flags, packages)
		}
	case "darwin":
		if flags.Sync {
			brew.Sync()
		}
		brew.Add(flags, packages)
	case "windows":
	default:
		log.Error("Unsupported OS")
	}
}
