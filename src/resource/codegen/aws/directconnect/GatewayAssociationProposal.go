package directconnect

type GatewayAssociationProposal struct {
	// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
	AllowedPrefixes []string `json:"allowedPrefixes,omitempty" yaml:"allowedPrefixes,omitempty"`

	// The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
	AssociatedGatewayId string `json:"associatedGatewayId,omitempty" yaml:"associatedGatewayId,omitempty"`

	// Direct Connect Gateway identifier.
	DxGatewayId string `json:"dxGatewayId,omitempty" yaml:"dxGatewayId,omitempty"`

	// AWS Account identifier of the Direct Connect Gateway's owner.
	DxGatewayOwnerAccountId string `json:"dxGatewayOwnerAccountId,omitempty" yaml:"dxGatewayOwnerAccountId,omitempty"`
}
