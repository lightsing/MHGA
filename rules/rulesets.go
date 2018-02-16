package rules

import (
	"github.com/lightsing/gcache"
	"path/filepath"
	"os"
	log "github.com/sirupsen/logrus"
)

type RuleSets struct {
	RuleSets []*RuleSet
	Targets gcache.Cache
	RuleCache gcache.Cache
}

func NewRuleSets() *RuleSets {
	return &RuleSets{
		RuleSets: make([]*RuleSet, 0),
		Targets: gcache.New(20).ARC().Build(),
		RuleCache: gcache.New(20).ARC().Build(),
	}
}

func LoadRuleSets(root string) (*RuleSets, error) {
	ruleSets := NewRuleSets()
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Warnf("WalkFunc Error, %s\n", err)
			return err
		}
		if ruleSet, err := LoadRuleSet(path); err == nil {
			log.Debugf("Adding: %s\n", path)
			ruleSets.RuleSets = append(ruleSets.RuleSets, ruleSet)
		} else {
			// possibly caused by re2 bug(feature)
			// or not a xml file
			log.Infof("Parse rule fail, ignore %s (caused by [%s])", path, err)
		}
		return nil
	})
	if err != nil {
		log.Warnf("Walk Error, %s\n", err)
		return nil, err
	}
	log.Infof("Load \"%s\" Complete", root)
	return ruleSets, nil
}