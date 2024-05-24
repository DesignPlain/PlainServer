package pinpoint

type SmsChannel struct {
	// ID of the application.
	ApplicationId string `json:"applicationId,omitempty" yaml:"applicationId,omitempty"`

	// Whether the channel is enabled or disabled. By default, it is set to `true`.
	Enabled bool `json:"enabled,omitempty" yaml:"enabled,omitempty"`

	// Identifier of the sender for your messages.
	SenderId string `json:"senderId,omitempty" yaml:"senderId,omitempty"`

	// Short Code registered with the phone provider.
	ShortCode string `json:"shortCode,omitempty" yaml:"shortCode,omitempty"`
}
