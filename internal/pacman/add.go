package pacman

import (
	"runtime"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/mistweaverco/tu/internal/config"
	"github.com/mistweaverco/tu/internal/shell"
)

func Add(flags config.ConfigFlags, packages []string) {
	if flags.DryRun {
		log.Info("Dry run", "Adding packages", packages, "flags", flags)
		return
	}
	if flags.Sync {
		sync(flags.Brew)
	}
	packagesString := strings.Join(packages, " ")
	if flags.Brew && runtime.GOOS != "windows" {
		shell.SetEnvironmentVariable("HOMEBREW_NO_AUTO_UPDATE", "1")
		shell.RunCommand(sh, "brew install "+packagesString)
		return
	}
	switch packageManager.Name {
	case "pacman":
		shell.RunCommand(sh, "sudo pacman -S "+packagesString)
	case "apt":
		shell.RunCommand(sh, "sudo apt install -y "+packagesString)
	case "dnf":
		shell.RunCommand(sh, "sudo dnf install -y "+packagesString)
	case "zypper":
		shell.RunCommand(sh, "sudo zypper install -y "+packagesString)
	case "yum":
		shell.RunCommand(sh, "sudo yum install -y "+packagesString)
	case "brew":
		shell.SetEnvironmentVariable("HOMEBREW_NO_AUTO_UPDATE", "1")
		shell.RunCommand(sh, "brew install "+packagesString)
	case "apk":
		shell.RunCommand(sh, "apk add "+packagesString)
	}
}
