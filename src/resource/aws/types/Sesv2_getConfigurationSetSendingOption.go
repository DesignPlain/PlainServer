package types

type Sesv2_getConfigurationSetSendingOption struct {
	// Specifies whether email sending is enabled.
	SendingEnabled bool `json:"sendingEnabled,omitempty" yaml:"sendingEnabled,omitempty"`
}
