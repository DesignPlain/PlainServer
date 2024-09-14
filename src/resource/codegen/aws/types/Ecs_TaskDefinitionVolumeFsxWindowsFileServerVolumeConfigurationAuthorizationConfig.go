package types

type Ecs_TaskDefinitionVolumeFsxWindowsFileServerVolumeConfigurationAuthorizationConfig struct {
	// The authorization credential option to use. The authorization credential options can be provided using either the Amazon Resource Name (ARN) of an AWS Secrets Manager secret or AWS Systems Manager Parameter Store parameter. The ARNs refer to the stored credentials.
	CredentialsParameter string `json:"credentialsParameter,omitempty" yaml:"credentialsParameter,omitempty"`

	// A fully qualified domain name hosted by an AWS Directory Service Managed Microsoft AD (Active Directory) or self-hosted AD on Amazon EC2.
	Domain string `json:"domain,omitempty" yaml:"domain,omitempty"`
}
