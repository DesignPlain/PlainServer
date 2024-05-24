package types

type Medialive_ChannelInputAttachment struct {
	// User-specified settings for defining what the conditions are for declaring the input unhealthy and failing over to a different input. See Automatic Input Failover Settings for more details.
	AutomaticInputFailoverSettings Medialive_ChannelInputAttachmentAutomaticInputFailoverSettings `json:"automaticInputFailoverSettings,omitempty" yaml:"automaticInputFailoverSettings,omitempty"`

	// User-specified name for the attachment.
	InputAttachmentName string `json:"inputAttachmentName,omitempty" yaml:"inputAttachmentName,omitempty"`

	// The ID of the input.
	InputId string `json:"inputId,omitempty" yaml:"inputId,omitempty"`

	// Settings of an input. See Input Settings for more details.
	InputSettings Medialive_ChannelInputAttachmentInputSettings `json:"inputSettings,omitempty" yaml:"inputSettings,omitempty"`
}
