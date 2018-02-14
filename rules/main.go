package main

import (
"fmt"
"encoding/xml"
"os"
	"github.com/ryanuber/go-glob"
	"regexp"
)

type Target struct {
	Host string `xml:"host,attr"`
}

type Exclusion struct {
	Pattern string `xml:"pattern,attr"`
}

type Rule struct {
	From string `xml:"from,attr"`
	To string `xml:"to,attr"`
	FromRe *regexp.Regexp
}

type RuleSet struct {
	Targets []Target `xml:"target"`
	Exclusions []Exclusion `xml:"exclusion"`
	Rules []Rule `xml:"rule"`
}

type Tester interface {
	Is(string) bool
}

func (t *Target) Is(test string) bool {
	return glob.Glob(t.Host, test)
}

func (e *Exclusion) Is(test string) bool {
	return glob.Glob(e.Pattern, test)
}

func (r *Rule) Init() *Rule {
	r.FromRe = regexp.MustCompile(r.From)
	return r
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
		for i := range ruleSet.Rules {
			ruleSet.Rules[i].Init()
		}
		fmt.Printf("%v\n", ruleSet)
		for _, target := range ruleSet.Targets {
			fmt.Println(target.Is("www.google.com.hk"))
		}
	}

}
