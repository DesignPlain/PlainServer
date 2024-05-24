package networkmanager

type GlobalNetwork struct {
	// Description of the Global Network.
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// Key-value tags for the Global Network. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
