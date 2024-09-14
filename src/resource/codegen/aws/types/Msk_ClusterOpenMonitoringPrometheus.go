package types

type Msk_ClusterOpenMonitoringPrometheus struct {
	// Configuration block for JMX Exporter. See below.
	JmxExporter Msk_ClusterOpenMonitoringPrometheusJmxExporter `json:"jmxExporter,omitempty" yaml:"jmxExporter,omitempty"`

	// Configuration block for Node Exporter. See below.
	NodeExporter Msk_ClusterOpenMonitoringPrometheusNodeExporter `json:"nodeExporter,omitempty" yaml:"nodeExporter,omitempty"`
}
