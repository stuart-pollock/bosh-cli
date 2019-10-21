package cmd

import (
	. "github.com/stuart-pollock/bosh-cli/cmd/opts"
	boshdir "github.com/stuart-pollock/bosh-cli/director"
	boshui "github.com/stuart-pollock/bosh-cli/ui"
)

type DeleteDiskCmd struct {
	ui       boshui.UI
	director boshdir.Director
}

func NewDeleteDiskCmd(ui boshui.UI, director boshdir.Director) DeleteDiskCmd {
	return DeleteDiskCmd{ui: ui, director: director}
}

func (c DeleteDiskCmd) Run(opts DeleteDiskOpts) error {
	err := c.ui.AskForConfirmation()
	if err != nil {
		return err
	}

	disk, err := c.director.FindOrphanDisk(opts.Args.CID)
	if err != nil {
		return err
	}

	return disk.Delete()
}
