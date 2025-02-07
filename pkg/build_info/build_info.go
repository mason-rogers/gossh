package build_info

import (
	"fmt"
	"github.com/mason-rogers/gossh/pkg/config"
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
	return os.Getenv("GOSSH_DEBUG") == "true" || config.Get().Debug
}
