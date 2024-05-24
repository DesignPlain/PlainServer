package wafregional

import types "DesignSphere_Server/src/resource/aws/types"

type SqlInjectionMatchSet struct {
	// The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
	SqlInjectionMatchTuples []types.Wafregional_SqlInjectionMatchSetSqlInjectionMatchTuple `json:"sqlInjectionMatchTuples,omitempty" yaml:"sqlInjectionMatchTuples,omitempty"`

	// The name or description of the SizeConstraintSet.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
