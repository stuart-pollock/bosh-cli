package cmd

import (
	. "github.com/stuart-pollock/bosh-cli/cmd/opts"
	boshreldir "github.com/stuart-pollock/bosh-cli/releasedir"
)

type InitReleaseCmd struct {
	releaseDir boshreldir.ReleaseDir
}

func NewInitReleaseCmd(releaseDir boshreldir.ReleaseDir) InitReleaseCmd {
	return InitReleaseCmd{releaseDir: releaseDir}
}

func (c InitReleaseCmd) Run(opts InitReleaseOpts) error {
	return c.releaseDir.Init(opts.Git)
}
