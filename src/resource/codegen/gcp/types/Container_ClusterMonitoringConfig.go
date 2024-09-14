package types

type Container_ClusterMonitoringConfig struct {
	// Configuration for Advanced Datapath Monitoring. Structure is documented below.
	AdvancedDatapathObservabilityConfigs []Container_ClusterMonitoringConfigAdvancedDatapathObservabilityConfig `json:"advancedDatapathObservabilityConfigs,omitempty" yaml:"advancedDatapathObservabilityConfigs,omitempty"`

	// The GKE components exposing metrics. Supported values include: `SYSTEM_COMPONENTS`, `APISERVER`, `SCHEDULER`, `CONTROLLER_MANAGER`, `STORAGE`, `HPA`, `POD`, `DAEMONSET`, `DEPLOYMENT` and `STATEFULSET`. In beta provider, `WORKLOADS` is supported on top of those 10 values. (`WORKLOADS` is deprecated and removed in GKE 1.24.)
	EnableComponents []string `json:"enableComponents,omitempty" yaml:"enableComponents,omitempty"`

	// Configuration for Managed Service for Prometheus. Structure is documented below.
	ManagedPrometheus Container_ClusterMonitoringConfigManagedPrometheus `json:"managedPrometheus,omitempty" yaml:"managedPrometheus,omitempty"`
}
