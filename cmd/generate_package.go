package cmd

import (
	. "github.com/stuart-pollock/bosh-cli/cmd/opts"
	boshreldir "github.com/stuart-pollock/bosh-cli/releasedir"
)

type GeneratePackageCmd struct {
	releaseDir boshreldir.ReleaseDir
}

func NewGeneratePackageCmd(releaseDir boshreldir.ReleaseDir) GeneratePackageCmd {
	return GeneratePackageCmd{releaseDir: releaseDir}
}

func (c GeneratePackageCmd) Run(opts GeneratePackageOpts) error {
	return c.releaseDir.GeneratePackage(opts.Args.Name)
}
