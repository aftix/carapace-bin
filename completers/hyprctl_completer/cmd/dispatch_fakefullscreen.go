package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dispatch_fakefullscreenCmd = &cobra.Command{
	Use:   "fakefullscreen",
	Short: "toggles the active window's internal fullscreen state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dispatch_fakefullscreenCmd).Standalone()
	dispatchCmd.AddCommand(dispatch_fakefullscreenCmd)
}
