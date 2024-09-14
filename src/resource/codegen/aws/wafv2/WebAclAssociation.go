package wafv2

type WebAclAssociation struct {
	// The Amazon Resource Name (ARN) of the resource to associate with the web ACL. This must be an ARN of an Application Load Balancer, an Amazon API Gateway stage (REST only, HTTP is unsupported), an Amazon Cognito User Pool, an Amazon AppSync GraphQL API, an Amazon App Runner service, or an Amazon Verified Access instance.
	ResourceArn string `json:"resourceArn,omitempty" yaml:"resourceArn,omitempty"`

	// The Amazon Resource Name (ARN) of the Web ACL that you want to associate with the resource.
	WebAclArn string `json:"webAclArn,omitempty" yaml:"webAclArn,omitempty"`
}
