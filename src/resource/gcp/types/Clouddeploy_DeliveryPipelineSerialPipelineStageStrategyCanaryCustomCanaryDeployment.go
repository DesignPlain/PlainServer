package types

type Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyCanaryCustomCanaryDeployment struct {
	// Required. Configuration for each phase in the canary deployment in the order executed.
	PhaseConfigs []Clouddeploy_DeliveryPipelineSerialPipelineStageStrategyCanaryCustomCanaryDeploymentPhaseConfig `json:"phaseConfigs,omitempty" yaml:"phaseConfigs,omitempty"`
}
