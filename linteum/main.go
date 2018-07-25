// lintuem package
package linteum

import (
	"io"
	"os"

	"github.com/davecgh/go-spew/spew"
	"github.com/moby/buildkit/frontend/dockerfile/parser"
)

type Linteum struct {
	paths []string
}

func New(paths []string) *Linteum {
	linteum := &Linteum{
		paths: paths,
	}

	return linteum
}

func (linteum *Linteum) Parse(file io.Reader) (*parser.Node, error) {
	result, err := parser.Parse(file)
	if err != nil {
		return nil, nil
	}

	ast := result.AST

	return ast, nil
}

func (linteum *Linteum) Check(file io.Reader) (*parser.Node, error) {
	ast, err := linteum.Parse(file)
	if err != nil {
		return nil, err
	}

	spew.Dump(ast)
	return nil, nil
}

func (linteum *Linteum) Lint() error {
	for _, path := range linteum.paths {
		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		_, err = linteum.Check(file)
		if err != nil {
			return err
		}
	}

	return nil
}
