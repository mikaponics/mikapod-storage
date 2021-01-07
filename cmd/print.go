package cmd

import (
	"fmt"
    "errors"
	"time"

	"github.com/spf13/cobra"

	"github.com/mikaponics/mikapod-storage/internal/storage"
)

func init() {
	rootCmd.AddCommand(printCmd)
	printCmd.PersistentFlags().Int32P("lines", "", 0, "Set number count of latest records to return")
}

func runPrintCmd(limit int32) {
	storage := storage.InitMikapodDB()
	data := storage.ListTimeSeriesData(limit)
	for _, datum := range(data) {
		if datum.Id != 0 {
			fmt.Println("ID:", datum.Id,"| Instrument:", datum.Instrument, "| Timestamp:", time.Unix(datum.Timestamp, 0), "| Value:", datum.Value)
		}
	}
}

var printCmd = &cobra.Command{
	Use:   "print",
	Short: "Prints latest records",
	Long:  `Prints the latest records in the database`,
    Args: func(cmd *cobra.Command, args []string) error {
        // Check to see if something was entered and if not then we need to generate an error.
		lines, _:= cmd.Flags().GetInt32("lines")
		if lines == 0 {
			return errors.New("requires a `--lines` argument")
		}

		// If we excute to this part of code then we have generated no errors.
        return nil
    },
	Run: func(cmd *cobra.Command, args []string) {
		lines, _:= cmd.Flags().GetInt32("lines")
		runPrintCmd(lines)
	},
}
