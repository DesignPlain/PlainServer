package types

type Medialive_ChannelEncoderSettingsCaptionDescriptionDestinationSettingsDvbSubDestinationSettingsFont struct {
	// Username to be used.
	Username string `json:"username,omitempty" yaml:"username,omitempty"`

	// Key used to extract the password from EC2 Parameter store.
	PasswordParam string `json:"passwordParam,omitempty" yaml:"passwordParam,omitempty"`

	// Path to a file accessible to the live stream.
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`
}
