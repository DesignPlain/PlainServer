package appconfig

type DeploymentStrategy struct {
	// Map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Total amount of time for a deployment to last. Minimum value of 0, maximum value of 1440.
	DeploymentDurationInMinutes int `json:"deploymentDurationInMinutes,omitempty" yaml:"deploymentDurationInMinutes,omitempty"`

	// Description of the deployment strategy. Can be at most 1024 characters.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Amount of time AWS AppConfig monitors for alarms before considering the deployment to be complete and no longer eligible for automatic roll back. Minimum value of 0, maximum value of 1440.
	FinalBakeTimeInMinutes int `json:"finalBakeTimeInMinutes,omitempty" yaml:"finalBakeTimeInMinutes,omitempty"`

	// Percentage of targets to receive a deployed configuration during each interval. Minimum value of 1.0, maximum value of 100.0.
	GrowthFactor float64 `json:"growthFactor,omitempty" yaml:"growthFactor,omitempty"`

	// Algorithm used to define how percentage grows over time. Valid value: `LINEAR` and `EXPONENTIAL`. Defaults to `LINEAR`.
	GrowthType string `json:"growthType,omitempty" yaml:"growthType,omitempty"`

	// Name for the deployment strategy. Must be between 1 and 64 characters in length.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Where to save the deployment strategy. Valid values: `NONE` and `SSM_DOCUMENT`.
	ReplicateTo string `json:"replicateTo,omitempty" yaml:"replicateTo,omitempty"`
}
