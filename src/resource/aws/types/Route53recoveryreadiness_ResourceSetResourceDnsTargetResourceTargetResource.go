package types

type Route53recoveryreadiness_ResourceSetResourceDnsTargetResourceTargetResource struct {
	// NLB resource a DNS Target Resource points to. Required if `r53_resource` is not set.
	NlbResource Route53recoveryreadiness_ResourceSetResourceDnsTargetResourceTargetResourceNlbResource `json:"nlbResource,omitempty" yaml:"nlbResource,omitempty"`

	// Route53 resource a DNS Target Resource record points to.
	R53Resource Route53recoveryreadiness_ResourceSetResourceDnsTargetResourceTargetResourceR53Resource `json:"r53Resource,omitempty" yaml:"r53Resource,omitempty"`
}
