package types

type Route53_getTrafficPolicyDocumentRuleSecondary struct {
	//
	RuleReference string `json:"ruleReference,omitempty" yaml:"ruleReference,omitempty"`

	//
	EndpointReference string `json:"endpointReference,omitempty" yaml:"endpointReference,omitempty"`

	//
	EvaluateTargetHealth bool `json:"evaluateTargetHealth,omitempty" yaml:"evaluateTargetHealth,omitempty"`

	//
	HealthCheck string `json:"healthCheck,omitempty" yaml:"healthCheck,omitempty"`
}
