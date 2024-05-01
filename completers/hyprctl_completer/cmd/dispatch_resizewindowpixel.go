package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dispatch_resizewindowpixelCmd = &cobra.Command{
	Use:   "resizewindowpixel",
	Short: "resizes a given window",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dispatch_resizewindowpixelCmd).Standalone()
	dispatchCmd.AddCommand(dispatch_resizewindowpixelCmd)
}
