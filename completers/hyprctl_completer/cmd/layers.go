package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var layersCmd = &cobra.Command{
	Use:     "layers",
	Short:   "list all the layers",
	GroupID: "info",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(layersCmd).Standalone()
	rootCmd.AddCommand(layersCmd)
}
