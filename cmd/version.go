package cmd

import (
  "fmt"

  "github.com/spf13/cobra"
)

func init() {
  rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
  Use:   "version",
  Short: "Print the version number of mikapod-storage",
  Long:  `All software has versions.`,
  Run: func(cmd *cobra.Command, args []string) {
    fmt.Println("Mikapod-Storage v1.0")
  },
}
