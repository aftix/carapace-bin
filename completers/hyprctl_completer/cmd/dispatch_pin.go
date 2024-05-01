package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hypr"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var dispatch_pinCmd = &cobra.Command{
	Use:   "pin",
	Short: "pins a window so it can be seen on every workspace (floating only)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dispatch_pinCmd).Standalone()

	carapace.Gen(dispatch_pinCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionStyledValuesDescribed("active", "the currently active window (default)", style.Blue),
			hypr.ActionWindows(),
		).ToA(),
	)

	dispatchCmd.AddCommand(dispatch_pinCmd)
}
