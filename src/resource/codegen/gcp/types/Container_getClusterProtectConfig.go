package types

type Container_getClusterProtectConfig struct {
	// WorkloadConfig defines which actions are enabled for a cluster's workload configurations.
	WorkloadConfigs []Container_getClusterProtectConfigWorkloadConfig `json:"workloadConfigs,omitempty" yaml:"workloadConfigs,omitempty"`

	// Sets which mode to use for Protect workload vulnerability scanning feature. Accepted values are DISABLED, BASIC.
	WorkloadVulnerabilityMode string `json:"workloadVulnerabilityMode,omitempty" yaml:"workloadVulnerabilityMode,omitempty"`
}
