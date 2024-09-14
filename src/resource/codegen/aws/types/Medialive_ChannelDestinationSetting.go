package types

type Medialive_ChannelDestinationSetting struct {
	// A URL specifying a destination.
	Url string `json:"url,omitempty" yaml:"url,omitempty"`

	// Username for destination.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`

	// Key used to extract the password from EC2 Parameter store.
	PasswordParam string `json:"passwordParam,omitempty" yaml:"passwordParam,omitempty"`

	// Stream name RTMP destinations (URLs of type rtmp://)
	StreamName string `json:"streamName,omitempty" yaml:"streamName,omitempty"`
}
