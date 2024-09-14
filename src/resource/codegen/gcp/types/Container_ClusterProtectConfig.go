package types

type Container_ClusterProtectConfig struct {
	// WorkloadConfig defines which actions are enabled for a cluster's workload configurations. Structure is documented below
	WorkloadConfig Container_ClusterProtectConfigWorkloadConfig `json:"workloadConfig,omitempty" yaml:"workloadConfig,omitempty"`

	// Sets which mode to use for Protect workload vulnerability scanning feature. Accepted values are DISABLED, BASIC.
	WorkloadVulnerabilityMode string `json:"workloadVulnerabilityMode,omitempty" yaml:"workloadVulnerabilityMode,omitempty"`
}
