package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicesCmd = &cobra.Command{
	Use:     "devices",
	Short:   "list all devices with their properties",
	GroupID: "info",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicesCmd).Standalone()
	rootCmd.AddCommand(devicesCmd)
}
