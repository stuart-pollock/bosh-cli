package opts_test

import (
	"reflect"
	"regexp"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/stuart-pollock/bosh-cli/cmd/opts"
)

var (
	dupSpaces = regexp.MustCompile("\\s{2,}")
)

func getStructTagForName(field string, opts interface{}) string {
	st, _ := reflect.TypeOf(opts).Elem().FieldByName(field)
	return dupSpaces.ReplaceAllString(string(st.Tag), " ")
}

func getStructTagForType(field string, opts interface{}) string {
	st, _ := reflect.TypeOf(opts).Elem().FieldByName(field)
	return dupSpaces.ReplaceAllString(string(st.Tag), " ")
}

var _ = Describe("Opts", func() {
	Describe("MainOpts", func() {
		var opts *MainOpts

		BeforeEach(func() {
			opts = &MainOpts{}
		})

		Describe("VersionOpt", func() {
			It("contains desired values", func() {
				Expect(getStructTagForName("VersionOpt", opts)).To(Equal(
					`long:"version" short:"v" description:"Show CLI version"`,
				))
			})
		})

		Describe("Interpolate", func() {
			It("contains desired values", func() {
				Expect(getStructTagForName("Interpolate", opts)).To(Equal(
					`command:"interpolate" alias:"int" description:"Interpolates variables into a manifest"`,
				))
			})
		})
	})
})
