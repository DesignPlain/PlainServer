package types

type Route53recoveryreadiness_ResourceSetResourceDnsTargetResource struct {
	// Target resource the R53 record specified with the above params points to.
	TargetResource Route53recoveryreadiness_ResourceSetResourceDnsTargetResourceTargetResource `json:"targetResource,omitempty" yaml:"targetResource,omitempty"`

	// DNS Name that acts as the ingress point to a portion of application.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// Hosted Zone ARN that contains the DNS record with the provided name of target resource.
	HostedZoneArn string `json:"hostedZoneArn,omitempty" yaml:"hostedZoneArn,omitempty"`

	// Route53 record set id to uniquely identify a record given a `domain_name` and a `record_type`.
	RecordSetId string `json:"recordSetId,omitempty" yaml:"recordSetId,omitempty"`

	// Type of DNS Record of target resource.
	RecordType string `json:"recordType,omitempty" yaml:"recordType,omitempty"`
}
