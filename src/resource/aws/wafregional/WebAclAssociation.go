package wafregional

type WebAclAssociation struct {
	// The ID of the WAF Regional WebACL to create an association.
	WebAclId string `json:"webAclId,omitempty" yaml:"webAclId,omitempty"`

	// ARN of the resource to associate with. For example, an Application Load Balancer or API Gateway Stage.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`
}
