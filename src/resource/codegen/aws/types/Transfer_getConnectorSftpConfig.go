package types

type Transfer_getConnectorSftpConfig struct {
	// List of the public portions of the host keys that are used to identify the servers the connector is connected to.
	TrustedHostKeys []string `json:"trustedHostKeys,omitempty" yaml:"trustedHostKeys,omitempty"`

	// Identifer for the secret in AWS Secrets Manager that contains the SFTP user's private key, and/or password.
	UserSecretId string `json:"userSecretId,omitempty" yaml:"userSecretId,omitempty"`
}
