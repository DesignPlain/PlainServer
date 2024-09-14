package types

type Rds_ClusterServerlessv2ScalingConfiguration struct {
	// Maximum capacity for an Aurora DB cluster in `provisioned` DB engine mode. The maximum capacity must be greater than or equal to the minimum capacity. Valid capacity values are in a range of `0.5` up to `128` in steps of `0.5`.
	MaxCapacity float64 `json:"maxCapacity,omitempty" yaml:"maxCapacity,omitempty"`

	// Minimum capacity for an Aurora DB cluster in `provisioned` DB engine mode. The minimum capacity must be lesser than or equal to the maximum capacity. Valid capacity values are in a range of `0.5` up to `128` in steps of `0.5`.
	MinCapacity float64 `json:"minCapacity,omitempty" yaml:"minCapacity,omitempty"`
}
