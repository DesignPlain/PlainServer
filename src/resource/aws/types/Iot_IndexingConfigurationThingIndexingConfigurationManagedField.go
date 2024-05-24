package types

type Iot_IndexingConfigurationThingIndexingConfigurationManagedField struct {
	// The name of the field.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The data type of the field. Valid values: `Number`, `String`, `Boolean`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
