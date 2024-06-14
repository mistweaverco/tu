package apk

import (
	"github.com/charmbracelet/log"
	"github.com/mistweaverco/tu/internal/config"
	"github.com/mistweaverco/tu/internal/shell"
)

func Update(flags config.ConfigFlags) {
	if flags.DryRun {
		log.Info("Dry run", "Updating all packages", flags)
		return
	}
	sh := shell.GetCurrentShell()
	shell.RunCommand(sh, "apk upgrade")
}
