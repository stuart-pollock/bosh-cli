package opts

import (
	"github.com/stuart-pollock/bosh-cli/director/template"
)

// Shared
type VarFlags struct {
	VarKVs    []template.VarKV       `long:"var"        short:"v" value-name:"VAR=VALUE" description:"Set variable"`
	VarFiles  []template.VarFileArg  `long:"var-file"             value-name:"VAR=PATH"  description:"Set variable to file contents"`
	VarsFiles []template.VarsFileArg `long:"vars-file"  short:"l" value-name:"PATH"      description:"Load variables from a YAML file"`
	VarsEnvs  []template.VarsEnvArg  `long:"vars-env"             value-name:"PREFIX"    description:"Load variables from environment variables (e.g.: 'MY' to load MY_var=value)"`
}

func (f VarFlags) AsVariables() template.Variables {
	var firstToUse []template.Variables

	staticVars := template.StaticVariables{}

	for i, _ := range f.VarsEnvs {
		for k, v := range f.VarsEnvs[i].Vars {
			staticVars[k] = v
		}
	}

	for i, _ := range f.VarsFiles {
		for k, v := range f.VarsFiles[i].Vars {
			staticVars[k] = v
		}
	}

	for i, _ := range f.VarFiles {
		for k, v := range f.VarFiles[i].Vars {
			staticVars[k] = v
		}
	}

	for _, kv := range f.VarKVs {
		staticVars[kv.Name] = kv.Value
	}

	firstToUse = append(firstToUse, staticVars)

	vars := template.NewMultiVars(firstToUse)

	return vars
}
