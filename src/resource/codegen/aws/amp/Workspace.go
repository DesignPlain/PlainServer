package amp

import types "libds/aws/types"

type Workspace struct {
	// Logging configuration for the workspace. See Logging Configuration below for details.
	LoggingConfiguration types.Amp_WorkspaceLoggingConfiguration `json:"loggingConfiguration,omitempty" yaml:"loggingConfiguration,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// The alias of the prometheus workspace. See more [in AWS Docs](https://docs.aws.amazon.com/prometheus/latest/userguide/AMP-onboard-create-workspace.html).
	Alias string `json:"alias,omitempty" yaml:"alias,omitempty"`

	// The ARN for the KMS encryption key. If this argument is not provided, then the AWS owned encryption key will be used to encrypt the data in the workspace. See more [in AWS Docs](https://docs.aws.amazon.com/prometheus/latest/userguide/encryption-at-rest-Amazon-Service-Prometheus.html)
	KmsKeyArn string `json:"kmsKeyArn,omitempty" yaml:"kmsKeyArn,omitempty"`
}
