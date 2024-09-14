package types

type Datafusion_InstanceEventPublishConfig struct {
	// Option to enable Event Publishing.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// The resource name of the Pub/Sub topic. Format: projects/{projectId}/topics/{topic_id}
	Topic string `json:"topic,omitempty" yaml:"topic,omitempty"`
}
