package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import hosts from other sources",
}

var importTermiusCmd = &cobra.Command{
	Use:   "termius",
	Short: "Import hosts from Termius",
	RunE:  runImportTermius,
}

func runImportTermius(cmd *cobra.Command, args []string) error {
	fmt.Println("Termius Import support coming soon!")

	return nil
}
