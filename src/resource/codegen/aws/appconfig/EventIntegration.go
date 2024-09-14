package appconfig

import types "libds/aws/types"

type EventIntegration struct {
	// EventBridge bus.
	EventbridgeBus string `json:"eventbridgeBus,omitempty" yaml:"eventbridgeBus,omitempty"`

	// Name of the Event Integration.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Tags to apply to the Event Integration. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Description of the Event Integration.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Block that defines the configuration information for the event filter. The Event Filter block is documented below.
	EventFilter types.Appconfig_EventIntegrationEventFilter `json:"eventFilter,omitempty" yaml:"eventFilter,omitempty"`
}
