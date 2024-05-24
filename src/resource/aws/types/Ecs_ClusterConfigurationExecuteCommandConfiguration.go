package types

type Ecs_ClusterConfigurationExecuteCommandConfiguration struct {
	// The AWS Key Management Service key ID to encrypt the data between the local client and the container.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// The log configuration for the results of the execute command actions Required when `logging` is `OVERRIDE`. Detailed below.
	LogConfiguration Ecs_ClusterConfigurationExecuteCommandConfigurationLogConfiguration `json:"logConfiguration,omitempty" yaml:"logConfiguration,omitempty"`

	// The log setting to use for redirecting logs for your execute command results. Valid values are `NONE`, `DEFAULT`, and `OVERRIDE`.
	Logging string `json:"logging,omitempty" yaml:"logging,omitempty"`
}
