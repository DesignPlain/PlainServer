package types

type Guardduty_DetectorFeatureAdditionalConfiguration struct {
	// The status of the additional configuration. Valid values: `ENABLED`, `DISABLED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// The name of the additional configuration. Valid values: `EKS_ADDON_MANAGEMENT`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
