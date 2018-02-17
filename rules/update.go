package rules

import (
	"github.com/lightsing/makehttps/config"
	"gopkg.in/src-d/go-git.v4"
	"os"
	"gopkg.in/src-d/go-git.v4/plumbing"
	"errors"
)
func CheckRule(ruleConfig *config.RuleConfig) error {
	switch ruleConfig.Type {
	case "git":
		return checkGitRule(ruleConfig)
	}
	return errors.New("non-supported rule type")
}

func checkGitRule(ruleConfig *config.RuleConfig) error {
	if _, err := os.Stat(ruleConfig.Git.Path); err == nil && ruleConfig.Update {
		repo, err := git.PlainOpen(ruleConfig.Git.Path);
		if err != nil {
			// try fix error
			err = os.RemoveAll(ruleConfig.Git.Path)
			if err == nil {
				return gitClone(ruleConfig)
			} else {
				return err
			}
		}
		if worktree, err := repo.Worktree(); err == nil {
			return worktree.Pull(&git.PullOptions{
				SingleBranch: ruleConfig.Git.SingleBranch,
				ReferenceName: plumbing.ReferenceName(ruleConfig.Git.Branch),
				Depth: ruleConfig.Git.Depth,
			})
		} else {
			return err
		}
	}
	return gitClone(ruleConfig)
}

func gitClone(ruleConfig *config.RuleConfig) error {
	_, err := git.PlainClone(ruleConfig.Git.Path, false, &git.CloneOptions{
		URL: ruleConfig.Git.Upstream,
		SingleBranch: ruleConfig.Git.SingleBranch,
		ReferenceName: plumbing.ReferenceName(ruleConfig.Git.Branch),
		Depth: ruleConfig.Git.Depth,
	})
	if err != nil {
		return err
	}
}