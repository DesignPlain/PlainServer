package types

type Dynamodb_getTableGlobalSecondaryIndex struct {
	//
	RangeKey string `json:"rangeKey,omitempty" yaml:"rangeKey,omitempty"`

	//
	ReadCapacity int `json:"readCapacity,omitempty" yaml:"readCapacity,omitempty"`

	//
	WriteCapacity int `json:"writeCapacity,omitempty" yaml:"writeCapacity,omitempty"`

	//
	HashKey string `json:"hashKey,omitempty" yaml:"hashKey,omitempty"`

	// Name of the DynamoDB table.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	NonKeyAttributes []string `json:"nonKeyAttributes,omitempty" yaml:"nonKeyAttributes,omitempty"`

	//
	ProjectionType string `json:"projectionType,omitempty" yaml:"projectionType,omitempty"`
}
