package cmd

import (
	. "github.com/stuart-pollock/bosh-cli/cmd/opts"
	boshdir "github.com/stuart-pollock/bosh-cli/director"
)

type UpdateResurrectionCmd struct {
	director boshdir.Director
}

func NewUpdateResurrectionCmd(director boshdir.Director) UpdateResurrectionCmd {
	return UpdateResurrectionCmd{director: director}
}

func (c UpdateResurrectionCmd) Run(opts UpdateResurrectionOpts) error {
	return c.director.EnableResurrection(bool(opts.Args.Enabled))
}
