package config

import (
	"os"

	"github.com/BurntSushi/toml"
	"github.com/charmbracelet/log"
)

type ConfigFlags struct {
	Brew   bool
	Sync   bool
	DryRun bool
}

type ConfigFilePackageManager struct {
	IsInstalled bool
}

type ConfileFilePackageManagers struct {
	Brew  ConfigFilePackageManager
	Choco ConfigFilePackageManager
}

type ConfigFile struct {
	PackageManagers ConfileFilePackageManagers
}

type Config struct {
	ConfigFile ConfigFile
	Flags      ConfigFlags
}

var PS = string(os.PathSeparator)

func getConfigDir() string {
	userConfigDir, err := os.UserConfigDir()
	if err != nil {
		log.Fatal(err)
	}
	return userConfigDir + PS + "tu"
}

func getConfigFilePath() string {
	return getConfigDir() + PS + "tu.toml"
}

func ensureConfigDir() {
	err := os.MkdirAll(getConfigDir(), 0755)
	if err != nil {
		log.Fatal(err)
	}
}

func ensureConfigFile() {
	ensureConfigDir()
	file, err := os.Create(getConfigFilePath())
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
}

func WriteConfig(config Config) {
	ensureConfigFile()
	configFilePath := getConfigFilePath()
	file, err := os.OpenFile(configFilePath, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	toml.NewEncoder(file).Encode(config)
}

func GetConfig() Config {
	ensureConfigFile()
	configFilePath := getConfigFilePath()
	config := Config{}
	toml.DecodeFile(configFilePath, &config)
	return config
}
