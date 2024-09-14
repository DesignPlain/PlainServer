package location

import types "libds/aws/types"

type Map struct {
	// Configuration block with the map style selected from an available data provider. Detailed below.
	Configuration types.Location_MapConfiguration `json:"configuration,omitempty" yaml:"configuration,omitempty"`

	// An optional description for the map resource.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	/*
	   The name for the map resource.

	   The following arguments are optional:
	*/
	MapName string `json:"mapName,omitempty" yaml:"mapName,omitempty"`

	// Key-value tags for the map. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
