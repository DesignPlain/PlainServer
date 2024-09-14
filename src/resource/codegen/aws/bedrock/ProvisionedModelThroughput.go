package bedrock

import types "libds/aws/types"

type ProvisionedModelThroughput struct {
	//
	Timeouts types.Bedrock_ProvisionedModelThroughputTimeouts `json:"timeouts,omitempty" yaml:"timeouts,omitempty"`

	// Commitment duration requested for the Provisioned Throughput. For custom models, you can purchase on-demand Provisioned Throughput by omitting this argument. Valid values: `OneMonth`, `SixMonths`.
	CommitmentDuration string `json:"commitmentDuration,omitempty" yaml:"commitmentDuration,omitempty"`

	// ARN of the model to associate with this Provisioned Throughput.
	ModelArn string `json:"modelArn,omitempty" yaml:"modelArn,omitempty"`

	// Number of model units to allocate. A model unit delivers a specific throughput level for the specified model.
	ModelUnits int `json:"modelUnits,omitempty" yaml:"modelUnits,omitempty"`

	// Unique name for this Provisioned Throughput.
	ProvisionedModelName string `json:"provisionedModelName,omitempty" yaml:"provisionedModelName,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
