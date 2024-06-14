package apt

import (
	"github.com/mistweaverco/tu/internal/shell"
)

func Sync() {
	sh := shell.GetCurrentShell()
	shell.RunCommand(sh, "sudo apt update")
}
