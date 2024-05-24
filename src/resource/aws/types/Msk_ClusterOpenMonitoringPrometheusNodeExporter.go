package types

type Msk_ClusterOpenMonitoringPrometheusNodeExporter struct {
	// Indicates whether you want to enable or disable the Node Exporter.
	EnabledInBroker bool `json:"enabledInBroker,omitempty" yaml:"enabledInBroker,omitempty"`
}
