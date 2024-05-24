package wafregional

import types "DesignSphere_Server/src/resource/aws/types"

type XssMatchSet struct {
	// The name of the set
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The parts of web requests that you want to inspect for cross-site scripting attacks.
	XssMatchTuples []types.Wafregional_XssMatchSetXssMatchTuple `json:"xssMatchTuples,omitempty" yaml:"xssMatchTuples,omitempty"`
}
