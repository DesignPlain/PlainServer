package types

type Lambda_EventSourceMappingFilterCriteriaFilter struct {
	// A filter pattern up to 4096 characters. See [Filter Rule Syntax](https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html#filtering-syntax).
	Pattern string `json:"pattern,omitempty" yaml:"pattern,omitempty"`
}
