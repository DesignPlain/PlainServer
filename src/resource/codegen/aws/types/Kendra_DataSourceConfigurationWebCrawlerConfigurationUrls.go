package types

type Kendra_DataSourceConfigurationWebCrawlerConfigurationUrls struct {
	// A block that specifies the configuration of the seed or starting point URLs of the websites you want to crawl. You can choose to crawl only the website host names, or the website host names with subdomains, or the website host names with subdomains and other domains that the webpages link to. You can list up to `100` seed URLs. Detailed below.
	SeedUrlConfiguration Kendra_DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfiguration `json:"seedUrlConfiguration,omitempty" yaml:"seedUrlConfiguration,omitempty"`

	// A block that specifies the configuration of the sitemap URLs of the websites you want to crawl. Only URLs belonging to the same website host names are crawled. You can list up to `3` sitemap URLs. Detailed below.
	SiteMapsConfiguration Kendra_DataSourceConfigurationWebCrawlerConfigurationUrlsSiteMapsConfiguration `json:"siteMapsConfiguration,omitempty" yaml:"siteMapsConfiguration,omitempty"`
}
