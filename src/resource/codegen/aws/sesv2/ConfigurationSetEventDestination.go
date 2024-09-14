package sesv2

import types "libds/aws/types"

type ConfigurationSetEventDestination struct {
	// An object that defines the event destination. See `event_destination` Block for details.
	EventDestinationName string `json:"eventDestinationName,omitempty" yaml:"eventDestinationName,omitempty"`

	// The name of the configuration set.
	ConfigurationSetName string `json:"configurationSetName,omitempty" yaml:"configurationSetName,omitempty"`

	// A name that identifies the event destination within the configuration set.
	EventDestination types.Sesv2_ConfigurationSetEventDestinationEventDestination `json:"eventDestination,omitempty" yaml:"eventDestination,omitempty"`
}
