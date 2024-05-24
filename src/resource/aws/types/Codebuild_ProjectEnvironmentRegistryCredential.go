package types

type Codebuild_ProjectEnvironmentRegistryCredential struct {
	// ARN or name of credentials created using AWS Secrets Manager.
	Credential string `json:"credential,omitempty" yaml:"credential,omitempty"`

	// Service that created the credentials to access a private Docker registry. Valid value: `SECRETS_MANAGER` (AWS Secrets Manager).
	CredentialProvider string `json:"credentialProvider,omitempty" yaml:"credentialProvider,omitempty"`
}
