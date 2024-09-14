package types

type Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyCanaryRuntimeConfigCloudRun struct {
	// Whether Cloud Deploy should update the traffic stanza in a Cloud Run Service on the user's behalf to facilitate traffic splitting. This is required to be true for CanaryDeployments, but optional for CustomCanaryDeployments.
	AutomaticTrafficControl bool `json:"automaticTrafficControl,omitempty" yaml:"automaticTrafficControl,omitempty"`

	// Optional. A list of tags that are added to the canary revision while the canary phase is in progress.
	CanaryRevisionTags []string `json:"canaryRevisionTags,omitempty" yaml:"canaryRevisionTags,omitempty"`

	// Optional. A list of tags that are added to the prior revision while the canary phase is in progress.
	PriorRevisionTags []string `json:"priorRevisionTags,omitempty" yaml:"priorRevisionTags,omitempty"`

	// Optional. A list of tags that are added to the final stable revision when the stable phase is applied.
	StableRevisionTags []string `json:"stableRevisionTags,omitempty" yaml:"stableRevisionTags,omitempty"`
}
