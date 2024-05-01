package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hypr"
	"github.com/spf13/cobra"
)

var dispatch_closewindowCmd = &cobra.Command{
	Use:   "closewindow",
	Short: "closes a specified window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dispatch_closewindowCmd).Standalone()

	carapace.Gen(dispatch_closewindowCmd).PositionalCompletion(hypr.ActionWindows())

	dispatchCmd.AddCommand(dispatch_closewindowCmd)
}
