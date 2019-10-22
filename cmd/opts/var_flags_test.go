package opts_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/stuart-pollock/bosh-cli/cmd/opts"
	. "github.com/stuart-pollock/bosh-cli/director/template"
)

var _ = Describe("VarFlags", func() {
	Describe("AsVariables", func() {
		It("prefers kvs, to var files, to vars files, to env variables (without fs store)", func() {
			flags := VarFlags{
				VarKVs: []VarKV{
					{Name: "kv", Value: "kv"},
					{Name: "kv_precedence", Value: "kv1"},
					{Name: "kv_precedence", Value: "kv2"},
					{Name: "kv_env_precedence", Value: "kv"},
					{Name: "kv_file_precedence", Value: "kv"},
					{Name: "kv_file_env_precedence", Value: "kv"},
				},
				VarFiles: []VarFileArg{
					{Vars: StaticVariables{
						"var_file":                 "var_file",
						"var_file_precedence":      "var_file",
						"var_file_file_precedence": "var_file",
					}},
					{Vars: StaticVariables{
						"var_file_precedence": "var_file2",
					}},
				},
				VarsFiles: []VarsFileArg{
					{Vars: StaticVariables{
						"file":                     "file",
						"file_precedence":          "file",
						"var_file_file_precedence": "file",
					}},
					{Vars: StaticVariables{
						"file_env_precedence":    "file2",
						"kv_file_env_precedence": "file2",
						"kv_file_precedence":     "file2",
						"file2":                  "file2",
						"file_precedence":        "file2",
					}},
				},
				VarsEnvs: []VarsEnvArg{
					{Vars: StaticVariables{
						"env":            "env",
						"env_precedence": "env",
					}},
					{Vars: StaticVariables{
						"kv_env_precedence":      "env2",
						"file_env_precedence":    "env2",
						"kv_file_env_precedence": "env2",
						"env2":                   "env2",
						"env_precedence":         "env2",
					}},
				},
			}

			vars := flags.AsVariables()

			expectedVals := map[string]string{
				"kv":                       "kv",
				"kv_precedence":            "kv2",
				"var_file":                 "var_file",
				"var_file_precedence":      "var_file2",
				"var_file_file_precedence": "var_file",
				"file":                     "file",
				"file_precedence":          "file2",
				"kv_file_precedence":       "kv",
				"file2":                    "file2",
				"env2":                     "env2",
				"env":                      "env",
				"env_precedence":           "env2",
				"kv_file_env_precedence":   "kv",
				"file_env_precedence":      "file2",
				"kv_env_precedence":        "kv",
			}

			for key, expectedVal := range expectedVals {
				val, found, err := vars.Get(VariableDefinition{Name: key})
				Expect(val).To(Equal(expectedVal), fmt.Sprintf("Expecting key '%s' value to match", key))
				Expect(found).To(BeTrue())
				Expect(err).ToNot(HaveOccurred())
			}
		})
	})
})
