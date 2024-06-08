package tu

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "tu",
	Short: "tu is a tiny 🤏🏾wrapper for the your package 📦 manager.",
	Long:  "tu is a tiny 🤏🏾wrapper for the your package 📦 manager. It'll make managing 🖇️ your packages fun 😊 again!",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			err := cmd.Help()
			if err != nil {
				os.Exit(1)
			}
		}
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
