package types

type Quicksight_DataSourceCredentialsCredentialPair struct {
	// Password, maximum length of 1024 characters.
	Password string `json:"password,omitempty" yaml:"password,omitempty"`

	// User name, maximum length of 64 characters.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`
}
