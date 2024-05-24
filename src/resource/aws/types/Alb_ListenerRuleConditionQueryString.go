package types

type Alb_ListenerRuleConditionQueryString struct {
	// Query string key pattern to match.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Query string value pattern to match.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
