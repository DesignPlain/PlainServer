package types

type Kendra_DataSourceConfigurationWebCrawlerConfigurationProxyConfiguration struct {
	// Your secret ARN, which you can create in AWS Secrets Manager. The credentials are optional. You use a secret if web proxy credentials are required to connect to a website host. Amazon Kendra currently support basic authentication to connect to a web proxy server. The secret stores your credentials.
	Credentials string `json:"credentials,omitempty" yaml:"credentials,omitempty"`

	// The name of the website host you want to connect to via a web proxy server. For example, the host name of `https://a.example.com/page1.html` is `"a.example.com"`.
	Host string `json:"host,omitempty" yaml:"host,omitempty"`

	// The port number of the website host you want to connect to via a web proxy server. For example, the port for `https://a.example.com/page1.html` is `443`, the standard port for HTTPS.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`
}
