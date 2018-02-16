package main

import (
	"github.com/lightsing/makehttps/rules"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	// Only log the warning severity or above.
	log.SetLevel(log.WarnLevel)
	start := time.Now()
	if ruleSets, err := rules.LoadRuleSets("rules/rules"); err == nil {
		log.Warnf("Load all rule in %s", time.Since(start))
		start = time.Now()
		ruleSet, err := ruleSets.Targets.Get("www.google.com.hk")
		if err != nil {
			log.Errorf("Error when get rule: %s", err)
		} else {
			ruleSet = ruleSet.(*rules.RuleSet)
			log.Warnf("Found rule (%s):\n\t %v", time.Since(start), ruleSet)
		}
	}

}