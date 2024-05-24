package ec2

type SerialConsoleAccess struct {
	// Whether or not serial console access is enabled. Valid values are `true` or `false`. Defaults to `true`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`
}
