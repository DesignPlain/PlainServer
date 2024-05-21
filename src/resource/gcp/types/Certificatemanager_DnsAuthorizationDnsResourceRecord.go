package types

type Certificatemanager_DnsAuthorizationDnsResourceRecord struct {
	/*
	   (Output)
	   Data of the DNS Resource Record.
	*/
	Data string `json:"data,omitempty" yaml:"data,omitempty"`

	/*
	   Name of the resource; provided by the client when the resource is created.
	   The name must be 1-64 characters long, and match the regular expression [a-zA-Z][a-zA-Z0-9_-]- which means the first character must be a letter,
	   and all following characters must be a dash, underscore, letter or digit.


	   - - -
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   (Output)
	   Type of the DNS Resource Record.
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
