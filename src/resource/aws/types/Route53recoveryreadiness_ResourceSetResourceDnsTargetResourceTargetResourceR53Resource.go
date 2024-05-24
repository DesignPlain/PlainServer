package types

type Route53recoveryreadiness_ResourceSetResourceDnsTargetResourceTargetResourceR53Resource struct {
	// DNS Name that acts as the ingress point to a portion of application.
	DomainName string `json:"domainName,omitempty" yaml:"domainName,omitempty"`

	// Route53 record set id to uniquely identify a record given a `domain_name` and a `record_type`.
	RecordSetId string `json:"recordSetId,omitempty" yaml:"recordSetId,omitempty"`
}
