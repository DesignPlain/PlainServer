package ec2transitgateway

type MulticastDomainAssociation struct {
	// The ID of the subnet to associate with the transit gateway multicast domain.
	SubnetId string `json:"subnetId,omitempty" yaml:"subnetId,omitempty"`

	// The ID of the transit gateway attachment.
	TransitGatewayAttachmentId string `json:"transitGatewayAttachmentId,omitempty" yaml:"transitGatewayAttachmentId,omitempty"`

	// The ID of the transit gateway multicast domain.
	TransitGatewayMulticastDomainId string `json:"transitGatewayMulticastDomainId,omitempty" yaml:"transitGatewayMulticastDomainId,omitempty"`
}
