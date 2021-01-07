package cmd

import (
	"fmt"
	"os"

	// homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	// "github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "mikapod-storage",
	Short: "Local storage application for IoT project",
	Long: `The purpose of this application is to provide a local database tailored for storage of
  data for a **Mikapod-Soil** device and be accessible with remote procedure calls (RPC). `,
	Run: func(cmd *cobra.Command, args []string) {
		// Do nothing...
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
