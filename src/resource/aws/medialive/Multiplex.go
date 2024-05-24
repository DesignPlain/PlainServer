package medialive

import types "DesignSphere_Server/src/resource/aws/types"

type Multiplex struct {
	// A list of availability zones. You must specify exactly two.
	AvailabilityZones []string `json:"availabilityZones,omitempty" yaml:"availabilityZones,omitempty"`

	// Multiplex settings. See Multiplex Settings for more details.
	MultiplexSettings types.Medialive_MultiplexMultiplexSettings `json:"multiplexSettings,omitempty" yaml:"multiplexSettings,omitempty"`

	/*
	   name of Multiplex.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Whether to start the Multiplex. Defaults to `false`.
	StartMultiplex bool `json:"startMultiplex,omitempty" yaml:"startMultiplex,omitempty"`

	// A map of tags to assign to the Multiplex. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
