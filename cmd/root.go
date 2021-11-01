package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:           "avwx",
		Short:         "A simple tool to access METAR and TAF data.",
		Long:          "AVWX is a tool to access aviation weather data (METARs, TAFs) provided by the Aviation Weather Center's Text Data Server.",
		SilenceErrors: true,
	}
)

func init() {

	rootCmd.CompletionOptions.DisableDefaultCmd = true
	rootCmd.AddCommand(metarCmd)
	rootCmd.AddCommand(tafCmd)

	// remove default help command
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:    "no-help",
		Hidden: true,
	})

}

// Execute starts the root AVWX command
func Execute() error {
	return rootCmd.Execute()
}
