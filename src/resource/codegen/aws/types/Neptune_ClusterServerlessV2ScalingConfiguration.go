package types

type Neptune_ClusterServerlessV2ScalingConfiguration struct {
	// The maximum Neptune Capacity Units (NCUs) for this cluster. Must be lower or equal than --128--. See [AWS Documentation](https://docs.aws.amazon.com/neptune/latest/userguide/neptune-serverless-capacity-scaling.html) for more details.
	MaxCapacity float64 `json:"maxCapacity,omitempty" yaml:"maxCapacity,omitempty"`

	// The minimum Neptune Capacity Units (NCUs) for this cluster. Must be greater or equal than --1--. See [AWS Documentation](https://docs.aws.amazon.com/neptune/latest/userguide/neptune-serverless-capacity-scaling.html) for more details.
	MinCapacity float64 `json:"minCapacity,omitempty" yaml:"minCapacity,omitempty"`
}
