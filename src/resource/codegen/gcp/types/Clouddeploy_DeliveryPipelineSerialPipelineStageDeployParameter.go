package types

type Clouddeploy_DeliveryPipelineSerialPipelineStageDeployParameter struct {
	// Optional. Deploy parameters are applied to targets with match labels. If unspecified, deploy parameters are applied to all targets (including child targets of a multi-target).
	MatchTargetLabels map[string]string `json:"matchTargetLabels,omitempty" yaml:"matchTargetLabels,omitempty"`

	// Required. Values are deploy parameters in key-value pairs.
	Values map[string]string `json:"values,omitempty" yaml:"values,omitempty"`
}
