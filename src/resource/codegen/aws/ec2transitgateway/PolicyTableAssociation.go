package ec2transitgateway

type PolicyTableAssociation struct {
	// Identifier of EC2 Transit Gateway Attachment.
	TransitGatewayAttachmentId string `json:"transitGatewayAttachmentId,omitempty" yaml:"transitGatewayAttachmentId,omitempty"`

	// Identifier of EC2 Transit Gateway Policy Table.
	TransitGatewayPolicyTableId string `json:"transitGatewayPolicyTableId,omitempty" yaml:"transitGatewayPolicyTableId,omitempty"`
}
