package types

type Compute_URLMapDefaultRouteActionUrlRewrite struct {
	/*
	   Prior to forwarding the request to the selected backend service, the matching portion of the
	   request's path is replaced by pathPrefixRewrite.
	   The value must be between 1 and 1024 characters.
	*/
	PathPrefixRewrite string `json:"pathPrefixRewrite,omitempty" yaml:"pathPrefixRewrite,omitempty"`

	/*
	   Prior to forwarding the request to the selected service, the request's host header is replaced
	   with contents of hostRewrite.
	   The value must be between 1 and 255 characters.
	*/
	HostRewrite string `json:"hostRewrite,omitempty" yaml:"hostRewrite,omitempty"`
}
