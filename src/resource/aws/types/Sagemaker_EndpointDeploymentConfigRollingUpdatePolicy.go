package types

type Sagemaker_EndpointDeploymentConfigRollingUpdatePolicy struct {
	// Batch size for rollback to the old endpoint fleet. Each rolling step to provision capacity and turn on traffic on the old endpoint fleet, and terminate capacity on the new endpoint fleet. If this field is absent, the default value will be set to 100%!o(MISSING)f total capacity which means to bring up the whole capacity of the old fleet at once during rollback. See Rollback Maximum Batch Size.
	RollbackMaximumBatchSize Sagemaker_EndpointDeploymentConfigRollingUpdatePolicyRollbackMaximumBatchSize `json:"rollbackMaximumBatchSize,omitempty" yaml:"rollbackMaximumBatchSize,omitempty"`

	// The length of the baking period, during which SageMaker monitors alarms for each batch on the new fleet. Valid values are between `0` and `3600`.
	WaitIntervalInSeconds int `json:"waitIntervalInSeconds,omitempty" yaml:"waitIntervalInSeconds,omitempty"`

	// Batch size for each rolling step to provision capacity and turn on traffic on the new endpoint fleet, and terminate capacity on the old endpoint fleet. Value must be between 5%!t(MISSING)o 50%!o(MISSING)f the variant's total instance count. See Maximum Batch Size.
	MaximumBatchSize Sagemaker_EndpointDeploymentConfigRollingUpdatePolicyMaximumBatchSize `json:"maximumBatchSize,omitempty" yaml:"maximumBatchSize,omitempty"`

	// The time limit for the total deployment. Exceeding this limit causes a timeout. Valid values are between `600` and `14400`.
	MaximumExecutionTimeoutInSeconds int `json:"maximumExecutionTimeoutInSeconds,omitempty" yaml:"maximumExecutionTimeoutInSeconds,omitempty"`
}
