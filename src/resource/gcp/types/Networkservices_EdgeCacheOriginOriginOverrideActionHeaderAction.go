package types

type Networkservices_EdgeCacheOriginOriginOverrideActionHeaderAction struct {
	/*
	   Describes a header to add.
	   You may add a maximum of 25 request headers.
	   Structure is documented below.
	*/
	RequestHeadersToAdds []Networkservices_EdgeCacheOriginOriginOverrideActionHeaderActionRequestHeadersToAdd `json:"requestHeadersToAdds,omitempty" yaml:"requestHeadersToAdds,omitempty"`
}
