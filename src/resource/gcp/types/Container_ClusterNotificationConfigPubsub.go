package types

type Container_ClusterNotificationConfigPubsub struct {
	// Whether or not the notification config is enabled
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Choose what type of notifications you want to receive. If no filters are applied, you'll receive all notification  Structure is documented below.
	Filter Container_ClusterNotificationConfigPubsubFilter `json:"filter,omitempty" yaml:"filter,omitempty"`

	// The pubsub topic to push upgrade notifications to. Must be in the same project as the cluster. Must be in the format: `projects/{project}/topics/{topic}`.
	Topic string `json:"topic,omitempty" yaml:"topic,omitempty"`
}
