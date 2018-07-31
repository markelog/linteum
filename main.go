package main

import (
	"github.com/markelog/eclectica/cmd/print"
	"github.com/markelog/linteum/linteum"
	"github.com/markelog/linteum/rules"
	"gopkg.in/alecthomas/kingpin.v2"

	"github.com/markelog/linteum/config"
)

type list []string

func (p *list) Set(value string) error {
	*p = append(*p, value)
	return nil
}

func (p *list) String() string {
	return ""
}

func (p *list) IsCumulative() bool {
	return true
}

func pathList(s kingpin.Settings) (target *[]string) {
	target = new([]string)

	s.SetValue((*list)(target))

	return
}

var (
	paths = pathList(kingpin.Arg("path", "Path to lint"))
	conf  = kingpin.Flag("config", "Path to config").Default("linteum.yaml").String()
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()

	test := linteum.New(*paths)

	err := test.Lint()
	if err != nil {
		print.Error(err)
	}

	configuration, err := config.New(*conf)
	if err != nil {
		print.Error(err)
	}

	config := configuration.Rules
	data := map[string]interface{}{}

	for name, rule := range config {
		data[name] = rule
	}

	rules.New(data)
}
