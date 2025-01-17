package cmd_test

import (
	"io/ioutil"

	"github.com/cloudfoundry/bosh-utils/logger"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	. "github.com/stuart-pollock/bosh-cli/cmd"
	. "github.com/stuart-pollock/bosh-cli/cmd/opts"
	"github.com/stuart-pollock/bosh-cli/ui"
)

// This placeholder is used for replacing arguments in the table test with the
// temporary file created in the BeforeEach
const filePlaceholder = "replace-me"

var _ = Describe("Factory", func() {
	var (
		factory Factory
		tmpFile string
	)

	BeforeEach(func() {
		log := logger.NewLogger(logger.LevelNone)

		f, err := ioutil.TempFile("", "file")
		Expect(err).NotTo(HaveOccurred())

		tmpFile = f.Name()

		myUi := ui.NewConfUI(log)
		defer myUi.Flush()

		deps := NewBasicDeps(myUi, log)

		factory = NewFactory(deps)
	})

	Context("extra args and flags", func() {
		DescribeTable("extra args and flags", func(cmd string, args []string) {
			for i, arg := range args {
				if arg == filePlaceholder {
					args[i] = tmpFile
				}
			}
			cmdWithArgs := append([]string{cmd}, args...)
			cmdWithArgs = append(cmdWithArgs, "extra", "args")

			_, err := factory.New(cmdWithArgs)
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("does not support extra arguments: extra, args"))
		},
			Entry("help", "help", []string{}),
			Entry("interpolate", "interpolate", []string{filePlaceholder}),
		)

		It("catches unknown commands and lists available commands", func() {
			_, err := factory.New([]string{"unknown-cmd"})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("Unknown command `unknown-cmd'. Please specify one command of: add-blob"))
		})

		It("catches unknown global flags", func() {
			_, err := factory.New([]string{"--unknown-flag"})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("unknown flag `unknown-flag'"))
		})

		It("catches unknown command flags", func() {
			_, err := factory.New([]string{"ssh", "--unknown-flag"})
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(Equal("unknown flag `unknown-flag'"))
		})
	})

	Describe("help command", func() {
		It("has a help command", func() {
			cmd, err := factory.New([]string{"help"})
			Expect(err).ToNot(HaveOccurred())

			opts := cmd.Opts.(*MessageOpts)
			Expect(opts.Message).To(ContainSubstring("Usage:"))
			Expect(opts.Message).To(ContainSubstring("Application Options:"))
			Expect(opts.Message).To(ContainSubstring("Available commands:"))
		})
	})

	Describe("help options", func() {
		It("has a help flag", func() {
			cmd, err := factory.New([]string{"--help"})
			Expect(err).ToNot(HaveOccurred())

			opts := cmd.Opts.(*MessageOpts)
			Expect(opts.Message).To(ContainSubstring("Usage:"))
			Expect(opts.Message).To(ContainSubstring(
				"SSH into instance(s)                               https://bosh.io/docs/cli-v2#ssh"))
			Expect(opts.Message).To(ContainSubstring("Application Options:"))
			Expect(opts.Message).To(ContainSubstring("Available commands:"))
		})

		It("has a command help flag", func() {
			cmd, err := factory.New([]string{"ssh", "--help"})
			Expect(err).ToNot(HaveOccurred())

			opts := cmd.Opts.(*MessageOpts)
			Expect(opts.Message).To(ContainSubstring("Usage:"))
			Expect(opts.Message).To(ContainSubstring("SSH into instance(s)\n\nhttps://bosh.io/docs/cli-v2#ssh"))
			Expect(opts.Message).To(ContainSubstring("Application Options:"))
			Expect(opts.Message).To(ContainSubstring("[ssh command options]"))
		})
	})

	Describe("version option", func() {
		It("has a version flag", func() {
			cmd, err := factory.New([]string{"--version"})
			Expect(err).ToNot(HaveOccurred())

			opts := cmd.Opts.(*MessageOpts)
			Expect(opts.Message).To(Equal("version [DEV BUILD]\n"))
		})
	})

	Describe("global options", func() {
		clearNonGlobalOpts := func(opts MainOpts) MainOpts {
			opts.VersionOpt = nil // can't compare functions
			opts.Interpolate = InterpolateOpts{}
			return opts
		}

		It("can set variety of options", func() {
			opts := []string{
				"--config", "config",
				"--json",
			}

			cmd, err := factory.New(opts)
			Expect(err).ToNot(HaveOccurred())

			Expect(clearNonGlobalOpts(cmd.MainOpts)).To(Equal(MainOpts{
				JSONOpt: true,
			}))
		})

		It("errors when --user is set", func() {
			opts := []string{
				"--user", "foo",
				"--json",
				"--tty",
			}

			_, err := factory.New(opts)
			Expect(err).To(HaveOccurred())
		})
	})
})
