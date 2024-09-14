package types

type Monitoring_NotificationChannelSensitiveLabels struct {
	/*
	   An authorization token for a notification channel. Channel types that support this field include: slack
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	AuthToken string `json:"authToken,omitempty" yaml:"authToken,omitempty"`

	/*
	   An password for a notification channel. Channel types that support this field include: webhook_basicauth
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	/*
	   An servicekey token for a notification channel. Channel types that support this field include: pagerduty
	   --Note--: This property is sensitive and will not be displayed in the plan.
	*/
	ServiceKey string `json:"serviceKey,omitempty" yaml:"serviceKey,omitempty"`
}
