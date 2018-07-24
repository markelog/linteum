// lintuem package
package linteum

type Linteum struct {
	paths []string
}

func New(paths []string) *Linteum {
	linteum := &Linteum{
		paths: paths,
	}

	return linteum
}

func (linteum *Linteum) Lint() *Linteum {
	for _, path := range linteum.paths {
		println(path)
	}

	return nil
}
