package cmd

import (
	"os"

	"github.com/charmbracelet/log"

	"github.com/joho/godotenv"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "hbk",
	Short: "A resource leasing manager",
	Long:  `Provisions, leases and clean up resources, such as cloud accounts, instances and containers.`,
	// Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	err = rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
