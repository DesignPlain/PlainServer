package types

type Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyStandardPredeploy struct {
	// Optional. A sequence of skaffold custom actions to invoke during execution of the predeploy job.
	Actions []string `json:"actions,omitempty" yaml:"actions,omitempty"`
}
