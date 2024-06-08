package pacman

import (
	"runtime"

	"github.com/charmbracelet/log"
	"github.com/mistweaverco/tu/internal/config"
	"github.com/mistweaverco/tu/internal/shell"
)

func Query(flags config.ConfigFlags, queryString string) {
	if flags.DryRun {
		log.Info("Dry run", "Querying with", queryString)
		return
	}
	if flags.Sync {
		sync(flags.Brew)
	}
	if flags.Brew && runtime.GOOS != "windows" {
		shell.SetEnvironmentVariable("HOMEBREW_NO_AUTO_UPDATE", "1")
		shell.RunCommand(sh, "brew search "+queryString)
		return
	}
	switch packageManager.Name {
	case "pacman":
		shell.RunCommand(sh, "sudo pacman -S "+queryString)
	case "apt":
		shell.RunCommand(sh, "sudo apt search "+queryString)
	case "dnf":
		shell.RunCommand(sh, "sudo dnf search "+queryString)
	case "zypper":
		shell.RunCommand(sh, "sudo zypper search "+queryString)
	case "yum":
		shell.RunCommand(sh, "sudo yum search "+queryString)
	case "brew":
		shell.SetEnvironmentVariable("HOMEBREW_NO_AUTO_UPDATE", "1")
		shell.RunCommand(sh, "brew search "+queryString)
	case "apk":
		shell.RunCommand(sh, "apk search "+queryString)
	}
}
