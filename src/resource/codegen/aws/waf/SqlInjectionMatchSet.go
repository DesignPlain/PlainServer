package waf

import types "libds/aws/types"

type SqlInjectionMatchSet struct {
	// The parts of web requests that you want AWS WAF to inspect for malicious SQL code and, if you want AWS WAF to inspect a header, the name of the header.
	SqlInjectionMatchTuples []types.Waf_SqlInjectionMatchSetSqlInjectionMatchTuple `json:"sqlInjectionMatchTuples,omitempty" yaml:"sqlInjectionMatchTuples,omitempty"`

	// The name or description of the SQL Injection Match Set.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
