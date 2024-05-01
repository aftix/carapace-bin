package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dispatch_fullscreenCmd = &cobra.Command{
	Use:   "fullscreen",
	Short: "toggles the focused window's fullscreen state",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dispatch_fullscreenCmd).Standalone()

	carapace.Gen(dispatch_fullscreenCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"0", "fullscreen",
			"1", "maximize (keeps gaps and bars)",
			"2", "fullscreen without altering window internal state",
		),
	)

	dispatchCmd.AddCommand(dispatch_fullscreenCmd)
}
