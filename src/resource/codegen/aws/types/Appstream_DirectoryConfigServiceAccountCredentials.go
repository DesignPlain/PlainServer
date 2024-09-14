package types

type Appstream_DirectoryConfigServiceAccountCredentials struct {
	// Password for the account.
	AccountPassword string `json:"accountPassword,omitempty" yaml:"accountPassword,omitempty"`

	// User name of the account. This account must have the following privileges: create computer objects, join computers to the domain, and change/reset the password on descendant computer objects for the organizational units specified.
	AccountName string `json:"accountName,omitempty" yaml:"accountName,omitempty"`
}
