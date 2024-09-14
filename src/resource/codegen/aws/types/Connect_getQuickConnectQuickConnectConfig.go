package types

type Connect_getQuickConnectQuickConnectConfig struct {
	// Phone configuration of the Quick Connect. This is returned only if `quick_connect_type` is `PHONE_NUMBER`. The `phone_config` block is documented below.
	PhoneConfigs []Connect_getQuickConnectQuickConnectConfigPhoneConfig `json:"phoneConfigs,omitempty" yaml:"phoneConfigs,omitempty"`

	// Queue configuration of the Quick Connect. This is returned only if `quick_connect_type` is `QUEUE`. The `queue_config` block is documented below.
	QueueConfigs []Connect_getQuickConnectQuickConnectConfigQueueConfig `json:"queueConfigs,omitempty" yaml:"queueConfigs,omitempty"`

	// Configuration type of the Quick Connect. Valid values are `PHONE_NUMBER`, `QUEUE`, `USER`.
	QuickConnectType string `json:"quickConnectType,omitempty" yaml:"quickConnectType,omitempty"`

	// User configuration of the Quick Connect. This is returned only if `quick_connect_type` is `USER`. The `user_config` block is documented below.
	UserConfigs []Connect_getQuickConnectQuickConnectConfigUserConfig `json:"userConfigs,omitempty" yaml:"userConfigs,omitempty"`
}
