package rules

import (
	"errors"
	"github.com/lightsing/gcache"
	log "github.com/sirupsen/logrus"
	"os"
	"path/filepath"
	"sync"
)

type RuleSets struct {
	RuleSets  []*RuleSet
	Targets   gcache.Cache
	RuleCache gcache.Cache
	Lock      sync.RWMutex
}

func (rs *RuleSets) get(test string) (*RuleSet, bool) {
	if ruleSet, err := rs.Targets.Get(test); err == nil {
		return ruleSet.(*RuleSet), true
	}
	return nil, false
}

func (rs *RuleSets) Apply(test string) (*string, bool) {
	if result, err := rs.RuleCache.Get(test); err == nil {
		log.Warnf("Cache hit for %s", test)
		if result, ok := result.(*string); ok {
			return result, true
		} else {
			return nil, false
		}
	}
	log.Warnf("Cache miss for %s", test)
	if ruleSet, ok := rs.get(test); ok {
		if result, ok := ruleSet.Apply(test); ok {
			rs.RuleCache.Set(test, result)
			return result, true
		}
	}
	rs.RuleCache.Set(test, nil)
	return nil, false
}

func NewRuleSets() *RuleSets {
	ruleSets := RuleSets{
		RuleSets:  make([]*RuleSet, 0),
		Targets:   nil,
		RuleCache: gcache.New(20).ARC().Build(),
	}
	ruleSets.Targets = gcache.New(20).ARC().LoaderFunc(func(test interface{}) (interface{}, error) {
		for _, ruleSet := range ruleSets.RuleSets {
			if ruleSet.Is(test.(string)) {
				//log.Warnf("%v", ruleSet)
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
