package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dispatch_resizeactiveCmd = &cobra.Command{
	Use:   "resizeactive",
	Short: "resizes the active window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dispatch_resizeactiveCmd).Standalone()

	carapace.Gen(dispatch_resizeactiveCmd).PositionalCompletion(
		carapace.ActionValuesDescribed("exact", "set the size using absolute units instead of relative"),
	)

	dispatchCmd.AddCommand(dispatch_resizeactiveCmd)
}
