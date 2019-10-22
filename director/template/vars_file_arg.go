package template

import (
	"github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cloudfoundry/bosh-utils/logger"
	"github.com/cloudfoundry/bosh-utils/system"
	"gopkg.in/yaml.v2"
)

type VarsFileArg struct {
	FS system.FileSystem

	Vars StaticVariables
}

func (a *VarsFileArg) UnmarshalFlag(filePath string) error {
	if a.FS == nil {
		a.FS = system.NewOsFileSystemWithStrictTempRoot(logger.NewLogger(logger.LevelNone))
	}

	if len(filePath) == 0 {
		return errors.Errorf("Expected file path to be non-empty")
	}

	bytes, err := a.FS.ReadFile(filePath)
	if err != nil {
		return errors.WrapErrorf(err, "Reading variables file '%s'", filePath)
	}

	var vars StaticVariables

	err = yaml.Unmarshal(bytes, &vars)
	if err != nil {
		return errors.WrapErrorf(err, "Deserializing variables file '%s'", filePath)
	}

	a.Vars = vars

	return nil
}
