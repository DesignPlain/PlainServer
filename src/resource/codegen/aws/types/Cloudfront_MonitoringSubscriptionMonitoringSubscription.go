package types

type Cloudfront_MonitoringSubscriptionMonitoringSubscription struct {
	// A subscription configuration for additional CloudWatch metrics. See below.
	RealtimeMetricsSubscriptionConfig Cloudfront_MonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfig `json:"realtimeMetricsSubscriptionConfig,omitempty" yaml:"realtimeMetricsSubscriptionConfig,omitempty"`
}
