package types

type Alb_getListenerDefaultAction struct {
	//
	Forwards []Alb_getListenerDefaultActionForward `json:"forwards,omitempty" yaml:"forwards,omitempty"`

	//
	Order int `json:"order,omitempty" yaml:"order,omitempty"`

	//
	Redirects []Alb_getListenerDefaultActionRedirect `json:"redirects,omitempty" yaml:"redirects,omitempty"`

	//
	TargetGroupArn string `json:"targetGroupArn,omitempty" yaml:"targetGroupArn,omitempty"`

	//
	Type string `json:"type,omitempty" yaml:"type,omitempty"`

	//
	AuthenticateCognitos []Alb_getListenerDefaultActionAuthenticateCognito `json:"authenticateCognitos,omitempty" yaml:"authenticateCognitos,omitempty"`

	//
	AuthenticateOidcs []Alb_getListenerDefaultActionAuthenticateOidc `json:"authenticateOidcs,omitempty" yaml:"authenticateOidcs,omitempty"`

	//
	FixedResponses []Alb_getListenerDefaultActionFixedResponse `json:"fixedResponses,omitempty" yaml:"fixedResponses,omitempty"`
}
