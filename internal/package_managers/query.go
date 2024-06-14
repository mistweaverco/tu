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

func Query(flags config.ConfigFlags, queryString string) {
	if flags.DryRun {
		log.Info("Dry run", "Query", queryString, "flags", flags)
		return
	}
	switch utils.GetOS() {
	case "linux":
		if flags.Brew {
			brew.Query(flags, queryString)
			return
		}
		pacman := utils.GetLinuxPackageManager()
		switch pacman {
		case "apk":
			apk.Query(flags, queryString)
		case "apt":
			apt.Query(flags, queryString)
		case "dnf":
			dnf.Query(flags, queryString)
		case "yay":
			yay.Query(flags, queryString)
		case "yum":
			yum.Query(flags, queryString)
		case "zypper":
			zypper.Query(flags, queryString)
		}
	case "darwin":
		brew.Query(flags, queryString)
	case "windows":
	default:
		log.Error("Unsupported OS")
	}
}
