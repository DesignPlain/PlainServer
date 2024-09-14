package types

type Xray_GroupInsightsConfiguration struct {
	// Specifies whether insights are enabled.
	InsightsEnabled bool `json:"insightsEnabled,omitempty" yaml:"insightsEnabled,omitempty"`

	// Specifies whether insight notifications are enabled.
	NotificationsEnabled bool `json:"notificationsEnabled,omitempty" yaml:"notificationsEnabled,omitempty"`
}
