package types

type Ec2transitgateway_getMulticastDomainAssociation struct {
	// The ID of the subnet associated with the transit gateway multicast domain.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`

	// The ID of the transit gateway attachment.
	TransitGatewayAttachmentId string `json:"transitGatewayAttachmentId,omitempty" yaml:"transitGatewayAttachmentId,omitempty"`
}
