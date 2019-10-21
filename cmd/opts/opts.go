package opts

import (
	"github.com/stuart-pollock/go-patch/patch"
)

type BoshOpts struct {
	// -----> Global options

	VersionOpt func() error `long:"version" short:"v" description:"Show CLI version"`

	ConfigPathOpt string `long:"config" description:"Config file path" env:"BOSH_CONFIG" default:"~/.bosh/config"`

	EnvironmentOpt string    `long:"environment" short:"e" description:"Director environment name or URL" env:"BOSH_ENVIRONMENT"`
	CACertOpt      CACertArg `long:"ca-cert"               description:"Director CA certificate path or value" env:"BOSH_CA_CERT"`
	Sha2           bool      `long:"sha2"                  description:"Use SHA256 checksums" env:"BOSH_SHA2"`
	Parallel       int       `long:"parallel" description:"The max number of parallel operations" default:"5"`

	// Specify client credentials
	ClientOpt       string `long:"client"        description:"Override username or UAA client"        env:"BOSH_CLIENT"`
	ClientSecretOpt string `long:"client-secret" description:"Override password or UAA client secret" env:"BOSH_CLIENT_SECRET"`

	DeploymentOpt string `long:"deployment" short:"d" description:"Deployment name" env:"BOSH_DEPLOYMENT"`

	// Output formatting
	ColumnOpt         []ColumnOpt `long:"column"                    description:"Filter to show only given column(s)"`
	JSONOpt           bool        `long:"json"                      description:"Output as JSON"`
	TTYOpt            bool        `long:"tty"                       description:"Force TTY-like output"`
	NoColorOpt        bool        `long:"no-color"                  description:"Toggle colorized output"`
	NonInteractiveOpt bool        `long:"non-interactive" short:"n" description:"Don't ask for user input" env:"BOSH_NON_INTERACTIVE"`

	Help HelpOpts `command:"help" description:"Show this help message"`

	Interpolate InterpolateOpts `command:"interpolate" alias:"int" description:"Interpolates variables into a manifest"`

	Variables VariablesOpts `command:"variables" alias:"vars" description:"List variables"`
}

type HelpOpts struct {
	cmd
}

type InterpolateOpts struct {
	Args InterpolateArgs `positional-args:"true" required:"true"`

	VarFlags
	OpsFlags

	Path            patch.Pointer `long:"path" value-name:"OP-PATH" description:"Extract value out of template (e.g.: /private_key)"`
	VarErrors       bool          `long:"var-errs"                  description:"Expect all variables to be found, otherwise error"`
	VarErrorsUnused bool          `long:"var-errs-unused"           description:"Expect all variables to be used, otherwise error"`

	cmd
}

type InterpolateArgs struct {
	Manifest FileBytesArg `positional-arg-name:"PATH" description:"Path to a template that will be interpolated"`
}

// MessageOpts is used for version and help flags
type MessageOpts struct {
	Message string
}

type VariablesOpts struct {
	Deployment string
	cmd
}

type cmd struct{}

// Execute is necessary for each command to be goflags.Commander
func (c cmd) Execute(_ []string) error {
	panic("Unreachable")
}
