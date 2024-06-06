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
			cmd.Help()
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
	// rootCmd.PersistentFlags().StringVar(&cfg.ConfigPath, "config", "shazam.yml", "config file")
	// rootCmd.PersistentFlags().BoolVar(&cfg.Flags.DryRun, "dry-run", false, "dry run")
	// rootCmd.PersistentFlags().BoolVar(&cfg.Flags.PullInExisting, "pull-in-existing", false, "pull in existing files")
	// rootCmd.PersistentFlags().StringVar(&cfg.Flags.Root, "root", "", "root workspace")
	// rootCmd.PersistentFlags().StringVar(&cfg.Flags.Only, "only", "", "only specific nodes matching a name")
	// rootCmd.PersistentFlags().StringVar(&cfg.Flags.DotfilesPath, "dotfiles-path", "", "dotfiles path")
	// rootCmd.PersistentFlags().StringVar(&cfg.Flags.Path, "path", "", "path to config file or dir")
}
