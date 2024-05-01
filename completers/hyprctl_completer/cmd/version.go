package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "prints the Hyprland version, flags, commit and branch of build",
	GroupID: "info",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionCmd).Standalone()
	rootCmd.AddCommand(versionCmd)
}
