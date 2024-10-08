package medialive

import types "libds/aws/types"

type Channel struct {
	// Settings for the VPC outputs. See VPC for more details.
	Vpc types.Medialive_ChannelVpc `json:"vpc,omitempty" yaml:"vpc,omitempty"`

	// Specification of CDI inputs for this channel. See CDI Input Specification for more details.
	CdiInputSpecification types.Medialive_ChannelCdiInputSpecification `json:"cdiInputSpecification,omitempty" yaml:"cdiInputSpecification,omitempty"`

	// Destinations for channel. See Destinations for more details.
	Destinations []types.Medialive_ChannelDestination `json:"destinations,omitempty" yaml:"destinations,omitempty"`

	// Input attachments for the channel. See Input Attachments for more details.
	InputAttachments []types.Medialive_ChannelInputAttachment `json:"inputAttachments,omitempty" yaml:"inputAttachments,omitempty"`

	// Specification of network and file inputs for the channel.
	InputSpecification types.Medialive_ChannelInputSpecification `json:"inputSpecification,omitempty" yaml:"inputSpecification,omitempty"`

	// Maintenance settings for this channel. See Maintenance for more details.
	Maintenance types.Medialive_ChannelMaintenance `json:"maintenance,omitempty" yaml:"maintenance,omitempty"`

	// Concise argument description.
	RoleArn string `json:"roleArn,omitempty" yaml:"roleArn,omitempty"`

	// A map of tags to assign to the channel. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Concise argument description.
	ChannelClass string `json:"channelClass,omitempty" yaml:"channelClass,omitempty"`

	// Encoder settings. See Encoder Settings for more details.
	EncoderSettings types.Medialive_ChannelEncoderSettings `json:"encoderSettings,omitempty" yaml:"encoderSettings,omitempty"`

	// The log level to write to Cloudwatch logs.
	LogLevel string `json:"logLevel,omitempty" yaml:"logLevel,omitempty"`

	/*
	   Name of the Channel.

	   The following arguments are optional:
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Whether to start/stop channel. Default: `false`
	StartChannel bool `json:"startChannel,omitempty" yaml:"startChannel,omitempty"`
}
