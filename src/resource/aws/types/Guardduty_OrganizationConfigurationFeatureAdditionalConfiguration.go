package types

type Guardduty_OrganizationConfigurationFeatureAdditionalConfiguration struct {
	// The status of the additional configuration that will be configured for the organization. Valid values: `NEW`, `ALL`, `NONE`.
	AutoEnable string `json:"autoEnable,omitempty" yaml:"autoEnable,omitempty"`

	// The name of the additional configuration that will be configured for the organization. Valid values: `EKS_ADDON_MANAGEMENT`.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
