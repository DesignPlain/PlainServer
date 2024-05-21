package types

type Compute_URLMapPathMatcherRouteRuleHeaderAction struct {
	/*
	   Headers to add to a matching request prior to forwarding the request to the backendService.
	   Structure is documented below.
	*/
	RequestHeadersToAdds []Compute_URLMapPathMatcherRouteRuleHeaderActionRequestHeadersToAdd `json:"requestHeadersToAdds,omitempty" yaml:"requestHeadersToAdds,omitempty"`

	/*
	   A list of header names for headers that need to be removed from the request prior to
	   forwarding the request to the backendService.
	*/
	RequestHeadersToRemoves []string `json:"requestHeadersToRemoves,omitempty" yaml:"requestHeadersToRemoves,omitempty"`

	/*
	   Headers to add the response prior to sending the response back to the client.
	   Structure is documented below.
	*/
	ResponseHeadersToAdds []Compute_URLMapPathMatcherRouteRuleHeaderActionResponseHeadersToAdd `json:"responseHeadersToAdds,omitempty" yaml:"responseHeadersToAdds,omitempty"`

	/*
	   A list of header names for headers that need to be removed from the response prior to sending the
	   response back to the client.
	*/
	ResponseHeadersToRemoves []string `json:"responseHeadersToRemoves,omitempty" yaml:"responseHeadersToRemoves,omitempty"`
}
