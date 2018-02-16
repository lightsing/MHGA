package main

import (
	"github.com/lightsing/makehttps/rules"
	log "github.com/sirupsen/logrus"
)

func main() {
	// Only log the warning severity or above.
	log.SetLevel(log.InfoLevel)
	rules.LoadRuleSets("rules/rules")
}