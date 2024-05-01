package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dispatch_movewindowpixelCmd = &cobra.Command{
	Use:   "movewindowpixel",
	Short: "moves a selected window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dispatch_movewindowpixelCmd).Standalone()
	dispatchCmd.AddCommand(dispatch_movewindowpixelCmd)
}
