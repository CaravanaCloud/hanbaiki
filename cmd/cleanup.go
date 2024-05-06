package cmd

import (
	"github.com/spf13/cobra"
)

var cleanupCmd = &cobra.Command{
	Use:   "cleanup [resource_type]",
	Short: "cleanup command",
	Long:  `cleanup a resoure of the specified type.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runResource("cleanup", args, nil)
	},
}

func init() {
	rootCmd.AddCommand(cleanupCmd)
}
