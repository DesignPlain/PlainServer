package types

type Ecs_ServiceServiceConnectConfigurationLogConfigurationSecretOption struct {
	// Name of the secret.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Secret to expose to the container. The supported values are either the full ARN of the AWS Secrets Manager secret or the full ARN of the parameter in the SSM Parameter Store.
	ValueFrom string `json:"valueFrom,omitempty" yaml:"valueFrom,omitempty"`
}
