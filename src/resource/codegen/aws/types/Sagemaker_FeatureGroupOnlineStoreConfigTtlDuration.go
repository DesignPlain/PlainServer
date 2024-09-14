package types

type Sagemaker_FeatureGroupOnlineStoreConfigTtlDuration struct {
	// TtlDuration time unit. Valid values are `Seconds`, `Minutes`, `Hours`, `Days`, or `Weeks`.
	Unit string `json:"unit,omitempty" yaml:"unit,omitempty"`

	// TtlDuration time value.
	Value int `json:"value,omitempty" yaml:"value,omitempty"`
}
