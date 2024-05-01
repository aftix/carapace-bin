package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hyprctl",
	Short: "Utility for controlling parts of Hyprland from a CLI or a script",
	Long:  "https://github.com/HyprWM/Hyprland",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	rootCmd.AddGroup(
		&cobra.Group{ID: "control", Title: "Control Commands"},
		&cobra.Group{ID: "info", Title: "Info Commands"},
	)

	// TODO: prevent positional completion when this flag is used
	rootCmd.Flags().String("batch", "", "Specify a batch of commands to execute")

	rootCmd.PersistentFlags().BoolS("json", "j", false, "Outputs information in JSON")

	carapace.Gen(rootCmd).FlagCompletion(carapace.ActionMap{
		"batch": carapace.ActionMultiParts(";", func(c carapace.Context) carapace.Action {
			return carapace.ActionCommands(rootCmd)
		}),
	})
}
