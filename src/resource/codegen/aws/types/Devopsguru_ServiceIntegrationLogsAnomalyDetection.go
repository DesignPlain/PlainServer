package types

type Devopsguru_ServiceIntegrationLogsAnomalyDetection struct {
	// Specifies if DevOps Guru is configured to perform log anomaly detection on CloudWatch log groups. Valid values are `DISABLED` and `ENABLED`.
	OptInStatus string `json:"optInStatus,omitempty" yaml:"optInStatus,omitempty"`
}
