package yay

import (
	"github.com/mistweaverco/tu/internal/shell"
)

func Sync() {
	sh := shell.GetCurrentShell()
	shell.RunCommand(sh, "yay -Syy")
}
