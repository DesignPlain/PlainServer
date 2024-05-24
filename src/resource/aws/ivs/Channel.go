package ivs

type Channel struct {
	// If `true`, channel is private (enabled for playback authorization).
	Authorized bool `json:"authorized,omitempty" yaml:"authorized,omitempty"`

	// Channel latency mode. Valid values: `NORMAL`, `LOW`.
	LatencyMode string `json:"latencyMode,omitempty" yaml:"latencyMode,omitempty"`

	// Channel name.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	// Recording configuration ARN.
	RecordingConfigurationArn string `json:"recordingConfigurationArn,omitempty" yaml:"recordingConfigurationArn,omitempty"`

	// A map of tags to assign to the resource. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Channel type, which determines the allowable resolution and bitrate. Valid values: `STANDARD`, `BASIC`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}
