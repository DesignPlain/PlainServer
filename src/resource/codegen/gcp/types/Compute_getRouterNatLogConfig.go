package types

type Compute_getRouterNatLogConfig struct {
	// Specifies the desired filtering of logs on this NAT. Possible values: ["ERRORS_ONLY", "TRANSLATIONS_ONLY", "ALL"]
	Filter string `json:"filter,omitempty" yaml:"filter,omitempty"`

	// Indicates whether or not to export logs.
	Enable bool `json:"enable,omitempty" yaml:"enable,omitempty"`
}
