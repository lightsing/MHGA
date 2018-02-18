package rules

import (
	"github.com/lightsing/makehttps/git"
	"github.com/lightsing/makehttps/config"
	"os"
	"errors"
	log "github.com/sirupsen/logrus"
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
		if git.Update(ruleConfig.Git) != nil {
			log.Errorf("Updating [%s] fail", ruleConfig.Name)
			// ignore error
		}
		return nil
	}
	return git.Clone(ruleConfig.Git)
}