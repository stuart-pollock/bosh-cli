package stemcell

import (
	bicloud "github.com/stuart-pollock/bosh-cli/cloud"
	biconfig "github.com/stuart-pollock/bosh-cli/config"
)

type ManagerFactory interface {
	NewManager(bicloud.Cloud) Manager
}

type managerFactory struct {
	repo biconfig.StemcellRepo
}

func NewManagerFactory(repo biconfig.StemcellRepo) ManagerFactory {
	return &managerFactory{
		repo: repo,
	}
}

func (f *managerFactory) NewManager(cloud bicloud.Cloud) Manager {
	return NewManager(f.repo, cloud)
}
