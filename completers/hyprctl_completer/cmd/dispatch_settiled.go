package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hypr"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var dispatch_settiledCmd = &cobra.Command{
	Use:   "settiled",
	Short: "set the current window to tiled",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dispatch_settiledCmd).Standalone()

	carapace.Gen(dispatch_settiledCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionStyledValuesDescribed("active", "the currently active window (default)", style.Blue),
			hypr.ActionWindows(),
		).ToA(),
	)

	dispatchCmd.AddCommand(dispatch_settiledCmd)
}
