package cmd

import (
	"fmt"

	. "github.com/stuart-pollock/bosh-cli/cmd/opts"
	boshtbl "github.com/stuart-pollock/bosh-cli/ui/table"
)

type Cmd struct {
	BoshOpts BoshOpts
	Opts     interface{}

	deps BasicDeps
}

func NewCmd(boshOpts BoshOpts, opts interface{}, deps BasicDeps) Cmd {
	return Cmd{boshOpts, opts, deps}
}

type cmdConveniencePanic struct {
	Err error
}

func (c Cmd) Execute() (cmdErr error) {
	// Catch convenience panics from panicIfErr
	defer func() {
		if r := recover(); r != nil {
			if cp, ok := r.(cmdConveniencePanic); ok {
				cmdErr = cp.Err
			} else {
				panic(r)
			}
		}
	}()

	c.configureUI()

	deps := c.deps

	switch opts := c.Opts.(type) {
	case *InterpolateOpts:
		return NewInterpolateCmd(deps.UI).Run(*opts)

	default:
		return fmt.Errorf("Unhandled command: %#v", c.Opts)
	}
}
func (c Cmd) configureUI() {
	c.deps.UI.EnableTTY(c.BoshOpts.TTYOpt)

	if !c.BoshOpts.NoColorOpt {
		c.deps.UI.EnableColor()
	}

	if c.BoshOpts.JSONOpt {
		c.deps.UI.EnableJSON()
	}

	if c.BoshOpts.NonInteractiveOpt {
		c.deps.UI.EnableNonInteractive()
	}

	if len(c.BoshOpts.ColumnOpt) > 0 {
		headers := []boshtbl.Header{}
		for _, columnOpt := range c.BoshOpts.ColumnOpt {
			headers = append(headers, columnOpt.Header)
		}

		c.deps.UI.ShowColumns(headers)
	}
}

func (c Cmd) panicIfErr(err error) {
	if err != nil {
		panic(cmdConveniencePanic{err})
	}
}
