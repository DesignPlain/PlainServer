package types

type Dynamodb_TableTtl struct {
	/*
	   Name of the table attribute to store the TTL timestamp in.
	   Required if `enabled` is `true`, must not be set otherwise.
	*/
	AttributeName string `json:"attributeName,omitempty" yaml:"attributeName,omitempty"`

	/*
	   Whether TTL is enabled.
	   Default value is `false`.
	*/
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
