package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var activewindowCmd = &cobra.Command{
	Use:     "activewindow",
	Short:   "print the active window name",
	GroupID: "info",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(activewindowCmd).Standalone()
	rootCmd.AddCommand(activewindowCmd)
}
