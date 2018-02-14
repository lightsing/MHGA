package main

import (
"fmt"
"encoding/xml"
"os"
)

type RuleSet struct {
	Targets []struct {
		Host string `xml:"host,attr"`
	} `xml:"target"`
	Exclusions []struct {
		Pattern string `xml:"pattern,attr"`
	} `xml:"exclusion"`
	Rules []struct {
		From string `xml:"from,attr"`
		To string `xml:"to,attr"`
	} `xml:"rule"`
}

func main() {
	reader, err:= os.Open("rules/rules/rules/Google.xml")
	if err != nil {
		panic(err)
	}
	var ruleSet RuleSet
	if err := xml.NewDecoder(reader).Decode(&ruleSet); err != nil {
		panic(err)
	} else {
		fmt.Printf("%v\n", ruleSet)
	}

}
