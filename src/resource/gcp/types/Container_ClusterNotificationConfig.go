package types



type Container_ClusterNotificationConfig struct {
	// The pubsub config for the cluster's upgrade notifications.
	Pubsub Container_ClusterNotificationConfigPubsub `json:"pubsub,omitempty" yaml:"pubsub,omitempty"`
}
