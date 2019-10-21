package cmd

import (
	cmdconf "github.com/stuart-pollock/bosh-cli/cmd/config"
	boshdir "github.com/stuart-pollock/bosh-cli/director"
	boshuaa "github.com/stuart-pollock/bosh-cli/uaa"
)

//go:generate counterfeiter . SessionContext

type SessionContext interface {
	Environment() string
	CACert() string
	Config() cmdconf.Config
	Credentials() cmdconf.Creds

	Deployment() string
}

//go:generate counterfeiter . Session

type Session interface {
	Environment() string
	Credentials() cmdconf.Creds

	UAA() (boshuaa.UAA, error)

	Director() (boshdir.Director, error)
	AnonymousDirector() (boshdir.Director, error)

	Deployment() (boshdir.Deployment, error)
}
