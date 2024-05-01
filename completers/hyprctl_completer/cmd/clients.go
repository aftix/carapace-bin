package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var clientsCmd = &cobra.Command{
	Use:     "clients",
	Short:   "list all clients with their properties",
	GroupID: "info",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(clientsCmd).Standalone()
	rootCmd.AddCommand(clientsCmd)
}
