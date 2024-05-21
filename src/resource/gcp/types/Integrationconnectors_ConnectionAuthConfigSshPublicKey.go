package types

type Integrationconnectors_ConnectionAuthConfigSshPublicKey struct {
	// Format of SSH Client cert.
	CertType string `json:"certType,omitempty" yaml:"certType,omitempty"`

	/*
	   SSH Client Cert. It should contain both public and private key.
	   Structure is documented below.
	*/
	SshClientCert Integrationconnectors_ConnectionAuthConfigSshPublicKeySshClientCert `json:"sshClientCert,omitempty" yaml:"sshClientCert,omitempty"`

	/*
	   Password (passphrase) for ssh client certificate if it has one.
	   Structure is documented below.
	*/
	SshClientCertPass Integrationconnectors_ConnectionAuthConfigSshPublicKeySshClientCertPass `json:"sshClientCertPass,omitempty" yaml:"sshClientCertPass,omitempty"`

	// The user account used to authenticate.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
