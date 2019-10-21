package cmd

import (
	. "github.com/stuart-pollock/bosh-cli/cmd/opts"
	boshdir "github.com/stuart-pollock/bosh-cli/director"
	boshui "github.com/stuart-pollock/bosh-cli/ui"
)

type CleanUpCmd struct {
	ui       boshui.UI
	director boshdir.Director
}

func NewCleanUpCmd(ui boshui.UI, director boshdir.Director) CleanUpCmd {
	return CleanUpCmd{ui: ui, director: director}
}

func (c CleanUpCmd) Run(opts CleanUpOpts) error {
	err := c.ui.AskForConfirmation()
	if err != nil {
		return err
	}

	return c.director.CleanUp(opts.All)
}
