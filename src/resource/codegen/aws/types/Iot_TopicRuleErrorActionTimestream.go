package types

type Iot_TopicRuleErrorActionTimestream struct {
	// The name of an Amazon Timestream database.
	DatabaseName string `json:"databaseName,omitempty" yaml:"databaseName,omitempty"`

	// Configuration blocks with metadata attributes of the time series that are written in each measure record. Nested arguments below.
	Dimensions []Iot_TopicRuleErrorActionTimestreamDimension `json:"dimensions,omitempty" yaml:"dimensions,omitempty"`

	// The ARN of the role that grants permission to write to the Amazon Timestream database table.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// The name of the database table into which to write the measure records.
	TableName string `json:"tableName,omitempty" yaml:"tableName,omitempty"`

	// Configuration block specifying an application-defined value to replace the default value assigned to the Timestream record's timestamp in the time column. Nested arguments below.
	Timestamp Iot_TopicRuleErrorActionTimestreamTimestamp `json:"timestamp,omitempty" yaml:"timestamp,omitempty"`
}
