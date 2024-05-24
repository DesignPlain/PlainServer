package types

type Dms_getEndpointKinesisSetting struct {
	//
	IncludeTableAlterOperations bool `json:"includeTableAlterOperations,omitempty" yaml:"includeTableAlterOperations,omitempty"`

	//
	MessageFormat string `json:"messageFormat,omitempty" yaml:"messageFormat,omitempty"`

	//
	ServiceAccessRoleArn string `json:"serviceAccessRoleArn,omitempty" yaml:"serviceAccessRoleArn,omitempty"`

	//
	StreamArn string `json:"streamArn,omitempty" yaml:"streamArn,omitempty"`

	//
	IncludePartitionValue bool `json:"includePartitionValue,omitempty" yaml:"includePartitionValue,omitempty"`

	//
	IncludeNullAndEmpty bool `json:"includeNullAndEmpty,omitempty" yaml:"includeNullAndEmpty,omitempty"`

	//
	IncludeTransactionDetails bool `json:"includeTransactionDetails,omitempty" yaml:"includeTransactionDetails,omitempty"`

	//
	PartitionIncludeSchemaTable bool `json:"partitionIncludeSchemaTable,omitempty" yaml:"partitionIncludeSchemaTable,omitempty"`

	//
	IncludeControlDetails bool `json:"includeControlDetails,omitempty" yaml:"includeControlDetails,omitempty"`
}
