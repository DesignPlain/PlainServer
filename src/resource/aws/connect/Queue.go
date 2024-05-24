package connect

import types "DesignSphere_Server/src/resource/aws/types"

type Queue struct {
	// Specifies the description of the Queue. Valid values are `ENABLED`, `DISABLED`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// Specifies the description of the Queue.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	// A block that defines the outbound caller ID name, number, and outbound whisper flow. The Outbound Caller Config block is documented below.
	OutboundCallerConfig types.Connect_QueueOutboundCallerConfig `json:"outboundCallerConfig,omitempty" yaml:"outboundCallerConfig,omitempty"`

	// Specifies a list of quick connects ids that determine the quick connects available to agents who are working the queue.
	QuickConnectIds []string `json:"quickConnectIds,omitempty" yaml:"quickConnectIds,omitempty"`

	// Specifies the identifier of the Hours of Operation.
	HoursOfOperationId string `json:"hoursOfOperationId,omitempty" yaml:"hoursOfOperationId,omitempty"`

	// Specifies the maximum number of contacts that can be in the queue before it is considered full. Minimum value of 0.
	MaxContacts int `json:"maxContacts,omitempty" yaml:"maxContacts,omitempty"`

	// Specifies the name of the Queue.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Tags to apply to the Queue. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
