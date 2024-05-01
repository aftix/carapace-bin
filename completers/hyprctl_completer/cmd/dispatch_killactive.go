package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dispatch_killactiveCmd = &cobra.Command{
	Use:   "killactive",
	Short: "closes (not kills) the active window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dispatch_killactiveCmd).Standalone()
	dispatchCmd.AddCommand(dispatch_killactiveCmd)
}
