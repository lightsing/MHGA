package rules

import (
	"github.com/lightsing/gcache"
	"path/filepath"
	"os"
	log "github.com/sirupsen/logrus"
	"errors"
	"sync"
)

type RuleSets struct {
	RuleSets []*RuleSet
	Targets gcache.Cache
	RuleCache gcache.Cache
	Lock sync.RWMutex
}

func NewRuleSets() *RuleSets {
	ruleSets := RuleSets{
		RuleSets: make([]*RuleSet, 0),
		Targets: nil,
		RuleCache: gcache.New(20).ARC().Build(),
	}
	ruleSets.Targets = gcache.New(20).ARC().LoaderFunc(func(test interface{}) (interface{}, error) {
		for _, ruleSet := range ruleSets.RuleSets {
			if ruleSet.Is(test.(string)) {
				return ruleSet, nil
			}
		}
		return nil, errors.New("applicable rule not found")
	}).Build()
	return &ruleSets
}

func LoadRuleSets(root string) (*RuleSets, error) {
	ruleSets := NewRuleSets()
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Warnf("WalkFunc Error, %s\n", err)
			return err
		}
		go func() {
			if ruleSet, err := LoadRuleSet(path); err == nil {
				log.Debugf("Adding: %s\n", path)
				ruleSets.Lock.Lock()
				ruleSets.RuleSets = append(ruleSets.RuleSets, ruleSet)
				ruleSets.Lock.Unlock()
			} else {
				// possibly caused by re2 bug(feature)
				// or not a xml file
				log.Infof("Parse rule fail, ignore %s (caused by [%s])", path, err)
			}
		}()
		return nil
	})
	if err != nil {
		log.Warnf("Walk Error, %s\n", err)
		return nil, err
	}
	log.Infof("Load \"%s\" Complete", root)
	return ruleSets, nil
}