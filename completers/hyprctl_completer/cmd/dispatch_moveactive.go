package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dispatch_moveactiveCmd = &cobra.Command{
	Use:   "moveactive",
	Short: "moves the active window using cordinates",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dispatch_moveactiveCmd).Standalone()

	carapace.Gen(dispatch_moveactiveCmd).PositionalCompletion(
		carapace.ActionValuesDescribed("exact", "set the size using absolute units instead of relative"),
	)

	dispatchCmd.AddCommand(dispatch_moveactiveCmd)
}
