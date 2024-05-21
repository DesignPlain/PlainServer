package types

type Networkservices_EdgeCacheServiceRoutingPathMatcherRouteRuleHeaderAction struct {
	/*
	   Describes a header to add.
	   Structure is documented below.
	*/
	RequestHeaderToAdds []Networkservices_EdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionRequestHeaderToAdd `json:"requestHeaderToAdds,omitempty" yaml:"requestHeaderToAdds,omitempty"`

	/*
	   A list of header names for headers that need to be removed from the request prior to forwarding the request to the origin.
	   Structure is documented below.
	*/
	RequestHeaderToRemoves []Networkservices_EdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionRequestHeaderToRemove `json:"requestHeaderToRemoves,omitempty" yaml:"requestHeaderToRemoves,omitempty"`

	/*
	   Headers to add to the response prior to sending it back to the client.
	   Response headers are only sent to the client, and do not have an effect on the cache serving the response.
	   Structure is documented below.
	*/
	ResponseHeaderToAdds []Networkservices_EdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionResponseHeaderToAdd `json:"responseHeaderToAdds,omitempty" yaml:"responseHeaderToAdds,omitempty"`

	/*
	   A list of header names for headers that need to be removed from the request prior to forwarding the request to the origin.
	   Structure is documented below.
	*/
	ResponseHeaderToRemoves []Networkservices_EdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionResponseHeaderToRemove `json:"responseHeaderToRemoves,omitempty" yaml:"responseHeaderToRemoves,omitempty"`
}
