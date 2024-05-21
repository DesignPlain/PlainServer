package types

type Clouddeploy_DeliveryPipelineConditionTargetsTypeCondition struct {
	// Human readable error message.
	ErrorDetails string `json:"errorDetails,omitempty" yaml:"errorDetails,omitempty"`

	// True if the targets are all a comparable type. For example this is true if all targets are GKE clusters. This is false if some targets are Cloud Run targets and others are GKE clusters.
	Status bool `json:"status,omitempty" yaml:"status,omitempty"`
}
