package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dispatch_movefocusCmd = &cobra.Command{
	Use:   "movefocus",
	Short: "moves the focus in a direction",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dispatch_movefocusCmd).Standalone()

	carapace.Gen(dispatch_movefocusCmd).PositionalCompletion(
		carapace.ActionValues("left", "right", "up", "down"),
	)

	dispatchCmd.AddCommand(dispatch_movefocusCmd)
}
