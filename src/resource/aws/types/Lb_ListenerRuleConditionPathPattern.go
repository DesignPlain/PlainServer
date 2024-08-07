package types

type Lb_ListenerRuleConditionPathPattern struct {
	// List of header value patterns to match. Maximum size of each pattern is 128 characters. Comparison is case insensitive. Wildcard characters supported: - (matches 0 or more characters) and ? (matches exactly 1 character). If the same header appears multiple times in the request they will be searched in order until a match is found. Only one pattern needs to match for the condition to be satisfied. To require that all of the strings are a match, create one condition block per string.
	Values []string `json:"values,omitempty" yaml:"values,omitempty"`
}
