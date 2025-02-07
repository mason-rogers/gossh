package main

import (
	"github.com/mason-rogers/gossh/pkg/cmd"
	"os"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
