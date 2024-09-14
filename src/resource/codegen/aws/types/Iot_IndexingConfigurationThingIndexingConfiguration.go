package types

type Iot_IndexingConfigurationThingIndexingConfiguration struct {
	// Thing indexing mode. Valid values: `REGISTRY`, `REGISTRY_AND_SHADOW`, `OFF`.
	ThingIndexingMode string `json:"thingIndexingMode,omitempty" yaml:"thingIndexingMode,omitempty"`

	// Contains custom field names and their data type. See below.
	CustomFields []Iot_IndexingConfigurationThingIndexingConfigurationCustomField `json:"customFields,omitempty" yaml:"customFields,omitempty"`

	// Device Defender indexing mode. Valid values: `VIOLATIONS`, `OFF`. Default: `OFF`.
	DeviceDefenderIndexingMode string `json:"deviceDefenderIndexingMode,omitempty" yaml:"deviceDefenderIndexingMode,omitempty"`

	// Required if `named_shadow_indexing_mode` is `ON`. Enables to add named shadows filtered by `filter` to fleet indexing configuration.
	Filter Iot_IndexingConfigurationThingIndexingConfigurationFilter `json:"filter,omitempty" yaml:"filter,omitempty"`

	// Contains fields that are indexed and whose types are already known by the Fleet Indexing service. See below.
	ManagedFields []Iot_IndexingConfigurationThingIndexingConfigurationManagedField `json:"managedFields,omitempty" yaml:"managedFields,omitempty"`

	// [Named shadow](https://docs.aws.amazon.com/iot/latest/developerguide/iot-device-shadows.html) indexing mode. Valid values: `ON`, `OFF`. Default: `OFF`.
	NamedShadowIndexingMode string `json:"namedShadowIndexingMode,omitempty" yaml:"namedShadowIndexingMode,omitempty"`

	// Thing connectivity indexing mode. Valid values: `STATUS`, `OFF`. Default: `OFF`.
	ThingConnectivityIndexingMode string `json:"thingConnectivityIndexingMode,omitempty" yaml:"thingConnectivityIndexingMode,omitempty"`
}
