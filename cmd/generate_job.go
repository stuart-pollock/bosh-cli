package cmd

import (
	. "github.com/stuart-pollock/bosh-cli/cmd/opts"
	boshreldir "github.com/stuart-pollock/bosh-cli/releasedir"
)

type GenerateJobCmd struct {
	releaseDir boshreldir.ReleaseDir
}

func NewGenerateJobCmd(releaseDir boshreldir.ReleaseDir) GenerateJobCmd {
	return GenerateJobCmd{releaseDir: releaseDir}
}

func (c GenerateJobCmd) Run(opts GenerateJobOpts) error {
	return c.releaseDir.GenerateJob(opts.Args.Name)
}
