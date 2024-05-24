package types

type Directoryservice_SharedDirectoryTarget struct {
	// Identifier of the directory consumer account.
	Id string `json:"id,omitempty" yaml:"id,omitempty"`

	// Type of identifier to be used in the `id` field. Valid value is `ACCOUNT`. Default is `ACCOUNT`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
