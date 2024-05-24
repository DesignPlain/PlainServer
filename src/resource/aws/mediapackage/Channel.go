package mediapackage

type Channel struct {
	// A unique identifier describing the channel
	ChannelId string `json:"channelId,omitempty" yaml:"channelId,omitempty"`

	// A description of the channel
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// A map of tags to assign to the resource. .If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
