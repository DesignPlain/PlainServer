package types

type Msk_ClusterOpenMonitoring struct {
	// Configuration block for Prometheus settings for open monitoring. See below.
	Prometheus Msk_ClusterOpenMonitoringPrometheus `json:"prometheus,omitempty" yaml:"prometheus,omitempty"`
}
