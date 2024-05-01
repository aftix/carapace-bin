package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dispatch_centerwindowCmd = &cobra.Command{
	Use:   "centerwindow",
	Short: "center the active window in the monitor (floating only)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dispatch_centerwindowCmd).Standalone()

	carapace.Gen(dispatch_centerwindowCmd).PositionalCompletion(
		carapace.ActionValuesDescribed("1", "respect monitor reserved area"),
	)

	dispatchCmd.AddCommand(dispatch_centerwindowCmd)
}
