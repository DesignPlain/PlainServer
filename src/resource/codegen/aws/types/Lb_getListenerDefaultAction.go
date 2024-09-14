package types

type Lb_getListenerDefaultAction struct {
	//
	Order int `json:"order,omitempty" yaml:"order,omitempty"`

	//
	Redirects []Lb_getListenerDefaultActionRedirect `json:"redirects,omitempty" yaml:"redirects,omitempty"`

	//
	TargetGroupArn string `json:"targetGroupArn,omitempty" yaml:"targetGroupArn,omitempty"`

	//
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	//
	AuthenticateCognitos []Lb_getListenerDefaultActionAuthenticateCognito `json:"authenticateCognitos,omitempty" yaml:"authenticateCognitos,omitempty"`

	//
	AuthenticateOidcs []Lb_getListenerDefaultActionAuthenticateOidc `json:"authenticateOidcs,omitempty" yaml:"authenticateOidcs,omitempty"`

	//
	FixedResponses []Lb_getListenerDefaultActionFixedResponse `json:"fixedResponses,omitempty" yaml:"fixedResponses,omitempty"`

	//
	Forwards []Lb_getListenerDefaultActionForward `json:"forwards,omitempty" yaml:"forwards,omitempty"`
}
