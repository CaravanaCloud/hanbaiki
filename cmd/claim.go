package cmd

import (
	"github.com/spf13/cobra"
)

var claimCmd = &cobra.Command{
	Use:   "claim [resource_type]",
	Short: "claim command",
	Long:  `claim a resoure of the specified type.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runResource("claim", args, nil)
	},
}

func init() {
	rootCmd.AddCommand(claimCmd)
}
