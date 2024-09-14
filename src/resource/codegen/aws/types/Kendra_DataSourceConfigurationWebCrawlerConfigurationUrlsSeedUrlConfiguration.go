package types

type Kendra_DataSourceConfigurationWebCrawlerConfigurationUrlsSeedUrlConfiguration struct {
	// The list of seed or starting point URLs of the websites you want to crawl. The list can include a maximum of `100` seed URLs. Array Members: Minimum number of `0` items. Maximum number of `100` items. Length Constraints: Minimum length of `1`. Maximum length of `2048`.
	SeedUrls []string `json:"seedUrls,omitempty" yaml:"seedUrls,omitempty"`

	/*
	   The default mode is set to `HOST_ONLY`. You can choose one of the following modes:
	   - `HOST_ONLY` – crawl only the website host names. For example, if the seed URL is `"abc.example.com"`, then only URLs with host name `"abc.example.com"` are crawled.
	   - `SUBDOMAINS` – crawl the website host names with subdomains. For example, if the seed URL is `"abc.example.com"`, then `"a.abc.example.com"` and `"b.abc.example.com"` are also crawled.
	   - `EVERYTHING` – crawl the website host names with subdomains and other domains that the webpages link to.
	*/
	WebCrawlerMode string `json:"webCrawlerMode,omitempty" yaml:"webCrawlerMode,omitempty"`
}
