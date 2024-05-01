package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var monitorsCmd = &cobra.Command{
	Use:     "monitors",
	Short:   "list all monitors with their properties",
	GroupID: "info",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(monitorsCmd).Standalone()
	rootCmd.AddCommand(monitorsCmd)
}
