package yum

import (
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
	sh := shell.GetCurrentShell()
	packagesString := strings.Join(packages, " ")
	shell.RunCommand(sh, "sudo yum remove -y "+packagesString)
}
