package controltower

type LandingZone struct {
	// The landing zone version.
	Version string `json:"version,omitempty" yaml:"version,omitempty"`

	// The manifest JSON file is a text file that describes your AWS resources. For examples, review [Launch your landing zone](https://docs.aws.amazon.com/controltower/latest/userguide/lz-api-launch).
	ManifestJson string `json:"manifestJson,omitempty" yaml:"manifestJson,omitempty"`

	// Tags to apply to the landing zone. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
