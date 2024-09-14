package iot

import types "libds/aws/types"

type IndexingConfiguration struct {
	// Thing group indexing configuration. See below.
	ThingGroupIndexingConfiguration types.Iot_IndexingConfigurationThingGroupIndexingConfiguration `json:"thingGroupIndexingConfiguration,omitempty" yaml:"thingGroupIndexingConfiguration,omitempty"`

	// Thing indexing configuration. See below.
	ThingIndexingConfiguration types.Iot_IndexingConfigurationThingIndexingConfiguration `json:"thingIndexingConfiguration,omitempty" yaml:"thingIndexingConfiguration,omitempty"`
}
