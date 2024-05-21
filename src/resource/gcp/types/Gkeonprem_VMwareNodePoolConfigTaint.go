package types

type Gkeonprem_VMwareNodePoolConfigTaint struct {
	// Key associated with the effect.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// Value associated with the effect.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	/*
	   Available taint effects.
	   Possible values are: `EFFECT_UNSPECIFIED`, `NO_SCHEDULE`, `PREFER_NO_SCHEDULE`, `NO_EXECUTE`.
	*/
	Effect string `json:"effect,omitempty" yaml:"effect,omitempty"`
}
