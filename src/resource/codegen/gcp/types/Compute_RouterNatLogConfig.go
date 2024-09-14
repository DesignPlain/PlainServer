package types

type Compute_RouterNatLogConfig struct {
	// Indicates whether or not to export logs.
	Enable bool `json:"enable,omitempty" yaml:"enable,omitempty"`

	/*
	   Specifies the desired filtering of logs on this NAT.
	   Possible values are: `ERRORS_ONLY`, `TRANSLATIONS_ONLY`, `ALL`.
	*/
	Filter string `json:"filter,omitempty" yaml:"filter,omitempty"`
}
