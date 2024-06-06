package pacman

import (
	"runtime"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/mistweaverco/tu/internal/config"
	"github.com/mistweaverco/tu/internal/shell"
)

func Remove(flags config.ConfigFlags, packages []string) {
	if flags.DryRun {
		log.Info("Dry run", "Removing packages", packages, "flags", flags)
		return
	}
	packagesString := strings.Join(packages, " ")
	if flags.Brew && runtime.GOOS != "windows" {
		shell.SetEnvironmentVariable("HOMEBREW_NO_AUTO_UPDATE", "1")
		shell.RunCommand(sh, "brew remove "+packagesString)
		return
	}
	switch packageManager.Name {
	case "pacman":
		shell.RunCommand(sh, "sudo pacman -Runs "+packagesString)
	case "apt":
		shell.RunCommand(sh, "sudo apt remove -y "+packagesString)
	case "dnf":
		shell.RunCommand(sh, "sudo dnf remove -y "+packagesString)
	case "zypper":
		shell.RunCommand(sh, "sudo zypper remove -y "+packagesString)
	case "yum":
		shell.RunCommand(sh, "sudo yum remove -y "+packagesString)
	case "brew":
		shell.SetEnvironmentVariable("HOMEBREW_NO_AUTO_UPDATE", "1")
		shell.RunCommand(sh, "brew remove "+packagesString)
	case "apk":
		shell.RunCommand(sh, "apk del "+packagesString)
	}
}
