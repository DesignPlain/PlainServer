package types

type Route53_getTrafficPolicyDocumentRuleItem struct {
	// References to an endpoint.
	EndpointReference string `json:"endpointReference,omitempty" yaml:"endpointReference,omitempty"`

	// If you want to associate a health check with the endpoint or rule.
	HealthCheck string `json:"healthCheck,omitempty" yaml:"healthCheck,omitempty"`
}
