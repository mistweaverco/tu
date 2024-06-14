package choco

import "github.com/mistweaverco/tu/internal/shell"

func Setup() {
	if !shell.CommandExists("choco") {
		shell.RunCommand("powershell", "Set-ExecutionPolicy AllSigned")
		shell.RunCommand("powershell", "Set-ExecutionPolicy Bypass -Scope Process -Force; [System.Net.ServicePointManager]::SecurityProtocol = [System.Net.ServicePointManager]::SecurityProtocol -bor 3072; iex ((New-Object System.Net.WebClient).DownloadString('https://chocolatey.org/install.ps1'))")
	}
}
