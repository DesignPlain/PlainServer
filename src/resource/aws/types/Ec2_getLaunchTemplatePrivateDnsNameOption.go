package types

type Ec2_getLaunchTemplatePrivateDnsNameOption struct {
	//
	EnableResourceNameDnsAaaaRecord bool `json:"enableResourceNameDnsAaaaRecord,omitempty" yaml:"enableResourceNameDnsAaaaRecord,omitempty"`

	//
	HostnameType string `json:"hostnameType,omitempty" yaml:"hostnameType,omitempty"`

	//
	EnableResourceNameDnsARecord bool `json:"enableResourceNameDnsARecord,omitempty" yaml:"enableResourceNameDnsARecord,omitempty"`
}
