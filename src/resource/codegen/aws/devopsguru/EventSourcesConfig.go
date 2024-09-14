package devopsguru

import types "libds/aws/types"

type EventSourcesConfig struct {
	// Configuration information about the integration of DevOps Guru as the Consumer via EventBridge with another AWS Service. See `event_sources` below.
	EventSources []types.Devopsguru_EventSourcesConfigEventSource `json:"eventSources,omitempty" yaml:"eventSources,omitempty"`
}
