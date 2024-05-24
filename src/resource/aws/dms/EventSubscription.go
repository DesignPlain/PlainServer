package dms

type EventSubscription struct {
	// List of event categories to listen for, see `DescribeEventCategories` for a canonical list.
	EventCategories []string `json:"eventCategories,omitempty" yaml:"eventCategories,omitempty"`

	// Name of event subscription.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// SNS topic arn to send events on.
	SnsTopicArn string `json:"snsTopicArn,omitempty" yaml:"snsTopicArn,omitempty"`

	// Ids of sources to listen to. If you don't specify a value, notifications are provided for all sources.
	SourceIds []string `json:"sourceIds,omitempty" yaml:"sourceIds,omitempty"`

	// Type of source for events. Valid values: `replication-instance` or `replication-task`
	SourceType string `json:"sourceType,omitempty" yaml:"sourceType,omitempty"`

	// Map of resource tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Whether the event subscription should be enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
