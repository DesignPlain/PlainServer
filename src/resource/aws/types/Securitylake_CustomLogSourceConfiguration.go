package types

type Securitylake_CustomLogSourceConfiguration struct {
	// The configuration for the Glue Crawler for the third-party custom source.
	CrawlerConfiguration Securitylake_CustomLogSourceConfigurationCrawlerConfiguration `json:"crawlerConfiguration,omitempty" yaml:"crawlerConfiguration,omitempty"`

	// The identity of the log provider for the third-party custom source.
	ProviderIdentity Securitylake_CustomLogSourceConfigurationProviderIdentity `json:"providerIdentity,omitempty" yaml:"providerIdentity,omitempty"`
}
