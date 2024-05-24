package types

type Route53_ResolverEndpointIpAddress struct {
	// The IP address in the subnet that you want to use for DNS queries.
	Ip string `json:"ip,omitempty" yaml:"ip,omitempty"`

	//
	IpId string `json:"ipId,omitempty" yaml:"ipId,omitempty"`

	// The ID of the subnet that contains the IP address.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`
}
