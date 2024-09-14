package securitycenter

import types "libds/gcp/types"

type NotificationConfig struct {
	// This must be unique within the organization.
	ConfigId string `json:"configId,omitempty" yaml:"configId,omitempty"`

	// The description of the notification config (max of 1024 characters).
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The organization whose Cloud Security Command Center the Notification
	   Config lives in.
	*/
	Organization string `json:"organization,omitempty" yaml:"organization,omitempty"`

	/*
	   The Pub/Sub topic to send notifications to. Its format is
	   "projects/[project_id]/topics/[topic]".
	*/
	PubsubTopic string `json:"pubsubTopic,omitempty" yaml:"pubsubTopic,omitempty"`

	/*
	   The config for triggering streaming-based notifications.
	   Structure is documented below.
	*/
	StreamingConfig types.Securitycenter_NotificationConfigStreamingConfig `json:"streamingConfig,omitempty" yaml:"streamingConfig,omitempty"`
}
