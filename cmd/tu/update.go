package tu

import (
	"github.com/mistweaverco/tu/internal/config"
	"github.com/mistweaverco/tu/internal/pacman"
	"github.com/spf13/cobra"
)

var updateCmd = &cobra.Command{
	Use:   "u",
	Short: "update all packages",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, packages []string) {
		f := config.ConfigFlags{
			Sync:   cmd.Flag("sync").Changed,
			DryRun: cmd.Flag("dry-run").Changed,
			Brew:   cmd.Flag("b").Changed,
		}
		pacman.Update(f)
	},
}

func init() {
	updateCmd.Flags().Bool("dry-run", false, "dry run mode")
	updateCmd.Flags().BoolP("sync", "s", false, "sync registry/mirrors before updating all packages")
	updateCmd.Flags().BoolP("brew", "b", false, "use brews package manager")
	updateCmd.Aliases = []string{"update", "upgrade", "up"}
	rootCmd.AddCommand(updateCmd)
}
