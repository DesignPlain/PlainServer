package types

type Lb_ListenerRuleAction struct {
	// Information for creating an action that distributes requests among one or more target groups. Specify only if `type` is `forward`. If you specify both `forward` block and `target_group_arn` attribute, you can specify only one target group using `forward` and it must be the same target group specified in `target_group_arn`.
	Forward Lb_ListenerRuleActionForward `json:"forward,omitempty" yaml:"forward,omitempty"`

	//
	Order int `json:"order,omitempty" yaml:"order,omitempty"`

	// Information for creating a redirect action. Required if `type` is `redirect`.
	Redirect Lb_ListenerRuleActionRedirect `json:"redirect,omitempty" yaml:"redirect,omitempty"`

	// The ARN of the Target Group to which to route traffic. Specify only if `type` is `forward` and you want to route to a single target group. To route to one or more target groups, use a `forward` block instead.
	TargetGroupArn string `json:"targetGroupArn,omitempty" yaml:"targetGroupArn,omitempty"`

	// The type of routing action. Valid values are `forward`, `redirect`, `fixed-response`, `authenticate-cognito` and `authenticate-oidc`.
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Information for creating an authenticate action using Cognito. Required if `type` is `authenticate-cognito`.
	AuthenticateCognito Lb_ListenerRuleActionAuthenticateCognito `json:"authenticateCognito,omitempty" yaml:"authenticateCognito,omitempty"`

	// Information for creating an authenticate action using OIDC. Required if `type` is `authenticate-oidc`.
	AuthenticateOidc Lb_ListenerRuleActionAuthenticateOidc `json:"authenticateOidc,omitempty" yaml:"authenticateOidc,omitempty"`

	// Information for creating an action that returns a custom HTTP response. Required if `type` is `fixed-response`.
	FixedResponse Lb_ListenerRuleActionFixedResponse `json:"fixedResponse,omitempty" yaml:"fixedResponse,omitempty"`
}
