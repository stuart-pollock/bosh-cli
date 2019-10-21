package cmd

import (
	boshdir "github.com/stuart-pollock/bosh-cli/director"
	boshui "github.com/stuart-pollock/bosh-cli/ui"
)

type ManifestCmd struct {
	ui         boshui.UI
	deployment boshdir.Deployment
}

func NewManifestCmd(ui boshui.UI, deployment boshdir.Deployment) ManifestCmd {
	return ManifestCmd{ui: ui, deployment: deployment}
}

func (c ManifestCmd) Run() error {
	manifest, err := c.deployment.Manifest()
	if err != nil {
		return err
	}

	c.ui.PrintBlock([]byte(manifest))

	return nil
}
