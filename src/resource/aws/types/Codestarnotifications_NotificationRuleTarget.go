package types

type Codestarnotifications_NotificationRuleTarget struct {
	// The type of the notification target. Default value is `SNS`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The ARN of notification rule target. For example, a SNS Topic ARN.
	Address string `json:"address,omitempty" yaml:"address,omitempty"`

	// The status of the notification rule. Possible values are `ENABLED` and `DISABLED`, default is `ENABLED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`
}
