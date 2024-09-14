package types

type Compute_getMachineTypesMachineTypeDeprecated struct {
	// The URL of the suggested replacement for a deprecated machine type.
	Replacement string `json:"replacement,omitempty" yaml:"replacement,omitempty"`

	// The deprecation state of this resource. This can be `ACTIVE`, `DEPRECATED`, `OBSOLETE`, or `DELETED`.
	State string `json:"state,omitempty" yaml:"state,omitempty"`
}
