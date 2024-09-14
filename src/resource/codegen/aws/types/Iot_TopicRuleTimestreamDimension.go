package types

type Iot_TopicRuleTimestreamDimension struct {
	// The metadata dimension name. This is the name of the column in the Amazon Timestream database table record.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// The value to write in this column of the database record.
	Value string `json:"value,omitempty" yaml:"value,omitempty"`
}
