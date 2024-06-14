package yay

import (
	"github.com/charmbracelet/log"
	"github.com/mistweaverco/tu/internal/config"
	"github.com/mistweaverco/tu/internal/shell"
)

func Query(flags config.ConfigFlags, queryString string) {
	sh := shell.GetCurrentShell()
	if flags.DryRun {
		log.Info("Dry run", "Querying with", queryString)
		return
	}
	shell.RunCommand(sh, "yay -Ss "+queryString)
}
