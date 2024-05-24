package types

type Quicksight_DataSourceCredentials struct {
	/*
	   The Amazon Resource Name (ARN) of a data source that has the credential pair that you want to use.
	   When the value is not null, the `credential_pair` from the data source in the ARN is used.
	*/
	CopySourceArn string `json:"copySourceArn,omitempty" yaml:"copySourceArn,omitempty"`

	// Credential pair. See Credential Pair below for more details.
	CredentialPair Quicksight_DataSourceCredentialsCredentialPair `json:"credentialPair,omitempty" yaml:"credentialPair,omitempty"`
}
