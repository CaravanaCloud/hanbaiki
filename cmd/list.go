package cmd

import (
	"github.com/spf13/cobra"
)

var ou string

var listCmd = &cobra.Command{
	Use:   "list [resource_type]",
	Short: "list command",
	Long:  `list a resoure of the specified type.`,
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		data := map[string]string{
			"ou": ou,
		}
		runResource("list", args, data)
	},
}

func init() {
	listCmd.Flags().StringVarP(&ou, "ou", "o", "", "Name of the organizational unit")
	rootCmd.AddCommand(listCmd)
}
