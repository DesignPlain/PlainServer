package types

type Compute_BackendServiceConsistentHash struct {
	/*
	   The minimum number of virtual nodes to use for the hash ring.
	   Larger ring sizes result in more granular load
	   distributions. If the number of hosts in the load balancing pool
	   is larger than the ring size, each host will be assigned a single
	   virtual node.
	   Defaults to 1024.
	*/
	MinimumRingSize int `json:"minimumRingSize,omitempty" yaml:"minimumRingSize,omitempty"`

	/*
	   Hash is based on HTTP Cookie. This field describes a HTTP cookie
	   that will be used as the hash key for the consistent hash load
	   balancer. If the cookie is not present, it will be generated.
	   This field is applicable if the sessionAffinity is set to HTTP_COOKIE.
	   Structure is documented below.
	*/
	HttpCookie Compute_BackendServiceConsistentHashHttpCookie `json:"httpCookie,omitempty" yaml:"httpCookie,omitempty"`

	/*
	   The hash based on the value of the specified header field.
	   This field is applicable if the sessionAffinity is set to HEADER_FIELD.
	*/
	HttpHeaderName string `json:"httpHeaderName,omitempty" yaml:"httpHeaderName,omitempty"`
}
