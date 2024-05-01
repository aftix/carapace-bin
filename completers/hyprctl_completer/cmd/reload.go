package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var reloadCmd = &cobra.Command{
	Use:     "reload",
	Short:   "force a reload of the config file",
	GroupID: "control",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(reloadCmd).Standalone()
	rootCmd.AddCommand(reloadCmd)
}
