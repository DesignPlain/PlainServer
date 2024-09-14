package types

type Container_AttachedClusterMonitoringConfig struct {
	/*
	   Enable Google Cloud Managed Service for Prometheus in the cluster.
	   Structure is documented below.
	*/
	ManagedPrometheusConfig Container_AttachedClusterMonitoringConfigManagedPrometheusConfig `json:"managedPrometheusConfig,omitempty" yaml:"managedPrometheusConfig,omitempty"`
}
