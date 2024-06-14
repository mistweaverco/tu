package choco

import (
	"strings"

	"github.com/charmbracelet/log"
	"github.com/mistweaverco/tu/internal/config"
	"github.com/mistweaverco/tu/internal/shell"
)

func Add(flags config.ConfigFlags, packages []string) {
	sh := shell.GetCurrentShell()
	packagesString := strings.Join(packages, " ")
	if flags.DryRun {
		log.Info("Dry run", "Adding packages", packages, "flags", flags)
		return
	}
	shell.RunCommand(sh, "choco install "+packagesString)
}
