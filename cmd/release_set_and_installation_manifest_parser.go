package cmd

import (
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	"github.com/stuart-pollock/go-patch/patch"

	boshtpl "github.com/stuart-pollock/bosh-cli/director/template"
	biinstallmanifest "github.com/stuart-pollock/bosh-cli/installation/manifest"
	birelsetmanifest "github.com/stuart-pollock/bosh-cli/release/set/manifest"
)

type ReleaseSetAndInstallationManifestParser struct {
	ReleaseSetParser   birelsetmanifest.Parser
	InstallationParser biinstallmanifest.Parser
}

func (y ReleaseSetAndInstallationManifestParser) ReleaseSetAndInstallationManifest(deploymentManifestPath string, vars boshtpl.Variables, op patch.Op) (birelsetmanifest.Manifest, biinstallmanifest.Manifest, error) {
	releaseSetManifest, err := y.ReleaseSetParser.Parse(deploymentManifestPath, vars, op)
	if err != nil {
		return birelsetmanifest.Manifest{}, biinstallmanifest.Manifest{}, bosherr.WrapErrorf(err, "Parsing release set manifest '%s'", deploymentManifestPath)
	}

	installationManifest, err := y.InstallationParser.Parse(deploymentManifestPath, vars, op, releaseSetManifest)
	if err != nil {
		return birelsetmanifest.Manifest{}, biinstallmanifest.Manifest{}, bosherr.WrapErrorf(err, "Parsing installation manifest '%s'", deploymentManifestPath)
	}

	return releaseSetManifest, installationManifest, nil
}
