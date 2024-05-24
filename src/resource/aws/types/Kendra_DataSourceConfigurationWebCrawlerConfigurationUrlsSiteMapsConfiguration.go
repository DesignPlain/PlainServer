package types

type Kendra_DataSourceConfigurationWebCrawlerConfigurationUrlsSiteMapsConfiguration struct {
	// The list of sitemap URLs of the websites you want to crawl. The list can include a maximum of `3` sitemap URLs.
	SiteMaps []string `json:"siteMaps,omitempty" yaml:"siteMaps,omitempty"`
}
