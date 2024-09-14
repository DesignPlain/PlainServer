package types

type Cloudformation_StackSetOperationPreferences struct {
	// The number of accounts, per Region, for which this operation can fail before AWS CloudFormation stops the operation in that Region.
	FailureToleranceCount int `json:"failureToleranceCount,omitempty" yaml:"failureToleranceCount,omitempty"`

	// The percentage of accounts, per Region, for which this stack operation can fail before AWS CloudFormation stops the operation in that Region.
	FailureTolerancePercentage int `json:"failureTolerancePercentage,omitempty" yaml:"failureTolerancePercentage,omitempty"`

	// The maximum number of accounts in which to perform this operation at one time.
	MaxConcurrentCount int `json:"maxConcurrentCount,omitempty" yaml:"maxConcurrentCount,omitempty"`

	// The maximum percentage of accounts in which to perform this operation at one time.
	MaxConcurrentPercentage int `json:"maxConcurrentPercentage,omitempty" yaml:"maxConcurrentPercentage,omitempty"`

	// The concurrency type of deploying StackSets operations in Regions, could be in parallel or one Region at a time.
	RegionConcurrencyType string `json:"regionConcurrencyType,omitempty" yaml:"regionConcurrencyType,omitempty"`

	// The order of the Regions in where you want to perform the stack operation.
	RegionOrders []string `json:"regionOrders,omitempty" yaml:"regionOrders,omitempty"`
}
