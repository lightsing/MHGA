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
		test := "http://www.google.com.hk"
		for range make([]int, 10) {
			start = time.Now()
			if result, ok := ruleSets.Apply(test); ok {
				log.Warnf("Apply rule [%s], from (%s) to (%s)", time.Since(start), test, *result)
			}
		}

	}

}
