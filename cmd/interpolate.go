package cmd

import (
	"github.com/stuart-pollock/go-patch/patch"

	. "github.com/stuart-pollock/bosh-cli/cmd/opts"
	"github.com/stuart-pollock/bosh-cli/director/template"
	"github.com/stuart-pollock/bosh-cli/ui"
)

type InterpolateCmd struct {
	ui ui.UI
}

func NewInterpolateCmd(myUi ui.UI) InterpolateCmd {
	return InterpolateCmd{ui: myUi}
}

func (c InterpolateCmd) Run(opts InterpolateOpts) error {
	tpl := template.NewTemplate(opts.Args.Manifest.Bytes)

	vars := opts.VarFlags.AsVariables()
	op := opts.OpsFlags.AsOp()
	evalOpts := template.EvaluateOpts{
		ExpectAllKeys:     opts.VarErrors,
		ExpectAllVarsUsed: opts.VarErrorsUnused,
	}

	if opts.Path.IsSet() {
		evalOpts.PostVarSubstitutionOp = patch.FindOp{Path: opts.Path}

		// Printing YAML indented multiline strings (eg SSH key) is not useful
		evalOpts.UnescapedMultiline = true
	}

	bytes, err := tpl.Evaluate(vars, op, evalOpts)
	if err != nil {
		return err
	}

	c.ui.PrintBlock(bytes)

	return nil
}
