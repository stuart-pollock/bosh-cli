package cmd

import (
	. "github.com/stuart-pollock/bosh-cli/cmd/opts"
	boshdir "github.com/stuart-pollock/bosh-cli/director"
	boshui "github.com/stuart-pollock/bosh-cli/ui"
)

type EventCmd struct {
	ui       boshui.UI
	director boshdir.Director
}

func NewEventCmd(ui boshui.UI, director boshdir.Director) EventCmd {
	return EventCmd{ui: ui, director: director}
}

func (c EventCmd) Run(opts EventOpts) error {
	event, err := c.director.Event(opts.Args.ID)
	if err != nil {
		return err
	}

	EventTable{event, c.ui}.Print()

	return nil
}
