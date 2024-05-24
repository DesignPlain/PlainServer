package wafregional

import types "DesignSphere_Server/src/resource/aws/types"

type RegexMatchSet struct {
	// The name or description of the Regex Match Set.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The regular expression pattern that you want AWS WAF to search for in web requests, the location in requests that you want AWS WAF to search, and other settings. See below.
	RegexMatchTuples []types.Wafregional_RegexMatchSetRegexMatchTuple `json:"regexMatchTuples,omitempty" yaml:"regexMatchTuples,omitempty"`
}
