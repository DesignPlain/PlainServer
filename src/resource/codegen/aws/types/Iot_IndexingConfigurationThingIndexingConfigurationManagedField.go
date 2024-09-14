package types

type Iot_IndexingConfigurationThingIndexingConfigurationManagedField struct {
	// The data type of the field. Valid values: `Number`, `String`, `Boolean`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// The name of the field.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
