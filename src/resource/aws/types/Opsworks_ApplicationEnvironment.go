package types

type Opsworks_ApplicationEnvironment struct {
	// Variable name.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Set visibility of the variable value to `true` or `false`.
	Secure bool `json:"secure,omitempty" yaml:"secure,omitempty"`

	// Variable value.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
