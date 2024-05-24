package connect

import types "DesignSphere_Server/src/resource/aws/types"

type QuickConnect struct {
	// Specifies the description of the Quick Connect.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	// Specifies the name of the Quick Connect.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A block that defines the configuration information for the Quick Connect: `quick_connect_type` and one of `phone_config`, `queue_config`, `user_config` . The Quick Connect Config block is documented below.
	QuickConnectConfig types.Connect_QuickConnectQuickConnectConfig `json:"quickConnectConfig,omitempty" yaml:"quickConnectConfig,omitempty"`

	// Tags to apply to the Quick Connect. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
