package accesscontextmanager

import types "DesignSphere_Server/src/resource/gcp/types"

type ServicePerimeterEgressPolicy struct {
	/*
	   The name of the Service Perimeter to add this resource to.


	   - - -
	*/
	Perimeter string `json:"perimeter,omitempty" yaml:"perimeter,omitempty"`

	/*
	   Defines conditions on the source of a request causing this `EgressPolicy` to apply.
	   Structure is documented below.
	*/
	EgressFrom types.Accesscontextmanager_ServicePerimeterEgressPolicyEgressFrom `json:"egressFrom,omitempty" yaml:"egressFrom,omitempty"`

	/*
	   Defines the conditions on the `ApiOperation` and destination resources that
	   cause this `EgressPolicy` to apply.
	   Structure is documented below.
	*/
	EgressTo types.Accesscontextmanager_ServicePerimeterEgressPolicyEgressTo `json:"egressTo,omitempty" yaml:"egressTo,omitempty"`
}
