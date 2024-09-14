package waf

type RegexPatternSet struct {
	// The name or description of the Regex Pattern Set.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A list of regular expression (regex) patterns that you want AWS WAF to search for, such as `B[a@]dB[o0]t`.
	RegexPatternStrings []string `json:"regexPatternStrings,omitempty" yaml:"regexPatternStrings,omitempty"`
}
