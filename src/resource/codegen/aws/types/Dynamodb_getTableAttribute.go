package types

type Dynamodb_getTableAttribute struct {
	// Name of the DynamoDB table.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	//
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
