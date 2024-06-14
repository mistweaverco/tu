package brew

import (
	"github.com/mistweaverco/tu/internal/shell"
)

func Sync() {
	sh := shell.GetCurrentShell()
	shell.SetEnvironmentVariable("HOMEBREW_NO_AUTO_UPDATE", "1")
	shell.RunCommand(sh, "brew update")
}
