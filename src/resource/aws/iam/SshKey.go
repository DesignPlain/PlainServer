package iam

type SshKey struct {
	// Specifies the public key encoding format to use in the response. To retrieve the public key in ssh-rsa format, use `SSH`. To retrieve the public key in PEM format, use `PEM`.
	Encoding string `json:"encoding,omitempty" yaml:"encoding,omitempty"`

	// The SSH public key. The public key must be encoded in ssh-rsa format or PEM format.
	PublicKey string `json:"publicKey,omitempty" yaml:"publicKey,omitempty"`

	// The status to assign to the SSH public key. Active means the key can be used for authentication with an AWS CodeCommit repository. Inactive means the key cannot be used. Default is `active`.
	Status string `json:"status,omitempty" yaml:"status,omitempty"`

	// The name of the IAM user to associate the SSH public key with.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
