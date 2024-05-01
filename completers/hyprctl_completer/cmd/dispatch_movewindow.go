package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hypr"
	"github.com/spf13/cobra"
)

var dispatch_movewindowCmd = &cobra.Command{
	Use:   "movewindow",
	Short: "moves the active window in a direction or to a monitor",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dispatch_movewindowCmd).Standalone()

	carapace.Gen(dispatch_movewindowCmd).PositionalCompletion(
		carapace.Batch(
			carapace.ActionValues("left", "right", "up", "down"),
			hypr.ActionMonitors().Prefix("mon:"),
		).ToA(),
		carapace.ActionValuesDescribed("silent", "prevent focus being moved"),
	)

	dispatchCmd.AddCommand(dispatch_movewindowCmd)
}
