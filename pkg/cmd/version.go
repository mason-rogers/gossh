package cmd

import (
	"fmt"
	"github.com/mason-rogers/gossh/pkg/build_info"
	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print version information",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(build_info.GetDescription())
	},
}
