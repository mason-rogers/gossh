package ssh

import (
	"fmt"
	"github.com/mason-rogers/gossh/pkg/build_info"
	"github.com/mason-rogers/gossh/pkg/config"
	"github.com/pkg/errors"
	"os"
	"os/exec"
	"strings"
)

func buildArgs(host config.Host) ([]string, error) {
	var args []string

	if host.User == "" {
		host.User = "root"
	}

	if host.Port == 0 {
		host.Port = 22
	}

	args = append(args, "-p", fmt.Sprintf("%d", host.Port))

	if host.KeyFile != "" {
		args = append(args, "-i", host.KeyFile)
	}

	if host.JumpHost != "" {
		jumpHost := config.Get().FindJumpHostByName(host.JumpHost)
		if jumpHost == nil {
			return []string{}, errors.New(fmt.Sprintf("Jump host [%s] not found.", host.JumpHost))
		}

		if jumpHost.Port == 0 {
			jumpHost.Port = 22
		}

		args = append(args, "-J", fmt.Sprintf("%s@%s:%d", jumpHost.User, jumpHost.Hostname, jumpHost.Port))
	}

	return append(args, fmt.Sprintf("%s@%s", host.User, host.Hostname)), nil
}

func ConnectToHost(host config.Host) error {
	args, err := buildArgs(host)
	if err != nil {
		return err
	}

	if build_info.RunningInDebug() {
		fmt.Printf("Executing `ssh %s`\n", strings.Join(args, " "))
	}

	cmd := exec.Command("ssh", args...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}

	return nil
}
