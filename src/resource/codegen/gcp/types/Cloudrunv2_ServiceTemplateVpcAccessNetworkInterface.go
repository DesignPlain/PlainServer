package types

type Cloudrunv2_ServiceTemplateVpcAccessNetworkInterface struct {
	/*
	   The VPC network that the Cloud Run resource will be able to send traffic to. At least one of network or subnetwork must be specified. If both
	   network and subnetwork are specified, the given VPC subnetwork must belong to the given VPC network. If network is not specified, it will be
	   looked up from the subnetwork.
	*/
	Network string `json:"network,omitempty" yaml:"network,omitempty"`

	/*
	   The VPC subnetwork that the Cloud Run resource will get IPs from. At least one of network or subnetwork must be specified. If both
	   network and subnetwork are specified, the given VPC subnetwork must belong to the given VPC network. If subnetwork is not specified, the
	   subnetwork with the same name with the network will be used.
	*/
	Subnetwork string `json:"subnetwork,omitempty" yaml:"subnetwork,omitempty"`

	// Network tags applied to this Cloud Run service.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
