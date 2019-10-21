package cmd

import boshdir "github.com/stuart-pollock/bosh-cli/director"
import . "github.com/stuart-pollock/bosh-cli/cmd/opts"

type UnignoreCmd struct {
	deployment boshdir.Deployment
}

func NewUnignoreCmd(deployment boshdir.Deployment) UnignoreCmd {
	return UnignoreCmd{deployment: deployment}
}

func (cmd UnignoreCmd) Run(opts UnignoreOpts) error {
	return cmd.deployment.Ignore(opts.Args.Slug, false)
}
