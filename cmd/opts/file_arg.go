package opts

import (
	"github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cloudfoundry/bosh-utils/logger"
	"github.com/cloudfoundry/bosh-utils/system"
)

type FileArg struct {
	ExpandedPath string
	FS           system.FileSystem
}

func (a *FileArg) UnmarshalFlag(data string) error {
	if a.FS == nil {
		a.FS = system.NewOsFileSystemWithStrictTempRoot(logger.NewLogger(logger.LevelNone))
	}

	expandedPath, err := a.FS.ExpandPath(data)
	if err != nil {
		return errors.WrapErrorf(err, "Checking file path")
	}
	a.ExpandedPath = expandedPath

	if a.FS.FileExists(expandedPath) {
		stat, err := a.FS.Stat(expandedPath)
		if err != nil {
			return errors.WrapErrorf(err, "Checking file path")
		}

		if stat.IsDir() {
			return errors.Errorf("Path must not be directory")
		}
	}

	return nil
}
