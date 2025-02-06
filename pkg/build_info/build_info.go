package build_info

import (
	"fmt"
	"runtime"
)

// Set by goreleaser
var (
	Version = "develop"
	Date    = "unknown"
)

func GetDescription() string {
	return fmt.Sprintf("gossh - %s built on %s, with %s", Version, Date, runtime.Version())
}
