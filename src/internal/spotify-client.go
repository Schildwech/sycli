package spotify

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "sycli",
	Short: "A simple CLI tool for Communication with Spotify",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Hello from mycli!")
	},
}

// Execute executes the root command.
func Execute() error {
	return rootCmd.Execute()
}
