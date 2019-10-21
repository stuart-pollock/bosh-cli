package cmd

import (
	. "github.com/stuart-pollock/bosh-cli/cmd/opts"
	boshdir "github.com/stuart-pollock/bosh-cli/director"
	biui "github.com/stuart-pollock/bosh-cli/ui"
	boshtbl "github.com/stuart-pollock/bosh-cli/ui/table"
)

type InspectStemcellTarballCmd struct {
	stemcellArchiveFactory func(string) boshdir.StemcellArchive
	ui                     biui.UI
}

func NewInspectStemcellTarballCmd(
	stemcellArchiveFactory func(string) boshdir.StemcellArchive,
	ui biui.UI,
) InspectStemcellTarballCmd {
	return InspectStemcellTarballCmd{
		stemcellArchiveFactory: stemcellArchiveFactory,
		ui:                     ui,
	}
}

func (c InspectStemcellTarballCmd) Run(opts InspectStemcellTarballOpts) error {
	archive := c.stemcellArchiveFactory(opts.Args.PathToStemcell)
	metadata, err := archive.Info()
	if err != nil {
		return err
	}

	infrastructure := metadata.CloudProperties["infrastructure"]
	if infrastructure == nil {
		infrastructure = "unknown"
	}

	metadataTable := boshtbl.Table{
		Content: "stemcell-metadata",
		Header: []boshtbl.Header{
			boshtbl.NewHeader("Name"),
			boshtbl.NewHeader("OS"),
			boshtbl.NewHeader("Version"),
			boshtbl.NewHeader("Infrastructure"),
		},
		SortBy: []boshtbl.ColumnSort{{Column: 0, Asc: true}},
		Rows: [][]boshtbl.Value{
			{
				boshtbl.NewValueString(metadata.Name),
				boshtbl.NewValueString(metadata.OS),
				boshtbl.NewValueString(metadata.Version),
				boshtbl.NewValueString(infrastructure.(string)),
			},
		},
	}

	c.ui.PrintTable(metadataTable)
	return nil
}
