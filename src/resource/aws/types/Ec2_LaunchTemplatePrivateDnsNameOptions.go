package types

type Ec2_LaunchTemplatePrivateDnsNameOptions struct {
	// Indicates whether to respond to DNS queries for instance hostnames with DNS A records.
	EnableResourceNameDnsARecord bool `json:"enableResourceNameDnsARecord,omitempty" yaml:"enableResourceNameDnsARecord,omitempty"`

	// Indicates whether to respond to DNS queries for instance hostnames with DNS AAAA records.
	EnableResourceNameDnsAaaaRecord bool `json:"enableResourceNameDnsAaaaRecord,omitempty" yaml:"enableResourceNameDnsAaaaRecord,omitempty"`

	// The type of hostname for Amazon EC2 instances. For IPv4 only subnets, an instance DNS name must be based on the instance IPv4 address. For IPv6 native subnets, an instance DNS name must be based on the instance ID. For dual-stack subnets, you can specify whether DNS names use the instance IPv4 address or the instance ID. Valid values: `ip-name` and `resource-name`.
	HostnameType string `json:"hostnameType,omitempty" yaml:"hostnameType,omitempty"`
}
