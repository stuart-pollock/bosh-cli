package cmd_test

import (
	"errors"

	"github.com/cloudfoundry/bosh-utils/logger"
	fakesys "github.com/cloudfoundry/bosh-utils/system/fakes"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/stuart-pollock/bosh-cli/cmd"
	. "github.com/stuart-pollock/bosh-cli/cmd/opts"
	"github.com/stuart-pollock/bosh-cli/ui"
	fakeui "github.com/stuart-pollock/bosh-cli/ui/fakes"
	"github.com/stuart-pollock/bosh-cli/ui/table"
)

var _ = Describe("Cmd", func() {
	var (
		fakeUI *fakeui.FakeUI
		confUI *ui.ConfUI
		fs     *fakesys.FakeFileSystem
		cmd    Cmd
	)

	BeforeEach(func() {
		fakeUI = &fakeui.FakeUI{}
		log := logger.NewLogger(logger.LevelNone)
		confUI = ui.NewWrappingConfUI(fakeUI, log)

		fs = fakesys.NewFakeFileSystem()

		deps := NewBasicDeps(confUI, log)

		cmd = NewCmd(MainOpts{}, nil, deps)
	})

	Describe("Execute", func() {
		It("succeeds executing at least one command", func() {
			cmd.Opts = &InterpolateOpts{}

			err := cmd.Execute()
			Expect(err).ToNot(HaveOccurred())

			Expect(fakeUI.Blocks).To(Equal([]string{"null\n"}))
		})

		It("prints message if specified", func() {
			cmd.Opts = &MessageOpts{Message: "output"}

			err := cmd.Execute()
			Expect(err).ToNot(HaveOccurred())

			Expect(fakeUI.Blocks).To(Equal([]string{"output"}))
		})

		It("allows to enable json output", func() {
			cmd.MainOpts = MainOpts{JSONOpt: true}
			cmd.Opts = &InterpolateOpts{}

			err := cmd.Execute()
			Expect(err).ToNot(HaveOccurred())

			confUI.Flush()

			Expect(fakeUI.Blocks[0]).To(ContainSubstring(`Blocks": [`))
		})

		Describe("color", func() {
			executeCmdAndPrintTable := func() {
				err := cmd.Execute()
				Expect(err).ToNot(HaveOccurred())

				// Tables have emboldened header values
				confUI.PrintTable(table.Table{Header: []table.Header{table.NewHeader("State")}})
			}

			It("has color in the output enabled by default", func() {
				cmd.MainOpts = MainOpts{}
				cmd.Opts = &InterpolateOpts{}

				executeCmdAndPrintTable()

				// Expect that header values are bold
				Expect(fakeUI.Tables[0].HeaderFormatFunc).ToNot(BeNil())
			})

			It("allows to disable color in the output", func() {
				cmd.MainOpts = MainOpts{NoColorOpt: true}
				cmd.Opts = &InterpolateOpts{}

				executeCmdAndPrintTable()

				// Expect that header values are empty because they were not emboldened
				Expect(fakeUI.Tables[0].HeaderFormatFunc).To(BeNil())
			})
		})

		It("returns error if changing tmp root fails", func() {
			fs.ChangeTempRootErr = errors.New("fake-err")

			err := cmd.Execute()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("fake-err"))
		})

		It("returns error for unknown commands", func() {
			err := cmd.Execute()
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("Unhandled command: <nil>"))
		})
	})
})
