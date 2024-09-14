package types

type Container_getClusterNotificationConfig struct {
	// Notification config for Cloud Pub/Sub
	Pubsubs []Container_getClusterNotificationConfigPubsub `json:"pubsubs,omitempty" yaml:"pubsubs,omitempty"`
}
