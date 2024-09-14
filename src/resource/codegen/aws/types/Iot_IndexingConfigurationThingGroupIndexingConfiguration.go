package types

type Iot_IndexingConfigurationThingGroupIndexingConfiguration struct {
	// A list of thing group fields to index. This list cannot contain any managed fields. See below.
	CustomFields []Iot_IndexingConfigurationThingGroupIndexingConfigurationCustomField `json:"customFields,omitempty" yaml:"customFields,omitempty"`

	// Contains fields that are indexed and whose types are already known by the Fleet Indexing service. See below.
	ManagedFields []Iot_IndexingConfigurationThingGroupIndexingConfigurationManagedField `json:"managedFields,omitempty" yaml:"managedFields,omitempty"`

	// Thing group indexing mode. Valid values: `OFF`, `ON`.
	ThingGroupIndexingMode string `json:"thingGroupIndexingMode,omitempty" yaml:"thingGroupIndexingMode,omitempty"`
}
