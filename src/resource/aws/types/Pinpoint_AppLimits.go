package types

type Pinpoint_AppLimits struct {
	// The maximum number of messages that the campaign can send daily.
	Daily int `json:"daily,omitempty" yaml:"daily,omitempty"`

	// The length of time (in seconds) that the campaign can run before it ends and message deliveries stop. This duration begins at the scheduled start time for the campaign. The minimum value is 60.
	MaximumDuration int `json:"maximumDuration,omitempty" yaml:"maximumDuration,omitempty"`

	// The number of messages that the campaign can send per second. The minimum value is 50, and the maximum is 20000.
	MessagesPerSecond int `json:"messagesPerSecond,omitempty" yaml:"messagesPerSecond,omitempty"`

	// The maximum total number of messages that the campaign can send.
	Total int `json:"total,omitempty" yaml:"total,omitempty"`
}
