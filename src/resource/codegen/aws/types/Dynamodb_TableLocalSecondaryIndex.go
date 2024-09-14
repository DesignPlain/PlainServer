package types

type Dynamodb_TableLocalSecondaryIndex struct {
	// Name of the index
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Only required with `INCLUDE` as a projection type; a list of attributes to project into the index. These do not need to be defined as attributes on the table.
	NonKeyAttributes []string `json:"nonKeyAttributes,omitempty" yaml:"nonKeyAttributes,omitempty"`

	// One of `ALL`, `INCLUDE` or `KEYS_ONLY` where `ALL` projects every attribute into the index, `KEYS_ONLY` projects  into the index only the table and index hash_key and sort_key attributes ,  `INCLUDE` projects into the index all of the attributes that are defined in `non_key_attributes` in addition to the attributes that that`KEYS_ONLY` project.
	ProjectionType string `json:"projectionType,omitempty" yaml:"projectionType,omitempty"`

	// Name of the range key.
	RangeKey string `json:"rangeKey,omitempty" yaml:"rangeKey,omitempty"`
}
