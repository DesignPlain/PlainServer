package types

type Route53_getTrafficPolicyDocumentRuleLocation struct {
	// Value of a subdivision.
	Subdivision string `json:"subdivision,omitempty" yaml:"subdivision,omitempty"`

	// Value of a continent.
	Continent string `json:"continent,omitempty" yaml:"continent,omitempty"`

	// Value of a country.
	Country string `json:"country,omitempty" yaml:"country,omitempty"`

	// References to an endpoint.
	EndpointReference string `json:"endpointReference,omitempty" yaml:"endpointReference,omitempty"`

	// Indicates whether you want Amazon Route 53 to evaluate the health of the endpoint and route traffic only to healthy endpoints.
	EvaluateTargetHealth bool `json:"evaluateTargetHealth,omitempty" yaml:"evaluateTargetHealth,omitempty"`

	// If you want to associate a health check with the endpoint or rule.
	HealthCheck string `json:"healthCheck,omitempty" yaml:"healthCheck,omitempty"`

	// Indicates whether this set of values represents the default location.
	IsDefault bool `json:"isDefault,omitempty" yaml:"isDefault,omitempty"`

	// References to a rule.
	RuleReference string `json:"ruleReference,omitempty" yaml:"ruleReference,omitempty"`
}
