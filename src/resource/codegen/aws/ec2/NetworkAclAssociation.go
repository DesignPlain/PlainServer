package ec2

type NetworkAclAssociation struct {
	// The ID of the network ACL.
	NetworkAclId string `json:"networkAclId,omitempty" yaml:"networkAclId,omitempty"`

	// The ID of the associated Subnet.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`
}
