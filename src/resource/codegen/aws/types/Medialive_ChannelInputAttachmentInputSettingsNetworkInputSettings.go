package types

type Medialive_ChannelInputAttachmentInputSettingsNetworkInputSettings struct {
	// Specifies HLS input settings when the uri is for a HLS manifest. See HLS Input Settings for more details.
	HlsInputSettings Medialive_ChannelInputAttachmentInputSettingsNetworkInputSettingsHlsInputSettings `json:"hlsInputSettings,omitempty" yaml:"hlsInputSettings,omitempty"`

	// Check HTTPS server certificates.
	ServerValidation string `json:"serverValidation,omitempty" yaml:"serverValidation,omitempty"`
}
