package zypper

import (
	"github.com/mistweaverco/tu/internal/shell"
)

func Sync() {
	sh := shell.GetCurrentShell()
	shell.RunCommand(sh, "sudo zypper refresh")
}
