package redshift

type EventSubscription struct {
	// A list of identifiers of the event sources for which events will be returned. If not specified, then all sources are included in the response. If specified, a `source_type` must also be specified.
	SourceIds []string `json:"sourceIds,omitempty" yaml:"sourceIds,omitempty"`

	// The type of source that will be generating the events. Valid options are `cluster`, `cluster-parameter-group`, `cluster-security-group`, `cluster-snapshot`, or `scheduled-action`. If not set, all sources will be subscribed to.
	SourceType string `json:"sourceType,omitempty" yaml:"sourceType,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// A boolean flag to enable/disable the subscription. Defaults to `true`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// A list of event categories for a SourceType that you want to subscribe to. See https://docs.aws.amazon.com/redshift/latest/mgmt/working-with-event-notifications.html or run `aws redshift describe-event-categories`.
	EventCategories []string `json:"eventCategories,omitempty" yaml:"eventCategories,omitempty"`

	// The name of the Redshift event subscription.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The event severity to be published by the notification subscription. Valid options are `INFO` or `ERROR`. Default value of `INFO`.
	Severity string `json:"severity,omitempty" yaml:"severity,omitempty"`

	// The ARN of the SNS topic to send events to.
	SnsTopicArn string `json:"snsTopicArn,omitempty" yaml:"snsTopicArn,omitempty"`
}
