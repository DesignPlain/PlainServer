package types

type Gkeonprem_BareMetalClusterProxy struct {
	/*
	   A list of IPs, hostnames, and domains that should skip the proxy.
	   Examples: ["127.0.0.1", "example.com", ".corp", "localhost"].
	*/
	NoProxies []string `json:"noProxies,omitempty" yaml:"noProxies,omitempty"`

	/*
	   Specifies the address of your proxy server.
	   Examples: http://domain
	   WARNING: Do not provide credentials in the format
	   http://(username:password@)domain these will be rejected by the server.
	*/
	Uri string `json:"uri,omitempty" yaml:"uri,omitempty"`
}
