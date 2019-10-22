package opts

import (
	"io/ioutil"
	"os"

	"github.com/cloudfoundry/bosh-utils/errors"
	"github.com/cloudfoundry/bosh-utils/logger"
	"github.com/cloudfoundry/bosh-utils/system"
)

type FileBytesArg struct {
	FS system.FileSystem

	Bytes []byte
}

func (a *FileBytesArg) UnmarshalFlag(data string) error {
	if a.FS == nil {
		a.FS = system.NewOsFileSystemWithStrictTempRoot(logger.NewLogger(logger.LevelNone))
	}

	if len(data) == 0 {
		return errors.Errorf("Expected file path to be non-empty")
	}

	if data == "-" {
		bs, err := ioutil.ReadAll(os.Stdin)
		if err != nil {
			return errors.WrapErrorf(err, "Reading from stdin")
		}

		(*a).Bytes = bs

		return nil
	}

	absPath, err := a.FS.ExpandPath(data)
	if err != nil {
		return errors.WrapErrorf(err, "Getting absolute path '%s'", data)
	}

	bytes, err := a.FS.ReadFile(absPath)
	if err != nil {
		return err
	}

	(*a).Bytes = bytes

	return nil
}
