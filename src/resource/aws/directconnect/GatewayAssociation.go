package directconnect

type GatewayAssociation struct {
	// VPC prefixes (CIDRs) to advertise to the Direct Connect gateway. Defaults to the CIDR block of the VPC associated with the Virtual Gateway. To enable drift detection, must be configured.
	AllowedPrefixes []string `json:"allowedPrefixes,omitempty" yaml:"allowedPrefixes,omitempty"`

	/*
	   The ID of the VGW or transit gateway with which to associate the Direct Connect gateway.
	   Used for single account Direct Connect gateway associations.
	*/
	AssociatedGatewayId string `json:"associatedGatewayId,omitempty" yaml:"associatedGatewayId,omitempty"`

	/*
	   The ID of the AWS account that owns the VGW or transit gateway with which to associate the Direct Connect gateway.
	   Used for cross-account Direct Connect gateway associations.
	*/
	AssociatedGatewayOwnerAccountId string `json:"associatedGatewayOwnerAccountId,omitempty" yaml:"associatedGatewayOwnerAccountId,omitempty"`

	// The ID of the Direct Connect gateway.
	DxGatewayId string `json:"dxGatewayId,omitempty" yaml:"dxGatewayId,omitempty"`

	/*
	   The ID of the Direct Connect gateway association proposal.
	   Used for cross-account Direct Connect gateway associations.
	*/
	ProposalId string `json:"proposalId,omitempty" yaml:"proposalId,omitempty"`

	//
	VpnGatewayId string `json:"vpnGatewayId,omitempty" yaml:"vpnGatewayId,omitempty"`
}
