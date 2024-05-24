package secretsmanager

import types "DesignSphere_Server/src/resource/aws/types"

type SecretRotation struct {
	// A structure that defines the rotation configuration for this secret. Defined below.
	RotationRules types.Secretsmanager_SecretRotationRotationRules `json:"rotationRules,omitempty" yaml:"rotationRules,omitempty"`

	// Specifies the secret to which you want to add a new version. You can specify either the Amazon Resource Name (ARN) or the friendly name of the secret. The secret must already exist.
	SecretId string `json:"secretId,omitempty" yaml:"secretId,omitempty"`

	// Specifies whether to rotate the secret immediately or wait until the next scheduled rotation window. The rotation schedule is defined in `rotation_rules`. For secrets that use a Lambda rotation function to rotate, if you don't immediately rotate the secret, Secrets Manager tests the rotation configuration by running the testSecret step (https://docs.aws.amazon.com/secretsmanager/latest/userguide/rotate-secrets_how.html) of the Lambda rotation function. The test creates an AWSPENDING version of the secret and then removes it. Defaults to `true`.
	RotateImmediately bool `json:"rotateImmediately,omitempty" yaml:"rotateImmediately,omitempty"`

	// Specifies the ARN of the Lambda function that can rotate the secret. Must be supplied if the secret is not managed by AWS.
	RotationLambdaArn string `json:"rotationLambdaArn,omitempty" yaml:"rotationLambdaArn,omitempty"`
}
