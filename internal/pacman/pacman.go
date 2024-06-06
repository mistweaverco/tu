package pacman

import (
	"runtime"

	"github.com/mistweaverco/tu/internal/shell"
	"github.com/mistweaverco/tu/internal/utils"
)

var packageManager = utils.GetPackageManager(runtime.GOOS)
var sh = shell.GetCurrentShell()
