package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/hypr"
	"github.com/spf13/cobra"
)

var dispatch_workspaceCmd = &cobra.Command{
	Use:   "workspace",
	Short: "changes the workspace",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dispatch_workspaceCmd).Standalone()

	carapace.Gen(dispatch_workspaceCmd).PositionalCompletion(hypr.ActionWorkspaces())

	dispatchCmd.AddCommand(dispatch_workspaceCmd)
}
