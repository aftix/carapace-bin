package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/spf13/cobra"
)

var dispatch_execrCmd = &cobra.Command{
	Use:   "exec",
	Short: "executes a raw shell command (does not support rules)",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dispatch_execrCmd).Standalone()

	carapace.Gen(dispatch_execrCmd).PositionalCompletion(carapace.ActionExecutables())
	carapace.Gen(dispatch_execrCmd).PositionalAnyCompletion(bridge.ActionCarapaceBin())

	dispatchCmd.AddCommand(dispatch_execrCmd)
}
