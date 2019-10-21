package fakes

import (
	"github.com/stuart-pollock/go-patch/patch"

	boshtpl "github.com/stuart-pollock/bosh-cli/director/template"
	biinstallmanifest "github.com/stuart-pollock/bosh-cli/installation/manifest"
	birelsetmanifest "github.com/stuart-pollock/bosh-cli/release/set/manifest"
)

type FakeParser struct {
	ParsePath          string
	ReleaseSetManifest birelsetmanifest.Manifest
	ParseManifest      biinstallmanifest.Manifest
	ParseErr           error
}

func NewFakeParser() *FakeParser {
	return &FakeParser{}
}

func (p *FakeParser) Parse(path string, vars boshtpl.Variables, op patch.Op, releaseSetManifest birelsetmanifest.Manifest) (biinstallmanifest.Manifest, error) {
	p.ParsePath = path
	p.ReleaseSetManifest = releaseSetManifest
	return p.ParseManifest, p.ParseErr
}
