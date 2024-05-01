package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var dispatch_execCmd = &cobra.Command{
	Use:   "exec",
	Short: "executes a shell command",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dispatch_execCmd).Standalone()

	// TODO: add rules completion
	carapace.Gen(dispatch_execCmd).PositionalCompletion(carapace.ActionExecutables())
	carapace.Gen(dispatch_execCmd).PositionalAnyCompletion(bridge.ActionCarapaceBin())

	dispatchCmd.AddCommand(dispatch_execCmd)
}
