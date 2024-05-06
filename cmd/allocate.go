package cmd

import (
	"github.com/spf13/cobra"
)

var allocateCmd = &cobra.Command{
	Use:   "allocate [resource_type]",
	Short: "Allocate command",
	Long:  `Allocate a resoure of the specified type.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runResource("allocate", args, nil)
	},
}

func init() {
	rootCmd.AddCommand(allocateCmd)
}
