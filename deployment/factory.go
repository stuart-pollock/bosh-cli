package deployment

import (
	"time"

	bidisk "github.com/stuart-pollock/bosh-cli/deployment/disk"
	biinstance "github.com/stuart-pollock/bosh-cli/deployment/instance"
	bistemcell "github.com/stuart-pollock/bosh-cli/stemcell"
)

type Factory interface {
	NewDeployment(
		[]biinstance.Instance,
		[]bidisk.Disk,
		[]bistemcell.CloudStemcell,
	) Deployment
}

type factory struct {
	pingTimeout time.Duration
	pingDelay   time.Duration
}

func NewFactory(
	pingTimeout time.Duration,
	pingDelay time.Duration,
) Factory {
	return &factory{
		pingTimeout: pingTimeout,
		pingDelay:   pingDelay,
	}
}

func (f *factory) NewDeployment(
	instances []biinstance.Instance,
	disks []bidisk.Disk,
	stemcells []bistemcell.CloudStemcell,
) Deployment {
	return NewDeployment(
		instances,
		disks,
		stemcells,
		f.pingTimeout,
		f.pingDelay,
	)
}
