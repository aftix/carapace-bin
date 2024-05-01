package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hypr"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var dispatch_setfloatingCmd = &cobra.Command{
	Use:   "setfloating",
	Short: "set the current window to floating",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dispatch_setfloatingCmd).Standalone()

	carapace.Gen(dispatch_setfloatingCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionStyledValuesDescribed("active", "the currently active window (default)", style.Blue),
			hypr.ActionWindows(),
		).ToA(),
	)

	dispatchCmd.AddCommand(dispatch_setfloatingCmd)
}
