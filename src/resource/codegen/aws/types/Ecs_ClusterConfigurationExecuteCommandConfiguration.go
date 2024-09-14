package types

type Ecs_ClusterConfigurationExecuteCommandConfiguration struct {
	// AWS Key Management Service key ID to encrypt the data between the local client and the container.
	KmsKeyId string `json:"kmsKeyId,omitempty" yaml:"kmsKeyId,omitempty"`

	// Log configuration for the results of the execute command actions. Required when `logging` is `OVERRIDE`. See `log_configuration` Block for details.
	LogConfiguration Ecs_ClusterConfigurationExecuteCommandConfigurationLogConfiguration `json:"logConfiguration,omitempty" yaml:"logConfiguration,omitempty"`

	// Log setting to use for redirecting logs for your execute command results. Valid values: `NONE`, `DEFAULT`, `OVERRIDE`.
	Logging string `json:"logging,omitempty" yaml:"logging,omitempty"`
}
