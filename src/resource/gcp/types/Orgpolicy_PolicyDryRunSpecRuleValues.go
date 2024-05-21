package types

type Orgpolicy_PolicyDryRunSpecRuleValues struct {
	// List of values denied at this resource.
	DeniedValues []string `json:"deniedValues,omitempty" yaml:"deniedValues,omitempty"`

	// List of values allowed at this resource.
	AllowedValues []string `json:"allowedValues,omitempty" yaml:"allowedValues,omitempty"`
}
