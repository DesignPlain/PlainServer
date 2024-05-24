package cloudfront

import types "DesignSphere_Server/src/resource/aws/types"

type MonitoringSubscription struct {
	// The ID of the distribution that you are enabling metrics for.
	DistributionId string `json:"distributionId,omitempty" yaml:"distributionId,omitempty"`

	// A monitoring subscription. This structure contains information about whether additional CloudWatch metrics are enabled for a given CloudFront distribution.
	MonitoringSubscription types.Cloudfront_MonitoringSubscriptionMonitoringSubscription `json:"monitoringSubscription,omitempty" yaml:"monitoringSubscription,omitempty"`
}
