package types



type Compute_RegionUrlMapPathMatcherRouteRuleHeaderAction struct {
	/*
	   Headers to add to a matching request before forwarding the request to the backendService.
	   Structure is documented below.
	*/
	RequestHeadersToAdds []Compute_RegionUrlMapPathMatcherRouteRuleHeaderActionRequestHeadersToAdd `json:"requestHeadersToAdds,omitempty" yaml:"requestHeadersToAdds,omitempty"`

	// A list of header names for headers that need to be removed from the request before forwarding the request to the backendService.
	RequestHeadersToRemoves []string `json:"requestHeadersToRemoves,omitempty" yaml:"requestHeadersToRemoves,omitempty"`

	/*
	   Headers to add the response before sending the response back to the client.
	   Structure is documented below.
	*/
	ResponseHeadersToAdds []Compute_RegionUrlMapPathMatcherRouteRuleHeaderActionResponseHeadersToAdd `json:"responseHeadersToAdds,omitempty" yaml:"responseHeadersToAdds,omitempty"`

	// A list of header names for headers that need to be removed from the response before sending the response back to the client.
	ResponseHeadersToRemoves []string `json:"responseHeadersToRemoves,omitempty" yaml:"responseHeadersToRemoves,omitempty"`
}
