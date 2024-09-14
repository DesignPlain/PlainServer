package wafregional

type WebAclAssociation struct {
	// ARN of the resource to associate with. For example, an Application Load Balancer or API Gateway Stage.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`

	// The ID of the WAF Regional WebACL to create an association.
	WebAclId string `json:"webAclId,omitempty" yaml:"webAclId,omitempty"`
}
