package rules

import (
	"errors"

	"github.com/davecgh/go-spew/spew"
	"github.com/markelog/linteum/rules/absolute-workdir"
)

type Rule interface {
	Name() string
	Visit() error
}

var rules = map[string]Rule{
	"absolute-workdir": &absoluteWorkdir.Rule{},
}

func New(initiate map[string]interface{}) ([]*Rule, error) {
	result := make([]*Rule, 0)

	for name, _ := range rules {
		arguments, ok := initiate[name]

		if ok == false {
			return nil, errors.New("Rule " + name + " does not exist")
		}

		// if arguments.enable.(bool) == false {
		// 	continue
		// }

		spew.Dump(arguments.(struct{ enable bool }))

		// 	result = append(result, rule)
	}

	return result, nil
}
