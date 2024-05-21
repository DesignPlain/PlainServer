package types

type Monitoring_UptimeCheckConfigContentMatcherJsonPathMatcher struct {
	/*
	   Options to perform JSONPath content matching.
	   Default value is `EXACT_MATCH`.
	   Possible values are: `EXACT_MATCH`, `REGEX_MATCH`.
	*/
	JsonMatcher string `json:"jsonMatcher,omitempty" yaml:"jsonMatcher,omitempty"`

	// JSONPath within the response output pointing to the expected `ContentMatcher::content` to match against.
	JsonPath string `json:"jsonPath,omitempty" yaml:"jsonPath,omitempty"`
}
