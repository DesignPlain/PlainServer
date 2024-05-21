package types

type Container_getClusterMonitoringConfig struct {
	// Configuration of Advanced Datapath Observability features.
	AdvancedDatapathObservabilityConfigs []Container_getClusterMonitoringConfigAdvancedDatapathObservabilityConfig `json:"advancedDatapathObservabilityConfigs,omitempty" yaml:"advancedDatapathObservabilityConfigs,omitempty"`

	// GKE components exposing metrics. Valid values include SYSTEM_COMPONENTS, APISERVER, SCHEDULER, CONTROLLER_MANAGER, STORAGE, HPA, POD, DAEMONSET, DEPLOYMENT, STATEFULSET and WORKLOADS.
	EnableComponents []string `json:"enableComponents,omitempty" yaml:"enableComponents,omitempty"`

	// Configuration for Google Cloud Managed Services for Prometheus.
	ManagedPrometheuses []Container_getClusterMonitoringConfigManagedPrometheus `json:"managedPrometheuses,omitempty" yaml:"managedPrometheuses,omitempty"`
}
