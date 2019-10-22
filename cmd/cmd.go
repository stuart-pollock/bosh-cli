package cmd

import (
	"fmt"

	. "github.com/stuart-pollock/bosh-cli/cmd/opts"
	"github.com/stuart-pollock/bosh-cli/ui/table"
)

type Cmd struct {
	MainOpts MainOpts
	Opts     interface{}

	deps BasicDeps
}

func NewCmd(mainOpts MainOpts, opts interface{}, deps BasicDeps) Cmd {
	return Cmd{mainOpts, opts, deps}
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
	c.deps.UI.EnableTTY(c.MainOpts.TTYOpt)

	if !c.MainOpts.NoColorOpt {
		c.deps.UI.EnableColor()
	}

	if c.MainOpts.JSONOpt {
		c.deps.UI.EnableJSON()
	}

	if c.MainOpts.NonInteractiveOpt {
		c.deps.UI.EnableNonInteractive()
	}

	if len(c.MainOpts.ColumnOpt) > 0 {
		headers := []table.Header{}
		for _, columnOpt := range c.MainOpts.ColumnOpt {
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
