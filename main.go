package main

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/markelog/eclectica/cmd/print"
	"github.com/markelog/linteum/linteum"
	"gopkg.in/alecthomas/kingpin.v2"
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
	paths = pathList(kingpin.Arg("path", "path to lint"))
)

func main() {
	kingpin.Version("0.0.1")
	kingpin.Parse()

	test := linteum.New(*paths)

	err := test.Lint()
	if err != nil {
		print.Error(err)
	}

	spew.Dump(err)
}
