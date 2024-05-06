package cmd

import (
	"github.com/spf13/cobra"
)

var leaseCmd = &cobra.Command{
	Use:   "lease [resource_type]",
	Short: "lease command",
	Long:  `lease a resoure of the specified type.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runResource("lease", args, nil)
	},
}

func init() {
	rootCmd.AddCommand(leaseCmd)
}
