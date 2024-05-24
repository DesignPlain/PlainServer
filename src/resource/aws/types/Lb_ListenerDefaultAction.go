package types

type Lb_ListenerDefaultAction struct {
	// Information for creating an action that returns a custom HTTP response. Required if `type` is `fixed-response`.
	FixedResponse Lb_ListenerDefaultActionFixedResponse `json:"fixedResponse,omitempty" yaml:"fixedResponse,omitempty"`

	// Configuration block for creating an action that distributes requests among one or more target groups. Specify only if `type` is `forward`. If you specify both `forward` block and `target_group_arn` attribute, you can specify only one target group using `forward` and it must be the same target group specified in `target_group_arn`. Detailed below.
	Forward Lb_ListenerDefaultActionForward `json:"forward,omitempty" yaml:"forward,omitempty"`

	// Order for the action. This value is required for rules with multiple actions. The action with the lowest value for order is performed first. Valid values are between `1` and `50000`.
	Order int `json:"order,omitempty" yaml:"order,omitempty"`

	// Configuration block for creating a redirect action. Required if `type` is `redirect`. Detailed below.
	Redirect Lb_ListenerDefaultActionRedirect `json:"redirect,omitempty" yaml:"redirect,omitempty"`

	// ARN of the Target Group to which to route traffic. Specify only if `type` is `forward` and you want to route to a single target group. To route to one or more target groups, use a `forward` block instead.
	TargetGroupArn string `json:"targetGroupArn,omitempty" yaml:"targetGroupArn,omitempty"`

	/*
	   Type of routing action. Valid values are `forward`, `redirect`, `fixed-response`, `authenticate-cognito` and `authenticate-oidc`.

	   The following arguments are optional:
	*/
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	// Configuration block for using Amazon Cognito to authenticate users. Specify only when `type` is `authenticate-cognito`. Detailed below.
	AuthenticateCognito Lb_ListenerDefaultActionAuthenticateCognito `json:"authenticateCognito,omitempty" yaml:"authenticateCognito,omitempty"`

	// Configuration block for an identity provider that is compliant with OpenID Connect (OIDC). Specify only when `type` is `authenticate-oidc`. Detailed below.
	AuthenticateOidc Lb_ListenerDefaultActionAuthenticateOidc `json:"authenticateOidc,omitempty" yaml:"authenticateOidc,omitempty"`
}
