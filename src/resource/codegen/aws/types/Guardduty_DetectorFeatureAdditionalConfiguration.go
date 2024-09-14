package types

type Guardduty_DetectorFeatureAdditionalConfiguration struct {
	// The status of the additional configuration. Valid values: `ENABLED`, `DISABLED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// The name of the additional configuration for a feature. Valid values: `EKS_ADDON_MANAGEMENT`, `ECS_FARGATE_AGENT_MANAGEMENT`, `EC2_AGENT_MANAGEMENT`. Refer to the [AWS Documentation](https://docs.aws.amazon.com/guardduty/latest/APIReference/API_DetectorAdditionalConfiguration.html) for the current list of supported values.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
