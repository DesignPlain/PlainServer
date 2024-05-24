package types

type Connect_QuickConnectQuickConnectConfig struct {
	// Specifies the user configuration of the Quick Connect. This is required only if `quick_connect_type` is `USER`. The `user_config` block is documented below.
	UserConfigs []Connect_QuickConnectQuickConnectConfigUserConfig `json:"userConfigs,omitempty" yaml:"userConfigs,omitempty"`

	// Specifies the phone configuration of the Quick Connect. This is required only if `quick_connect_type` is `PHONE_NUMBER`. The `phone_config` block is documented below.
	PhoneConfigs []Connect_QuickConnectQuickConnectConfigPhoneConfig `json:"phoneConfigs,omitempty" yaml:"phoneConfigs,omitempty"`

	// Specifies the queue configuration of the Quick Connect. This is required only if `quick_connect_type` is `QUEUE`. The `queue_config` block is documented below.
	QueueConfigs []Connect_QuickConnectQuickConnectConfigQueueConfig `json:"queueConfigs,omitempty" yaml:"queueConfigs,omitempty"`

	// Specifies the configuration type of the quick connect. valid values are `PHONE_NUMBER`, `QUEUE`, `USER`.
	QuickConnectType string `json:"quickConnectType,omitempty" yaml:"quickConnectType,omitempty"`
}
