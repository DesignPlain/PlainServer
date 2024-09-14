package types

type Kendra_DataSourceConfigurationWebCrawlerConfigurationAuthenticationConfigurationBasicAuthentication struct {
	// The name of the website host you want to connect to using authentication credentials. For example, the host name of `https://a.example.com/page1.html` is `"a.example.com"`.
	Host string `json:"host,omitempty" yaml:"host,omitempty"`

	// The port number of the website host you want to connect to using authentication credentials. For example, the port for `https://a.example.com/page1.html` is `443`, the standard port for HTTPS.
	Port int `json:"port,omitempty" yaml:"port,omitempty"`

	// Your secret ARN, which you can create in AWS Secrets Manager. You use a secret if basic authentication credentials are required to connect to a website. The secret stores your credentials of user name and password.
	Credentials string `json:"credentials,omitempty" yaml:"credentials,omitempty"`
}
