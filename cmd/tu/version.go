package tu

import (
	"github.com/charmbracelet/log"
	"github.com/mistweaverco/tu/internal/utils"
	"github.com/spf13/cobra"
)

var VERSION string

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of tu",
	Run: func(cmd *cobra.Command, args []string) {
		os := utils.GetOS()
		log.Info("Version", os, VERSION)
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}
