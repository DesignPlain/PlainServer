package waf

import types "libds/aws/types"

type XssMatchSet struct {
	// The name or description of the SizeConstraintSet.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The parts of web requests that you want to inspect for cross-site scripting attacks.
	XssMatchTuples []types.Waf_XssMatchSetXssMatchTuple `json:"xssMatchTuples,omitempty" yaml:"xssMatchTuples,omitempty"`
}
