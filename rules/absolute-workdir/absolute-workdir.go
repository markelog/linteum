package absoluteWorkdir

type Rule struct{}

type arguments struct {
	enable bool
}

func (rule *Rule) New() string {
	return "absolute-workdir"
}

func (rule *Rule) Name() string {
	return "absolute-workdir"
}

func (rule *Rule) Visit() error {
	return nil
}
