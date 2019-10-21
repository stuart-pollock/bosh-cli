package cmd

import boshdir "github.com/stuart-pollock/bosh-cli/director"
import . "github.com/stuart-pollock/bosh-cli/cmd/opts"

type AttachDiskCmd struct {
	deployment boshdir.Deployment
}

func NewAttachDiskCmd(deployment boshdir.Deployment) AttachDiskCmd {
	return AttachDiskCmd{
		deployment: deployment,
	}
}

func (c AttachDiskCmd) Run(opts AttachDiskOpts) error {
	return c.deployment.AttachDisk(opts.Args.Slug, opts.Args.DiskCID, opts.DiskProperties)
}
