package types

type Devopsguru_getNotificationChannelFilter struct {
	// Events to receive notifications for.
	MessageTypes []string `json:"messageTypes,omitempty" yaml:"messageTypes,omitempty"`

	// Severity levels to receive notifications for.
	Severities []string `json:"severities,omitempty" yaml:"severities,omitempty"`
}
