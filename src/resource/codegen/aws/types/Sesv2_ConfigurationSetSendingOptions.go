package types

type Sesv2_ConfigurationSetSendingOptions struct {
	// If `true`, email sending is enabled for the configuration set. If `false`, email sending is disabled for the configuration set.
	SendingEnabled bool `json:"sendingEnabled,omitempty" yaml:"sendingEnabled,omitempty"`
}
