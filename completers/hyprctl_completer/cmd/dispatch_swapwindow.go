package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dispatch_swapwindowCmd = &cobra.Command{
	Use:   "swapwindow",
	Short: "swaps the active window with another window in the given direction",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dispatch_swapwindowCmd).Standalone()

	carapace.Gen(dispatch_swapwindowCmd).PositionalCompletion(
		carapace.ActionValues("left", "right", "up", "down"),
	)

	dispatchCmd.AddCommand(dispatch_swapwindowCmd)
}
