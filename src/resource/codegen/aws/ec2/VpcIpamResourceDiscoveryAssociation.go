package ec2

type VpcIpamResourceDiscoveryAssociation struct {
	// The ID of the IPAM to associate.
	IpamId string `json:"ipamId,omitempty" yaml:"ipamId,omitempty"`

	// The ID of the Resource Discovery to associate.
	IpamResourceDiscoveryId string `json:"ipamResourceDiscoveryId,omitempty" yaml:"ipamResourceDiscoveryId,omitempty"`

	// A map of tags to add to the IPAM resource discovery association resource.
	Tags map[string]string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
