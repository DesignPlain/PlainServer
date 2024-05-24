package types

type Ec2_getInstancePrivateDnsNameOption struct {
	// Indicates whether to respond to DNS queries for instance hostnames with DNS A records.
	EnableResourceNameDnsARecord bool `json:"enableResourceNameDnsARecord,omitempty" yaml:"enableResourceNameDnsARecord,omitempty"`

	// Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records.
	EnableResourceNameDnsAaaaRecord bool `json:"enableResourceNameDnsAaaaRecord,omitempty" yaml:"enableResourceNameDnsAaaaRecord,omitempty"`

	// Type of hostname for EC2 instances.
	HostnameType string `json:"hostnameType,omitempty" yaml:"hostnameType,omitempty"`
}
