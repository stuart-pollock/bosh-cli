package cmd

import (
	. "github.com/stuart-pollock/bosh-cli/cmd/opts"
	boshdir "github.com/stuart-pollock/bosh-cli/director"
	boshui "github.com/stuart-pollock/bosh-cli/ui"
)

type OrphanDiskCmd struct {
	ui       boshui.UI
	director boshdir.Director
}

func NewOrphanDiskCmd(ui boshui.UI, director boshdir.Director) OrphanDiskCmd {
	return OrphanDiskCmd{ui: ui, director: director}
}

func (c OrphanDiskCmd) Run(opts OrphanDiskOpts) error {
	err := c.ui.AskForConfirmation()
	if err != nil {
		return err
	}

	return c.director.OrphanDisk(opts.Args.CID)
}
