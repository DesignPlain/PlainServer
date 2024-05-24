package types

type Keyspaces_TableSchemaDefinitionClusteringKey struct {
	// The name of the clustering key column.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The order modifier. Valid values: `ASC`, `DESC`.
	OrderBy string `json:"orderBy,omitempty" yaml:"orderBy,omitempty"`
}
