package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hypr"
	"github.com/spf13/cobra"
)

var dispatch_movetoworkspaceCmd = &cobra.Command{
	Use:   "movetoworkspace",
	Short: "moves the focused window to a workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dispatch_movetoworkspaceCmd).Standalone()

	carapace.Gen(dispatch_movetoworkspaceCmd).PositionalCompletion(hypr.ActionWorkspaces())

	dispatchCmd.AddCommand(dispatch_movetoworkspaceCmd)
}
