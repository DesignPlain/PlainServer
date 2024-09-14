package types

type Guardduty_OrganizationConfigurationFeatureAdditionalConfiguration struct {
	// The status of the additional configuration that will be configured for the organization. Valid values: `NEW`, `ALL`, `NONE`.
	AutoEnable string `json:"autoEnable,omitempty" yaml:"autoEnable,omitempty"`

	// The name of the additional configuration for a feature that will be configured for the organization. Valid values: `EKS_ADDON_MANAGEMENT`, `ECS_FARGATE_AGENT_MANAGEMENT`, `EC2_AGENT_MANAGEMENT`. Refer to the [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DetectorAdditionalConfiguration.html) for the current list of supported values.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
