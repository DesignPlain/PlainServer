package types

type Codepipeline_WebhookFilter struct {
	// The [JSON path](https://github.com/json-path/JsonPath) to filter on.
	JsonPath string `json:"jsonPath,omitempty" yaml:"jsonPath,omitempty"`

	// The value to match on (e.g., `refs/heads/{Branch}`). See [AWS docs](https://docs.aws.amazon.com/codepipeline/latest/APIReference/API_WebhookFilterRule.html) for details.
	MatchEquals string `json:"matchEquals,omitempty" yaml:"matchEquals,omitempty"`
}
