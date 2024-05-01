package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hypr"
	"github.com/spf13/cobra"
)

var dispatch_passCmd = &cobra.Command{
	Use:   "pass",
	Short: "passes the key (with mods) to a specified window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dispatch_passCmd).Standalone()

	carapace.Gen(dispatch_passCmd).PositionalCompletion(hypr.ActionWindows())

	dispatchCmd.AddCommand(dispatch_passCmd)
}
