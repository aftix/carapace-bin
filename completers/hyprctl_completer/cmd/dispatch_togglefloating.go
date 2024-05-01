package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hypr"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var dispatch_togglefloatingCmd = &cobra.Command{
	Use:   "togglefloating",
	Short: "toggle the current window's floating state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dispatch_togglefloatingCmd).Standalone()

	carapace.Gen(dispatch_togglefloatingCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionStyledValuesDescribed("active", "the currently active window (default)", style.Blue),
			hypr.ActionWindows(),
		).ToA(),
	)

	dispatchCmd.AddCommand(dispatch_togglefloatingCmd)
}
