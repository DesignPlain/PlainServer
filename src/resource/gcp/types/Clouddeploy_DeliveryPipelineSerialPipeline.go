package types

type Clouddeploy_DeliveryPipelineSerialPipeline struct {
	// Each stage specifies configuration for a `Target`. The ordering of this list defines the promotion flow.
	Stages []Clouddeploy_DeliveryPipelineSerialPipelineStage `json:"stages,omitempty" yaml:"stages,omitempty"`
}
