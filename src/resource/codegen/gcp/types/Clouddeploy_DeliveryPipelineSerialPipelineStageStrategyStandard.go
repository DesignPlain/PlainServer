package types

type Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyStandard struct {
	// Optional. Configuration for the postdeploy job. If this is not configured, postdeploy job will not be present.
	Postdeploy Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyStandardPostdeploy `json:"postdeploy,omitempty" yaml:"postdeploy,omitempty"`

	// Optional. Configuration for the predeploy job. If this is not configured, predeploy job will not be present.
	Predeploy Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyStandardPredeploy `json:"predeploy,omitempty" yaml:"predeploy,omitempty"`

	// Whether to verify a deployment.
	Verify bool `json:"verify,omitempty" yaml:"verify,omitempty"`
}
