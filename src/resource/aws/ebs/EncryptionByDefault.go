package ebs

type EncryptionByDefault struct {
	// Whether or not default EBS encryption is enabled. Valid values are `true` or `false`. Defaults to `true`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
