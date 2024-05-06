package cmd

import (
	"github.com/spf13/cobra"
)

var disposeCmd = &cobra.Command{
	Use:   "dispose [resource_type]",
	Short: "dispose command",
	Long:  `dispose a resoure of the specified type.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		runResource("dispose", args, nil)
	},
}

func init() {
	rootCmd.AddCommand(disposeCmd)
}
