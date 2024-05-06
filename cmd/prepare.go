package cmd

import (
	"github.com/spf13/cobra"
)

var prepareCmd = &cobra.Command{
	Use:   "prepare [resource_type]",
	Short: "prepare command",
	Long:  `prepare a resoure of the specified type.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runResource("prepare", args, nil)
	},
}

func init() {
	rootCmd.AddCommand(prepareCmd)
}
