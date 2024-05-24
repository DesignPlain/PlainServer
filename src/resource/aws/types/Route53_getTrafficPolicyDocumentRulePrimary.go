package types

type Route53_getTrafficPolicyDocumentRulePrimary struct {
	// References to an endpoint.
	EndpointReference string `json:"endpointReference,omitempty" yaml:"endpointReference,omitempty"`

	// Indicates whether you want Amazon Route 53 to evaluate the health of the endpoint and route traffic only to healthy endpoints.
	EvaluateTargetHealth bool `json:"evaluateTargetHealth,omitempty" yaml:"evaluateTargetHealth,omitempty"`

	// If you want to associate a health check with the endpoint or rule.
	HealthCheck string `json:"healthCheck,omitempty" yaml:"healthCheck,omitempty"`

	// References to a rule.
	RuleReference string `json:"ruleReference,omitempty" yaml:"ruleReference,omitempty"`
}
