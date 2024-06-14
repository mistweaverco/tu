package tu

import (
	"github.com/mistweaverco/tu/internal/config"
	"github.com/mistweaverco/tu/internal/package_managers"
	"github.com/spf13/cobra"
)

var removeCmd = &cobra.Command{
	Use:   "d",
	Short: "remove packages",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, packages []string) {
		f := config.ConfigFlags{
			DryRun: cmd.Flag("dry-run").Changed,
			Brew:   cmd.Flag("b").Changed,
		}
		package_managers.Remove(f, packages)
	},
}

func init() {
	removeCmd.Flags().Bool("dry-run", false, "dry run mode")
	removeCmd.Flags().BoolP("brew", "b", false, "use brew package manager")
	removeCmd.Aliases = []string{"remove", "delete", "del", "uninstall", "rm"}
	rootCmd.AddCommand(removeCmd)
}
