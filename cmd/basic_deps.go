package cmd

import (
	"code.cloudfoundry.org/clock"
	"github.com/cloudfoundry/bosh-utils/logger"
	"github.com/cloudfoundry/bosh-utils/system"

	mainUI "github.com/stuart-pollock/bosh-cli/ui"
)

type BasicDeps struct {
	UI     *mainUI.ConfUI
	Logger logger.Logger

	CmdRunner system.CmdRunner

	Time clock.Clock
}

func NewBasicDeps(ui *mainUI.ConfUI, log logger.Logger) BasicDeps {
	cmdRunner := system.NewExecCmdRunner(log)

	return BasicDeps{
		UI:     ui,
		Logger: log,

		CmdRunner: cmdRunner,
		Time:      clock.NewClock(),
	}
}
