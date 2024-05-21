package types

type Firebase_HostingVersionConfig struct {
	/*
	   An array of objects (called rewrite rules), where each rule specifies a URL pattern that, if matched to the
	   request URL path, triggers Hosting to respond as if the service were given the specified destination URL.
	   Structure is documented below.
	*/
	Rewrites []Firebase_HostingVersionConfigRewrite `json:"rewrites,omitempty" yaml:"rewrites,omitempty"`

	/*
	   An array of objects (called redirect rules), where each rule specifies a URL pattern that, if matched to the request URL path,
	   triggers Hosting to respond with a redirect to the specified destination path.
	   Structure is documented below.
	*/
	Redirects []Firebase_HostingVersionConfigRedirect `json:"redirects,omitempty" yaml:"redirects,omitempty"`
}
