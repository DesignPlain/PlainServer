package types

type Container_ClusterWorkloadAltsConfig struct {
	// Whether the alts handshaker should be enabled or not for direct-path. Requires Workload Identity (workloadPool) must be non-empty).
	EnableAlts bool `json:"enableAlts,omitempty" yaml:"enableAlts,omitempty"`
}
