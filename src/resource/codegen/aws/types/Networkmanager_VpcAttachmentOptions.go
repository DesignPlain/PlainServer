package types

type Networkmanager_VpcAttachmentOptions struct {
	/*
	   Indicates whether IPv6 is supported.
	   If the VPC attachment is pending acceptance, changing this value will recreate the resource.
	*/
	Ipv6Support bool `json:"ipv6Support,omitempty" yaml:"ipv6Support,omitempty"`

	/*
	   Indicates whether appliance mode is supported.
	   If enabled, traffic flow between a source and destination use the same Availability Zone for the VPC attachment for the lifetime of that flow.
	   If the VPC attachment is pending acceptance, changing this value will recreate the resource.
	*/
	ApplianceModeSupport bool `json:"applianceModeSupport,omitempty" yaml:"applianceModeSupport,omitempty"`
}
