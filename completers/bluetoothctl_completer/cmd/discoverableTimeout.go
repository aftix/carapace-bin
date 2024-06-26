package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discoverableTimeoutCmd = &cobra.Command{
	Use:   "discoverable-timeout [value]",
	Short: "Set controller discoverable timeout",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discoverableTimeoutCmd).Standalone()
	rootCmd.AddCommand(discoverableTimeoutCmd)
}
