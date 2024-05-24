package types

type Dynamodb_getTableLocalSecondaryIndex struct {
	// Name of the DynamoDB table.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	NonKeyAttributes []string `json:"nonKeyAttributes,omitempty" yaml:"nonKeyAttributes,omitempty"`

	//
	ProjectionType string `json:"projectionType,omitempty" yaml:"projectionType,omitempty"`

	//
	RangeKey string `json:"rangeKey,omitempty" yaml:"rangeKey,omitempty"`
}
