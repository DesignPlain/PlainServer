package types

type Cloudfront_MonitoringSubscriptionMonitoringSubscriptionRealtimeMetricsSubscriptionConfig struct {
	// A flag that indicates whether additional CloudWatch metrics are enabled for a given CloudFront distribution. Valid values are `Enabled` and `Disabled`. See below.
	RealtimeMetricsSubscriptionStatus string `json:"realtimeMetricsSubscriptionStatus,omitempty" yaml:"realtimeMetricsSubscriptionStatus,omitempty"`
}
