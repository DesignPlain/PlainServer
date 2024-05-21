package types

type Networkservices_EdgeCacheServiceRoutingPathMatcherRouteRuleHeaderActionResponseHeaderToRemove struct {
	/*
	   Headers to remove from the response prior to sending it back to the client.
	   Response headers are only sent to the client, and do not have an effect on the cache serving the response.
	*/
	HeaderName string `json:"headerName,omitempty" yaml:"headerName,omitempty"`
}
