package types

type Accesscontextmanager_ServicePerimetersServicePerimeterSpecIngressPolicy struct {
	/*
	   Defines the conditions on the source of a request causing this `IngressPolicy`
	   to apply.
	   Structure is documented below.
	*/
	IngressFrom Accesscontextmanager_ServicePerimetersServicePerimeterSpecIngressPolicyIngressFrom `json:"ingressFrom,omitempty" yaml:"ingressFrom,omitempty"`

	/*
	   Defines the conditions on the `ApiOperation` and request destination that cause
	   this `IngressPolicy` to apply.
	   Structure is documented below.
	*/
	IngressTo Accesscontextmanager_ServicePerimetersServicePerimeterSpecIngressPolicyIngressTo `json:"ingressTo,omitempty" yaml:"ingressTo,omitempty"`
}
