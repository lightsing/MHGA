package git

import (
	"os/exec"
	"github.com/lightsing/makehttps/config"
	"strconv"
)

func Clone(config config.GitConfig) error {
	args := []string{ "clone",
		"--depth", strconv.Itoa(config.Depth),
		"-b", config.Branch,
		config.Upstream, config.Path,
		}
	cmd := exec.Command("git", args...)
	go pipeStdout(cmd.StderrPipe())
	go pipeStdout(cmd.StdoutPipe())
	return cmd.Run()
}