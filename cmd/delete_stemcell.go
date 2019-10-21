package cmd

import (
	. "github.com/stuart-pollock/bosh-cli/cmd/opts"
	boshdir "github.com/stuart-pollock/bosh-cli/director"
	boshui "github.com/stuart-pollock/bosh-cli/ui"
)

type DeleteStemcellCmd struct {
	ui       boshui.UI
	director boshdir.Director
}

func NewDeleteStemcellCmd(ui boshui.UI, director boshdir.Director) DeleteStemcellCmd {
	return DeleteStemcellCmd{ui: ui, director: director}
}

func (c DeleteStemcellCmd) Run(opts DeleteStemcellOpts) error {
	err := c.ui.AskForConfirmation()
	if err != nil {
		return err
	}

	stemcell, err := c.director.FindStemcell(opts.Args.Slug)
	if err != nil {
		return err
	}

	return stemcell.Delete(opts.Force)
}
