package types

type Networkservices_TlsRouteRuleActionDestination struct {
	/*
	   Specifies the proportion of requests forwarded to the backend referenced by the serviceName field.

	   - - -
	*/
	Weight int `json:"weight,omitempty" yaml:"weight,omitempty"`

	// The URL of a BackendService to route traffic to.
	ServiceName string `json:"serviceName,omitempty" yaml:"serviceName,omitempty"`
}
