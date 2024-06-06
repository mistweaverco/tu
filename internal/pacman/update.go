package pacman

import (
	"runtime"

	"github.com/charmbracelet/log"
	"github.com/mistweaverco/tu/internal/config"
	"github.com/mistweaverco/tu/internal/shell"
)

func Update(flags config.ConfigFlags) {
	if flags.DryRun {
		log.Info("Dry run", "Updating all packages", flags)
		return
	}
	if flags.Sync {
		sync(flags.Brew)
	}
	if flags.Brew && runtime.GOOS != "windows" {
		shell.SetEnvironmentVariable("HOMEBREW_NO_AUTO_UPDATE", "1")
		shell.RunCommand(sh, "brew upgrade")
		return
	}
	switch packageManager.Name {
	case "pacman":
		shell.RunCommand(sh, "sudo pacman -Syu")
	case "apt":
		shell.RunCommand(sh, "sudo apt upgrade -y")
	case "dnf":
		shell.RunCommand(sh, "sudo dnf upgrade -y")
	case "zypper":
		shell.RunCommand(sh, "sudo zypper update -y")
	case "yum":
		shell.RunCommand(sh, "sudo yum update -y")
	case "brew":
		shell.SetEnvironmentVariable("HOMEBREW_NO_AUTO_UPDATE", "1")
		shell.RunCommand(sh, "brew upgrade")
	case "apk":
		shell.RunCommand(sh, "apk upgrade")
	}
}
