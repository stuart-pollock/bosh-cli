package opts

import (
	"github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cloudfoundry/bosh-utils/logger"
	"github.com/cloudfoundry/bosh-utils/system"
	"github.com/stuart-pollock/go-patch/patch"
	"gopkg.in/yaml.v2"
)

type OpsFileArg struct {
	FS system.FileSystem

	Ops patch.Ops
}

func (a *OpsFileArg) UnmarshalFlag(filePath string) error {
	if a.FS == nil {
		a.FS = system.NewOsFileSystemWithStrictTempRoot(logger.NewLogger(logger.LevelNone))
	}

	if len(filePath) == 0 {
		return errors.Errorf("Expected file path to be non-empty")
	}

	bytes, err := a.FS.ReadFile(filePath)
	if err != nil {
		return errors.WrapErrorf(err, "Reading ops file '%s'", filePath)
	}

	var opDefs []patch.OpDefinition

	err = yaml.Unmarshal(bytes, &opDefs)
	if err != nil {
		return errors.WrapErrorf(err, "Deserializing ops file '%s'", filePath)
	}

	ops, err := patch.NewOpsFromDefinitions(opDefs)
	if err != nil {
		return errors.WrapErrorf(err, "Building ops")
	}

	(*a).Ops = ops

	return nil
}
