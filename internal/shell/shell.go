package shell

import (
	"os"
	"os/exec"
	"runtime"

	"github.com/charmbracelet/log"
)

func GetCurrentShell() string {
	shell := os.Getenv("SHELL")
	if shell == "" {
		log.Warn("SHELL environment variable not set")
		return "sh"
	}
	return shell
}

func SetEnvironmentVariable(key string, value string) {
	os.Setenv(key, value)
}

func RunCommand(shell string, command string) {
	c := "-c"
	if runtime.GOOS == "windows" {
		switch shell {
		case "powershell":
			c = "-Command"
		case "cmd":
			c = "/c"
		default:
			c = "-c"
		}
	}
	cmd := exec.Command(shell, c, command)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Error("Error running command", "error", err)
	}
}

func CommandExists(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}
