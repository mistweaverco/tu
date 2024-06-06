package pacman

import "github.com/mistweaverco/tu/internal/shell"

func InstallHomebrew() {
	if !shell.CommandExists("brew") {
		shell.RunCommand("/bin/bash", "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)")
	}
}

func InstallChocolatey() {
	if !shell.CommandExists("choco") {
		shell.RunCommand("powershell", "Set-ExecutionPolicy AllSigned")
		shell.RunCommand("powershell", "Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))")
	}
}
