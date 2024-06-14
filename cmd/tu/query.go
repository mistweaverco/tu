package tu

import (
	"github.com/mistweaverco/tu/internal/config"
	"github.com/mistweaverco/tu/internal/package_managers"
	"github.com/spf13/cobra"
)

var queryCmd = &cobra.Command{
	Use:   "q",
	Short: "query/search packages",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		f := config.ConfigFlags{
			Brew:   cmd.Flag("brew").Changed,
			Sync:   cmd.Flag("sync").Changed,
			DryRun: cmd.Flag("dry-run").Changed,
		}
		package_managers.Query(f, args[0])
	},
}

func init() {
	queryCmd.Flags().BoolP("sync", "s", false, "sync with registry/mirrorlist")
	queryCmd.Flags().BoolP("brew", "b", false, "use brew package manager")
	queryCmd.Flags().Bool("dry-run", false, "dry run mode")
	queryCmd.Aliases = []string{"query", "search"}
	rootCmd.AddCommand(queryCmd)
}
