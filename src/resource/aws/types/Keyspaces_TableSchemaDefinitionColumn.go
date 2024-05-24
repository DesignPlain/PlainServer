package types

type Keyspaces_TableSchemaDefinitionColumn struct {
	// The name of the column.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The data type of the column. See the [Developer Guide](https://docs.aws.amazon.com/keyspaces/latest/devguide/cql.elements.html#cql.data-types) for a list of available data
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
