package types

type Dns_ManagedZoneCloudLoggingConfig struct {
	// If set, enable query logging for this ManagedZone. False by default, making logging opt-in.
	EnableLogging bool `json:"enableLogging,omitempty" yaml:"enableLogging,omitempty"`
}
