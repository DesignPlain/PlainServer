package types

type Dms_EndpointKinesisSettings struct {
	// Shows the partition value within the Kinesis message output, unless the partition type is schema-table-type. Default is `false`.
	IncludePartitionValue bool `json:"includePartitionValue,omitempty" yaml:"includePartitionValue,omitempty"`

	// Includes any data definition language (DDL) operations that change the table in the control data. Default is `false`.
	IncludeTableAlterOperations bool `json:"includeTableAlterOperations,omitempty" yaml:"includeTableAlterOperations,omitempty"`

	// Output format for the records created. Default is `json`. Valid values are `json` and `json-unformatted` (a single line with no tab).
	MessageFormat string `json:"messageFormat,omitempty" yaml:"messageFormat,omitempty"`

	// Prefixes schema and table names to partition values, when the partition type is primary-key-type. Default is `false`.
	PartitionIncludeSchemaTable bool `json:"partitionIncludeSchemaTable,omitempty" yaml:"partitionIncludeSchemaTable,omitempty"`

	// ARN of the IAM Role with permissions to write to the Kinesis data stream.
	ServiceAccessRoleArn string `json:"serviceAccessRoleArn,omitempty" yaml:"serviceAccessRoleArn,omitempty"`

	// ARN of the Kinesis data stream.
	StreamArn string `json:"streamArn,omitempty" yaml:"streamArn,omitempty"`

	// Shows detailed control information for table definition, column definition, and table and column changes in the Kinesis message output. Default is `false`.
	IncludeControlDetails bool `json:"includeControlDetails,omitempty" yaml:"includeControlDetails,omitempty"`

	// Include NULL and empty columns in the target. Default is `false`.
	IncludeNullAndEmpty bool `json:"includeNullAndEmpty,omitempty" yaml:"includeNullAndEmpty,omitempty"`

	// Provides detailed transaction information from the source database. Default is `false`.
	IncludeTransactionDetails bool `json:"includeTransactionDetails,omitempty" yaml:"includeTransactionDetails,omitempty"`
}
