package types

type Monitoring_getNotificationChannelSensitiveLabel struct {
	// An servicekey token for a notification channel. Channel types that support this field include: pagerduty
	ServiceKey string `json:"serviceKey,omitempty" yaml:"serviceKey,omitempty"`

	// An authorization token for a notification channel. Channel types that support this field include: slack
	AuthToken string `json:"authToken,omitempty" yaml:"authToken,omitempty"`

	// An password for a notification channel. Channel types that support this field include: webhook_basicauth
	Password string `json:"password,omitempty" yaml:"password,omitempty"`
}
