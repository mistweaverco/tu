package tu

import (
	"github.com/mistweaverco/tu/internal/config"
	"github.com/mistweaverco/tu/internal/package_managers"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "a",
	Short: "add packages",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, packages []string) {
		f := config.ConfigFlags{
			Brew:   cmd.Flag("brew").Changed,
			Sync:   cmd.Flag("sync").Changed,
			DryRun: cmd.Flag("dry-run").Changed,
		}
		package_managers.Add(f, packages)
	},
}

func init() {
	addCmd.Flags().BoolP("brew", "b", false, "use brew package manager")
	addCmd.Flags().BoolP("sync", "s", false, "sync with registry/mirrorlist")
	addCmd.Flags().Bool("dry-run", false, "dry run mode")
	addCmd.Aliases = []string{"add", "install", "i", "get"}
	rootCmd.AddCommand(addCmd)
}
