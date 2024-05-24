package types

type Route53recoveryreadiness_ResourceSetResource struct {
	//
	ComponentId string `json:"componentId,omitempty" yaml:"componentId,omitempty"`

	// Component for DNS/Routing Control Readiness Checks.
	DnsTargetResource Route53recoveryreadiness_ResourceSetResourceDnsTargetResource `json:"dnsTargetResource,omitempty" yaml:"dnsTargetResource,omitempty"`

	// Recovery group ARN or cell ARN that contains this resource set.
	ReadinessScopes []string `json:"readinessScopes,omitempty" yaml:"readinessScopes,omitempty"`

	// ARN of the resource.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`
}
