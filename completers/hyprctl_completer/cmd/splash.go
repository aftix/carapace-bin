package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var splashCmd = &cobra.Command{
	Use:     "splash",
	Short:   "prints the current random splash",
	GroupID: "info",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(splashCmd).Standalone()
	rootCmd.AddCommand(splashCmd)
}
