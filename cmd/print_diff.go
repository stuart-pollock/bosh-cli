package cmd

import (
	"fmt"

	"github.com/stuart-pollock/bosh-cli/ui"
)

type Diff struct {
	lines [][]interface{}
}

func NewDiff(lines [][]interface{}) Diff {
	return Diff{
		lines: lines,
	}
}

func (d Diff) Print(myUi ui.UI) {
	for _, line := range d.lines {
		lineMod, _ := line[1].(string)

		if lineMod == "added" {
			myUi.BeginLinef("+ %s\n", line[0])
		} else if lineMod == "removed" {
			myUi.BeginLinef("- %s\n", line[0])
		} else {
			myUi.BeginLinef("  %s\n", line[0])
		}
	}
}

func (d Diff) String() string {
	var result string
	for _, line := range d.lines {
		lineMod, _ := line[1].(string)

		if lineMod == "added" {
			result += fmt.Sprintf("+ %s\n", line[0])
		} else if lineMod == "removed" {
			result += fmt.Sprintf("- %s\n", line[0])
		} else {
			result += fmt.Sprintf("  %s\n", line[0])
		}
	}
	return result
}
