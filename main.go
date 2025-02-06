package main

import (
	"github.com/mason-rogers/gossh/pkg/cmd"
	"github.com/mason-rogers/gossh/pkg/config"
	"os"
)

func main() {
	config.Load()

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
