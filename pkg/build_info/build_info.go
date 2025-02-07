package build_info

import (
	"fmt"
	"os"
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

func RunningInDebug() bool {
	return os.Getenv("GOSSH_DEBUG") == "true"
}
