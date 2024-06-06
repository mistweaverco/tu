package pacman

import (
	"runtime"

	"github.com/mistweaverco/tu/internal/shell"
)

func sync(brew bool) {
	if brew && runtime.GOOS != "windows" {
		shell.RunCommand(sh, "brew update")
		return
	}
	switch packageManager.Name {
	case "pacman":
		shell.RunCommand(sh, "sudo pacman -Sy")
	case "apt":
		shell.RunCommand(sh, "sudo apt update")
	case "dnf":
		shell.RunCommand(sh, "sudo dnf check-update")
	case "zypper":
		shell.RunCommand(sh, "sudo zypper refresh")
	case "yum":
		shell.RunCommand(sh, "sudo yum check-update")
	case "brew":
		shell.RunCommand(sh, "brew update")
	case "apk":
		shell.RunCommand(sh, "apk update")
	}
}
