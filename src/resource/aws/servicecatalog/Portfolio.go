package servicecatalog

type Portfolio struct {
	// Name of the person or organization who owns the portfolio.
	ProviderName string `json:"providerName,omitempty" yaml:"providerName,omitempty"`

	// Tags to apply to the connection. If configured with a provider `default_tags` configuration block present, tags with matching keys will overwrite those defined at the provider-level.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`

	// Description of the portfolio
	Description string `json:"description,omitempty" yaml:"description,omitempty"`

	// The name of the portfolio.
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
}
