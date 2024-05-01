package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hypr"
	"github.com/spf13/cobra"
)

var dispatch_movetoworkspacesilentCmd = &cobra.Command{
	Use:   "movetoworkspacesilent",
	Short: "move the focused window to a workspace but don't switch the active workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dispatch_movetoworkspacesilentCmd).Standalone()

	carapace.Gen(dispatch_movetoworkspacesilentCmd).PositionalCompletion(hypr.ActionWorkspaces())

	dispatchCmd.AddCommand(dispatch_movetoworkspacesilentCmd)
}
