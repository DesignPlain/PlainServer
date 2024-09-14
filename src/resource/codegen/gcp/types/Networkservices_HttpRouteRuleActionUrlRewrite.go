package types

type Networkservices_HttpRouteRuleActionUrlRewrite struct {
	// Prior to forwarding the request to the selected destination, the requests host header is replaced by this value.
	HostRewrite string `json:"hostRewrite,omitempty" yaml:"hostRewrite,omitempty"`

	// Prior to forwarding the request to the selected destination, the matching portion of the requests path is replaced by this value.
	PathPrefixRewrite string `json:"pathPrefixRewrite,omitempty" yaml:"pathPrefixRewrite,omitempty"`
}
