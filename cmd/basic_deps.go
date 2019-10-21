package cmd

import (
	"code.cloudfoundry.org/clock"
	boshlog "github.com/cloudfoundry/bosh-utils/logger"
	boshsys "github.com/cloudfoundry/bosh-utils/system"

	boshui "github.com/stuart-pollock/bosh-cli/ui"
)

type BasicDeps struct {
	UI     *boshui.ConfUI
	Logger boshlog.Logger

	CmdRunner boshsys.CmdRunner

	Time clock.Clock
}

func NewBasicDeps(ui *boshui.ConfUI, logger boshlog.Logger) BasicDeps {
	return NewBasicDepsWithFS(ui, boshsys.NewOsFileSystemWithStrictTempRoot(logger), logger)
}

func NewBasicDepsWithFS(ui *boshui.ConfUI, fs boshsys.FileSystem, logger boshlog.Logger) BasicDeps {
	cmdRunner := boshsys.NewExecCmdRunner(logger)

	return BasicDeps{
		UI:     ui,
		Logger: logger,

		CmdRunner: cmdRunner,
		Time:      clock.NewClock(),
	}
}
