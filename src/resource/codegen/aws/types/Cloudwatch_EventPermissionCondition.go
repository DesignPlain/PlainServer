package types

type Cloudwatch_EventPermissionCondition struct {
	// Key for the condition. Valid values: `aws:PrincipalOrgID`.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Type of condition. Value values: `StringEquals`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Value for the key.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
