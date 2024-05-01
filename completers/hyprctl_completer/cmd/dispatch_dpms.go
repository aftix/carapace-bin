package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hypr"
	"github.com/spf13/cobra"
)

var dispatch_dpmsCmd = &cobra.Command{
	Use:   "dpms",
	Short: "sets all monitors' DPMS status",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dispatch_dpmsCmd).Standalone()

	carapace.Gen(dispatch_dpmsCmd).PositionalCompletion(
		carapace.ActionValues("on", "off", "toggle"),
		hypr.ActionMonitors(),
	)

	dispatchCmd.AddCommand(dispatch_dpmsCmd)
}
