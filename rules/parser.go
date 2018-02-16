package rules

import (
	"encoding/xml"
	"github.com/ryanuber/go-glob"
	"regexp"
	"os"
	"errors"
)

type Target struct {
	Host string `xml:"host,attr"`
}

type Exclusion struct {
	Pattern string `xml:"pattern,attr"`
}

type Rule struct {
	From   string `xml:"from,attr"`
	To     string `xml:"to,attr"`
	FromRe *regexp.Regexp
}

type RuleSet struct {
	Targets    []Target    `xml:"target"`
	Exclusions []Exclusion `xml:"exclusion"`
	Rules      []Rule      `xml:"rule"`
}

type Tester interface {
	Is(string) bool
}

var hostRegexp = regexp.MustCompile(`(?im)^(https?://)?([0-9a-zA-Z.]+).*$`)

func hostGlob(pattern, subj string) bool {
	match := hostRegexp.FindStringSubmatch(subj)
	return glob.Glob(pattern, match[len(match)-1])
}

func (t *Target) Is(test string) bool {
	return hostGlob(t.Host, test)
}

func (e *Exclusion) Is(test string) bool {
	return hostGlob(e.Pattern, test)
}

func (rs *RuleSet) Is(test string) bool {
	for _, exclusion := range rs.Exclusions {
		if exclusion.Is(test) {
			return false
		}
	}
	for _, target := range rs.Targets {
		if target.Is(test) {
			return true
		}
	}
	return false
}

func (r *Rule) Init() error {
	var err error
	if r.FromRe, err = regexp.Compile(r.From); err != nil{
		return errors.New("regex parse error")
	}
	return nil
}

func (r *Rule) Apply(uri string) string {
	return r.FromRe.ReplaceAllString(uri, r.To);
}

func (rs *RuleSet) Apply(urispec string) (*string, bool){
	for _, exclusion := range rs.Exclusions {
		if exclusion.Is(urispec) {
			return nil, false
		}
	}
	for _, target := range rs.Targets {
		if target.Is(urispec) {
			for _, rule := range rs.Rules {
				if result := rule.Apply(urispec); result != urispec {
					return &result, true
				}
			}
		}
	}
	return nil, false
}

func LoadRuleSet(any interface{}) (*RuleSet, error){
	if name, ok := any.(string); ok {
		if reader, err := os.Open(name); err != nil{
			return nil, err
		} else {
			return LoadRuleSet(reader)
		}
	} else if reader, ok := any.(*os.File); ok {
		var ruleSet RuleSet
		if err := xml.NewDecoder(reader).Decode(&ruleSet); err != nil {
			return nil, err
		} else {
			for i := range ruleSet.Rules {
				if err := ruleSet.Rules[i].Init(); err != nil {
					return nil, err
				}
			}
			return &ruleSet, nil
		}
	} else {
		return nil, errors.New("arg type error")
	}
}
