package types

type Container_AwsClusterControlPlaneAwsServicesAuthentication struct {
	// The Amazon Resource Name (ARN) of the role that the Anthos Multi-Cloud API will assume when managing AWS resources on your account.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// Optional. An identifier for the assumed role session. When unspecified, it defaults to `multicloud-service-agent`.
	RoleSessionName string `json:"roleSessionName,omitempty" yaml:"roleSessionName,omitempty"`
}
