package types

type Route53_getTrafficPolicyDocumentRuleGeoProximityLocation struct {
	// Represents the location west (negative) or east (positive) of the prime meridian. Valid values are -180 degrees to 180 degrees.
	Longitude string `json:"longitude,omitempty" yaml:"longitude,omitempty"`

	// If your endpoint is an AWS resource, specify the AWS Region that you created the resource in.
	Region string `json:"region,omitempty" yaml:"region,omitempty"`

	// References to a rule.
	RuleReference string `json:"ruleReference,omitempty" yaml:"ruleReference,omitempty"`

	// Specify a value for `bias` if you want to route more traffic to an endpoint from nearby endpoints (positive values) or route less traffic to an endpoint (negative values).
	Bias string `json:"bias,omitempty" yaml:"bias,omitempty"`

	// References to an endpoint.
	EndpointReference string `json:"endpointReference,omitempty" yaml:"endpointReference,omitempty"`

	// Indicates whether you want Amazon Route 53 to evaluate the health of the endpoint and route traffic only to healthy endpoints.
	EvaluateTargetHealth bool `json:"evaluateTargetHealth,omitempty" yaml:"evaluateTargetHealth,omitempty"`

	// If you want to associate a health check with the endpoint or rule.
	HealthCheck string `json:"healthCheck,omitempty" yaml:"healthCheck,omitempty"`

	// Represents the location south (negative) or north (positive) of the equator. Valid values are -90 degrees to 90 degrees.
	Latitude string `json:"latitude,omitempty" yaml:"latitude,omitempty"`
}
