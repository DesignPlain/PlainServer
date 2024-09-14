package types

type Devopsguru_NotificationChannelFilters struct {
	// Events to receive notifications for. Valid values are `NEW_INSIGHT`, `CLOSED_INSIGHT`, `NEW_ASSOCIATION`, `SEVERITY_UPGRADED`, and `NEW_RECOMMENDATION`.
	MessageTypes []string `json:"messageTypes,omitempty" yaml:"messageTypes,omitempty"`

	// Severity levels to receive notifications for. Valid values are `LOW`, `MEDIUM`, and `HIGH`.
	Severities []string `json:"severities,omitempty" yaml:"severities,omitempty"`
}
