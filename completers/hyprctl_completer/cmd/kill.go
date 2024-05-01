package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var killCmd = &cobra.Command{
	Use:     "kill",
	Short:   "enter kill mode, where you can kill an app by clicking on it (ESCAPE to exit)",
	GroupID: "control",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(killCmd).Standalone()
	rootCmd.AddCommand(killCmd)
}
