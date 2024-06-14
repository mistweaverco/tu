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

func Sync(flags config.ConfigFlags) {
	if flags.DryRun {
		log.Info("Dry run: Sync", "flags", flags)
		return
	}
	switch utils.GetOS() {
	case "linux":
		if flags.Brew {
			brew.Sync()
			return
		}
		pacman := utils.GetLinuxPackageManager()
		switch pacman {
		case "apk":
			apk.Sync()
		case "apt":
			apt.Sync()
		case "dnf":
			dnf.Sync()
		case "yay":
			yay.Sync()
		case "yum":
			yum.Sync()
		case "zypper":
			zypper.Sync()
		}
	case "darwin":
		brew.Sync()
	case "windows":
	default:
		log.Error("Unsupported OS")
	}
}
