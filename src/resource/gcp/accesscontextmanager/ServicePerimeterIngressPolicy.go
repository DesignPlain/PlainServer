package accesscontextmanager

import types "DesignSphere_Server/src/resource/gcp/types"

type ServicePerimeterIngressPolicy struct {
	/*
	   Defines the conditions on the source of a request causing this `IngressPolicy`
	   to apply.
	   Structure is documented below.
	*/
	IngressFrom types.Accesscontextmanager_ServicePerimeterIngressPolicyIngressFrom `json:"ingressFrom,omitempty" yaml:"ingressFrom,omitempty"`

	/*
	   Defines the conditions on the `ApiOperation` and request destination that cause
	   this `IngressPolicy` to apply.
	   Structure is documented below.
	*/
	IngressTo types.Accesscontextmanager_ServicePerimeterIngressPolicyIngressTo `json:"ingressTo,omitempty" yaml:"ingressTo,omitempty"`

	/*
	   The name of the Service Perimeter to add this resource to.


	   - - -
	*/
	Perimeter string `json:"perimeter,omitempty" yaml:"perimeter,omitempty"`
}
