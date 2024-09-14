package cloudwatch

type EventArchive struct {
	// The description of the new event archive.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Instructs the new event archive to only capture events matched by this pattern. By default, it attempts to archive every event received in the `event_source_arn`.
	EventPattern string `json:"eventPattern,omitempty" yaml:"eventPattern,omitempty"`

	// Event bus source ARN from where these events should be archived.
	EventSourceArn string `json:"eventSourceArn,omitempty" yaml:"eventSourceArn,omitempty"`

	// The name of the new event archive. The archive name cannot exceed 48 characters.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The maximum number of days to retain events in the new event archive. By default, it archives indefinitely.
	RetentionDays int `json:"retentionDays,omitempty" yaml:"retentionDays,omitempty"`
}
