package types

type Cloudformation_StackSetInstanceOperationPreferences struct {
	// Order of the Regions in where you want to perform the stack operation.
	RegionOrders []string `json:"regionOrders,omitempty" yaml:"regionOrders,omitempty"`

	// Specifies how the concurrency level behaves during the operation execution. Valid values are `STRICT_FAILURE_TOLERANCE` and `SOFT_FAILURE_TOLERANCE`.
	ConcurrencyMode string `json:"concurrencyMode,omitempty" yaml:"concurrencyMode,omitempty"`

	// Number of accounts, per Region, for which this operation can fail before AWS CloudFormation stops the operation in that Region.
	FailureToleranceCount int `json:"failureToleranceCount,omitempty" yaml:"failureToleranceCount,omitempty"`

	// Percentage of accounts, per Region, for which this stack operation can fail before AWS CloudFormation stops the operation in that Region.
	FailureTolerancePercentage int `json:"failureTolerancePercentage,omitempty" yaml:"failureTolerancePercentage,omitempty"`

	// Maximum number of accounts in which to perform this operation at one time.
	MaxConcurrentCount int `json:"maxConcurrentCount,omitempty" yaml:"maxConcurrentCount,omitempty"`

	// Maximum percentage of accounts in which to perform this operation at one time.
	MaxConcurrentPercentage int `json:"maxConcurrentPercentage,omitempty" yaml:"maxConcurrentPercentage,omitempty"`

	// Concurrency type of deploying StackSets operations in Regions, could be in parallel or one Region at a time. Valid values are `SEQUENTIAL` and `PARALLEL`.
	RegionConcurrencyType string `json:"regionConcurrencyType,omitempty" yaml:"regionConcurrencyType,omitempty"`
}
