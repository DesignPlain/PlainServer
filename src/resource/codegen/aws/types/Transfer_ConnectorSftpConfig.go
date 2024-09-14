package types

type Transfer_ConnectorSftpConfig struct {
	// A list of public portion of the host key, or keys, that are used to authenticate the user to the external server to which you are connecting.(https://docs.aws.amazon.com/transfer/latest/userguide/API_SftpConnectorConfig.html)
	TrustedHostKeys []string `json:"trustedHostKeys,omitempty" yaml:"trustedHostKeys,omitempty"`

	// The identifier for the secret (in AWS Secrets Manager) that contains the SFTP user's private key, password, or both. The identifier can be either the Amazon Resource Name (ARN) or the name of the secret.
	UserSecretId string `json:"userSecretId,omitempty" yaml:"userSecretId,omitempty"`
}
