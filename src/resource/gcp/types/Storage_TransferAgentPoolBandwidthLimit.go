package types

type Storage_TransferAgentPoolBandwidthLimit struct {
	// Bandwidth rate in megabytes per second, distributed across all the agents in the pool.
	LimitMbps string `json:"limitMbps,omitempty" yaml:"limitMbps,omitempty"`
}
