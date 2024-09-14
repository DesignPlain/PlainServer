package types

type Kendra_DataSourceConfigurationWebCrawlerConfigurationAuthenticationConfiguration struct {
	// The list of configuration information that's required to connect to and crawl a website host using basic authentication credentials. The list includes the name and port number of the website host. Detailed below.
	BasicAuthentications []Kendra_DataSourceConfigurationWebCrawlerConfigurationAuthenticationConfigurationBasicAuthentication `json:"basicAuthentications,omitempty" yaml:"basicAuthentications,omitempty"`
}
