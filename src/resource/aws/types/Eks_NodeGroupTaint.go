package types

type Eks_NodeGroupTaint struct {
	// The key of the taint. Maximum length of 63.
	Key string `json:"key,omitempty" yaml:"key,omitempty"`

	// The value of the taint. Maximum length of 63.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`

	// The effect of the taint. Valid values: `NO_SCHEDULE`, `NO_EXECUTE`, `PREFER_NO_SCHEDULE`.
	Effect string `json:"effect,omitempty" yaml:"effect,omitempty"`
}
