package types

type Networkservices_HttpRouteRuleAction struct {
	/*
	   Specifies the retry policy associated with this route.
	   Structure is documented below.
	*/
	RetryPolicy Networkservices_HttpRouteRuleActionRetryPolicy `json:"retryPolicy,omitempty" yaml:"retryPolicy,omitempty"`

	// Specifies the timeout for selected route.
	Timeout string `json:"timeout,omitempty" yaml:"timeout,omitempty"`

	/*
	   The specification for allowing client side cross-origin requests.
	   Structure is documented below.
	*/
	CorsPolicy Networkservices_HttpRouteRuleActionCorsPolicy `json:"corsPolicy,omitempty" yaml:"corsPolicy,omitempty"`

	/*
	   The destination to which traffic should be forwarded.
	   Structure is documented below.
	*/
	Destinations []Networkservices_HttpRouteRuleActionDestination `json:"destinations,omitempty" yaml:"destinations,omitempty"`

	/*
	   The specification for fault injection introduced into traffic to test the resiliency of clients to backend service failure.
	   Structure is documented below.
	*/
	FaultInjectionPolicy Networkservices_HttpRouteRuleActionFaultInjectionPolicy `json:"faultInjectionPolicy,omitempty" yaml:"faultInjectionPolicy,omitempty"`

	/*
	   If set, the request is directed as configured by this field.
	   Structure is documented below.
	*/
	Redirect Networkservices_HttpRouteRuleActionRedirect `json:"redirect,omitempty" yaml:"redirect,omitempty"`

	/*
	   The specification for modifying the headers of a matching request prior to delivery of the request to the destination.
	   Structure is documented below.
	*/
	RequestHeaderModifier Networkservices_HttpRouteRuleActionRequestHeaderModifier `json:"requestHeaderModifier,omitempty" yaml:"requestHeaderModifier,omitempty"`

	/*
	   Specifies the policy on how requests intended for the routes destination are shadowed to a separate mirrored destination.
	   Structure is documented below.
	*/
	RequestMirrorPolicy Networkservices_HttpRouteRuleActionRequestMirrorPolicy `json:"requestMirrorPolicy,omitempty" yaml:"requestMirrorPolicy,omitempty"`

	/*
	   The specification for modifying the headers of a response prior to sending the response back to the client.
	   Structure is documented below.
	*/
	ResponseHeaderModifier Networkservices_HttpRouteRuleActionResponseHeaderModifier `json:"responseHeaderModifier,omitempty" yaml:"responseHeaderModifier,omitempty"`

	/*
	   The specification for rewrite URL before forwarding requests to the destination.
	   Structure is documented below.
	*/
	UrlRewrite Networkservices_HttpRouteRuleActionUrlRewrite `json:"urlRewrite,omitempty" yaml:"urlRewrite,omitempty"`
}
