package ivschat

import types "libds/aws/types"

type Room struct {
	/*
	   Maximum number of characters in a single
	   message. Messages are expected to be UTF-8 encoded and this limit applies
	   specifically to rune/code-point count, not number of bytes.
	*/
	MaximumMessageLength int `json:"maximumMessageLength,omitempty" yaml:"maximumMessageLength,omitempty"`

	/*
	   Maximum number of messages per
	   second that can be sent to the room (by all clients).
	*/
	MaximumMessageRatePerSecond int `json:"maximumMessageRatePerSecond,omitempty" yaml:"maximumMessageRatePerSecond,omitempty"`

	/*
	   Configuration information for optional
	   review of messages.
	*/
	MessageReviewHandler types.Ivschat_RoomMessageReviewHandler `json:"messageReviewHandler,omitempty" yaml:"messageReviewHandler,omitempty"`

	// Room name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	/*
	   List of Logging Configuration
	   ARNs to attach to the room.
	*/
	LoggingConfigurationIdentifiers []string `json:"loggingConfigurationIdentifiers,omitempty" yaml:"loggingConfigurationIdentifiers,omitempty"`
}
