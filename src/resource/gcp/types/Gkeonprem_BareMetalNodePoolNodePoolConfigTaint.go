package types

type Gkeonprem_BareMetalNodePoolNodePoolConfigTaint struct {
	/*
	   Specifies the nodes operating system (default: LINUX).
	   Possible values are: `EFFECT_UNSPECIFIED`, `PREFER_NO_SCHEDULE`, `NO_EXECUTE`.

	   - - -
	*/
	Effect string `json:"effect,omitempty" yaml:"effect,omitempty"`

	// Key associated with the effect.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Value associated with the effect.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
