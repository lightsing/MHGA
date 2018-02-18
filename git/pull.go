package git

import (
	"github.com/lightsing/makehttps/config"
	"github.com/sirupsen/logrus"
	"os/exec"
)

func Update(config config.GitConfig) error {
	// git pull --rebase --stat origin master
	logrus.Infof("Updating %s:%s", gitNameRegex.FindString(config.Upstream), config.Branch)
	args := []string{"pull", "--rebase", "--stat", "origin", config.Branch}
	cmd := exec.Command("git", args...)
	cmd.Dir = config.Path
	go pipeStdout(cmd.StderrPipe())
	go pipeStdout(cmd.StdoutPipe())
	return cmd.Run()
}
