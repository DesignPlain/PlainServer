package types

type Networkservices_HttpRouteRuleActionRequestMirrorPolicy struct {
	/*
	   The destination the requests will be mirrored to.
	   Structure is documented below.
	*/
	Destination Networkservices_HttpRouteRuleActionRequestMirrorPolicyDestination `json:"destination,omitempty" yaml:"destination,omitempty"`
}
