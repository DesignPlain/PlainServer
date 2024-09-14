package types

type Networkservices_GrpcRouteRuleActionDestination struct {
	// The URL of a BackendService to route traffic to.
	ServiceName string `json:"serviceName,omitempty" yaml:"serviceName,omitempty"`

	// Specifies the proportion of requests forwarded to the backend referenced by the serviceName field.
	Weight int `json:"weight,omitempty" yaml:"weight,omitempty"`
}
