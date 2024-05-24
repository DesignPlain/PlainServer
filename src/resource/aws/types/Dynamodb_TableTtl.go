package types

type Dynamodb_TableTtl struct {
	// Name of the table attribute to store the TTL timestamp in.
	AttributeName string `json:"attributeName,omitempty" yaml:"attributeName,omitempty"`

	// Whether TTL is enabled.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
