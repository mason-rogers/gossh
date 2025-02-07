package cmd

import (
	"fmt"
	"github.com/mason-rogers/gossh/pkg/config"
	"github.com/spf13/cobra"
)

var importCmd = &cobra.Command{
	Use:   "import",
	Short: "Import hosts from other sources",
}

var importTermiusCmd = &cobra.Command{
	Use:   "termius",
	Short: "Import hosts from Termius",
	PreRun: func(cmd *cobra.Command, args []string) {
		config.Load()
	},
	Run: runImportTermius,
}

func runImportTermius(cmd *cobra.Command, args []string) {
	fmt.Println("Termius Import support coming soon!")
}
