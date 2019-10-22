package template

import (
	"strings"

	"github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cloudfoundry/bosh-utils/logger"
	"github.com/cloudfoundry/bosh-utils/system"
)

type VarFileArg struct {
	FS system.FileSystem

	Vars StaticVariables
}

func (a *VarFileArg) UnmarshalFlag(data string) error {
	if a.FS == nil {
		a.FS = system.NewOsFileSystemWithStrictTempRoot(logger.NewLogger(logger.LevelNone))
	}

	pieces := strings.SplitN(data, "=", 2)
	if len(pieces) != 2 {
		return errors.Errorf("Expected var '%s' to be in format 'name=path'", data)
	}

	if len(pieces[0]) == 0 {
		return errors.Errorf("Expected var '%s' to specify non-empty name", data)
	}

	if len(pieces[1]) == 0 {
		return errors.Errorf("Expected var '%s' to specify non-empty path", data)
	}

	absPath, err := a.FS.ExpandPath(pieces[1])
	if err != nil {
		return errors.WrapErrorf(err, "Getting absolute path '%s'", pieces[1])
	}

	bytes, err := a.FS.ReadFile(absPath)
	if err != nil {
		return errors.WrapErrorf(err, "Reading variable from file '%s'", absPath)
	}

	(*a).Vars = StaticVariables{pieces[0]: string(bytes)}

	return nil
}
