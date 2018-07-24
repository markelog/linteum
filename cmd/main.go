package paths

import kingpin "gopkg.in/alecthomas/kingpin.v2"

type paths []string

func (p *paths) Set(value string) error {
	*p = append(*p, value)
	return nil
}

func (p *paths) String() string {
	return ""
}

func (p *paths) IsCumulative() bool {
	return true
}

func Paths(s kingpin.Settings) (target *[]string) {
	target = new([]string)
	s.SetValue((*paths)(target))

	return
}
