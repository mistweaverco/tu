package brew

import (
	"github.com/charmbracelet/huh"
	"github.com/mistweaverco/tu/internal/shell"
	"github.com/mistweaverco/tu/internal/utils"
)

func SetupHint() {
	if !shell.CommandExists("brew") {
		confirm := false
		huh.NewConfirm().
			Title("Homebrew is not installed. Do you want to install it?").
			Affirmative("Yes").
			Negative("No").
			Value(&confirm).
			Run()
		if confirm {
			utils.OpenBrowser("https://brew.sh")
		}
	}
}
