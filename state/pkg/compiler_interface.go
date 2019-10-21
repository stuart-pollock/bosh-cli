package pkg

import (
	birelpkg "github.com/stuart-pollock/bosh-cli/release/pkg"
)

type Compiler interface {
	Compile(birelpkg.Compilable) (CompiledPackageRecord, bool, error)
}
