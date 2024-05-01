package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var keywordCmd = &cobra.Command{
	Use:     "keyword",
	Short:   "set a config keyword dynamically",
	GroupID: "control",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(keywordCmd).Standalone()
	rootCmd.AddCommand(keywordCmd)
}
