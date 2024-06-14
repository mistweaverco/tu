package brew

import (
	"github.com/charmbracelet/log"
	"github.com/mistweaverco/tu/internal/config"
	"github.com/mistweaverco/tu/internal/shell"
)

func Query(flags config.ConfigFlags, queryString string) {
	SetupHint()
	sh := shell.GetCurrentShell()
	if flags.DryRun {
		log.Info("Dry run", "Querying with", queryString)
		return
	}
	shell.SetEnvironmentVariable("HOMEBREW_NO_AUTO_UPDATE", "1")
	shell.RunCommand(sh, "brew search "+queryString)
}
