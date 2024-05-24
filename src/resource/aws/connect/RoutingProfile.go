package connect

import types "DesignSphere_Server/src/resource/aws/types"

type RoutingProfile struct {
	// Specifies the default outbound queue for the Routing Profile.
	DefaultOutboundQueueId string `json:"defaultOutboundQueueId,omitempty" yaml:"defaultOutboundQueueId,omitempty"`

	// Specifies the description of the Routing Profile.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Specifies the identifier of the hosting Amazon Connect Instance.
	InstanceId string `json:"instanceId,omitempty" yaml:"instanceId,omitempty"`

	// One or more `media_concurrencies` blocks that specify the channels that agents can handle in the Contact Control Panel (CCP) for this Routing Profile. The `media_concurrencies` block is documented below.
	MediaConcurrencies []types.Connect_RoutingProfileMediaConcurrency `json:"mediaConcurrencies,omitempty" yaml:"mediaConcurrencies,omitempty"`

	// Specifies the name of the Routing Profile.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// One or more `queue_configs` blocks that specify the inbound queues associated with the routing profile. If no queue is added, the agent only can make outbound calls. The `queue_configs` block is documented below.
	QueueConfigs []types.Connect_RoutingProfileQueueConfig `json:"queueConfigs,omitempty" yaml:"queueConfigs,omitempty"`

	/*
	   Tags to apply to the Routing Profile. If configured with a provider
	   `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	*/
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
