package costexplorer

import types "libds/aws/types"

type AnomalySubscription struct {
	// The name for the subscription.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A subscriber configuration. Multiple subscribers can be defined.
	Subscribers []types.Costexplorer_AnomalySubscriptionSubscriber `json:"subscribers,omitempty" yaml:"subscribers,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// An Expression object used to specify the anomalies that you want to generate alerts for. See Threshold Expression.
	ThresholdExpression types.Costexplorer_AnomalySubscriptionThresholdExpression `json:"thresholdExpression,omitempty" yaml:"thresholdExpression,omitempty"`

	// The unique identifier for the AWS account in which the anomaly subscription ought to be created.
	AccountId string `json:"accountId,omitempty" yaml:"accountId,omitempty"`

	// The frequency that anomaly reports are sent. Valid Values: `DAILY` | `IMMEDIATE` | `WEEKLY`.
	Frequency string `json:"frequency,omitempty" yaml:"frequency,omitempty"`

	// A list of cost anomaly monitors.
	MonitorArnLists []string `json:"monitorArnLists,omitempty" yaml:"monitorArnLists,omitempty"`
}
