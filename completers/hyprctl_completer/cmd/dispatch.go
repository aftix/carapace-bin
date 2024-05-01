package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dispatchCmd = &cobra.Command{
	Use:     "dispatch",
	Short:   "call a dispatcher with an argument",
	GroupID: "control",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dispatchCmd).Standalone()
	rootCmd.AddCommand(dispatchCmd)
}
