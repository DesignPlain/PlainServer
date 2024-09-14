package types

type Sesv2_getDedicatedIpPoolDedicatedIp struct {
	// IPv4 address.
	Ip string `json:"ip,omitempty" yaml:"ip,omitempty"`

	// Indicates how complete the dedicated IP warm-up process is. When this value equals `1`, the address has completed the warm-up process and is ready for use.
	WarmupPercentage int `json:"warmupPercentage,omitempty" yaml:"warmupPercentage,omitempty"`

	// The warm-up status of a dedicated IP address. Valid values: `IN_PROGRESS`, `DONE`.
	WarmupStatus string `json:"warmupStatus,omitempty" yaml:"warmupStatus,omitempty"`
}
