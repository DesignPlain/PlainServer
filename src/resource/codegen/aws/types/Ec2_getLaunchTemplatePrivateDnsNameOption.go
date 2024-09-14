package types

type Ec2_getLaunchTemplatePrivateDnsNameOption struct {
	//
	EnableResourceNameDnsARecord bool `json:"enableResourceNameDnsARecord,omitempty" yaml:"enableResourceNameDnsARecord,omitempty"`

	//
	EnableResourceNameDnsAaaaRecord bool `json:"enableResourceNameDnsAaaaRecord,omitempty" yaml:"enableResourceNameDnsAaaaRecord,omitempty"`

	//
	HostnameType string `json:"hostnameType,omitempty" yaml:"hostnameType,omitempty"`
}
