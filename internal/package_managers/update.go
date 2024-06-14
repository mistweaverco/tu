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

func Update(flags config.ConfigFlags) {
	if flags.DryRun {
		log.Info("Dry run: Update", "flags", flags)
		return
	}
	switch utils.GetOS() {
	case "linux":
		if flags.Brew {
			brew.Update(flags)
			return
		}
		pacman := utils.GetLinuxPackageManager()
		switch pacman {
		case "apk":
			if flags.Sync {
				apk.Sync()
			}
			apk.Update(flags)
		case "apt":
			if flags.Sync {
				apt.Sync()
			}
			apt.Update(flags)
		case "dnf":
			if flags.Sync {
				dnf.Sync()
			}
			dnf.Update(flags)
		case "yay":
			if flags.Sync {
				yay.Sync()
			}
			yay.Update(flags)
		case "yum":
			if flags.Sync {
				yum.Sync()
			}
			yum.Update(flags)
		case "zypper":
			if flags.Sync {
				zypper.Sync()
			}
			zypper.Update(flags)
		}
	case "darwin":
		if flags.Sync {
			brew.Sync()
		}
		brew.Update(flags)
	case "windows":
	default:
		log.Error("Unsupported OS")
	}
}
