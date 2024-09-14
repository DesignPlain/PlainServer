package types

type Route53recoverycontrol_SafetyRuleRuleConfig struct {
	// Logical negation of the rule.
	Inverted bool `json:"inverted,omitempty" yaml:"inverted,omitempty"`

	// Number of controls that must be set when you specify an `ATLEAST` type rule.
	Threshold int `json:"threshold,omitempty" yaml:"threshold,omitempty"`

	// Rule type. Valid values are `ATLEAST`, `AND`, and `OR`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
