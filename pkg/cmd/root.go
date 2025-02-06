package cmd

import (
	"fmt"
	"github.com/mason-rogers/gossh/pkg/build_info"
	"github.com/mason-rogers/gossh/pkg/config"
	"github.com/mason-rogers/gossh/pkg/menu"
	"github.com/mason-rogers/gossh/pkg/ssh"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"os"
)

func init() {
	rootCmd.Version = build_info.GetDescription()
	rootCmd.SetVersionTemplate(`{{printf .Version }}`)

	rootCmd.AddCommand(importCmd)
	rootCmd.AddCommand(versionCmd)

	importCmd.AddCommand(importTermiusCmd)
}

var rootCmd = &cobra.Command{
	Use:   "gossh [host]",
	Short: "SSH host manager",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error

		if len(args) > 0 {
			err = handleNonInteractive(args)
		} else {
			err = handleInteractive()
		}

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
			os.Exit(1)
		}

		return nil
	},
}

func Execute() error {
	return rootCmd.Execute()
}

func handleNonInteractive(args []string) error {
	host := config.FindHostByPath(args[0])
	if host == nil {
		return errors.New(fmt.Sprintf("Host %s not found", args[0]))
	}

	return ssh.ConnectToHost(*host)
}

func handleInteractive() error {
	host, err := menu.PromptForHost()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}

	return ssh.ConnectToHost(host)
}
