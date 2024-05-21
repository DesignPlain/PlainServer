package types

type Accesscontextmanager_ServicePerimetersServicePerimeterSpecEgressPolicy struct {
	/*
	   Defines conditions on the source of a request causing this `EgressPolicy` to apply.
	   Structure is documented below.
	*/
	EgressFrom Accesscontextmanager_ServicePerimetersServicePerimeterSpecEgressPolicyEgressFrom `json:"egressFrom,omitempty" yaml:"egressFrom,omitempty"`

	/*
	   Defines the conditions on the `ApiOperation` and destination resources that
	   cause this `EgressPolicy` to apply.
	   Structure is documented below.
	*/
	EgressTo Accesscontextmanager_ServicePerimetersServicePerimeterSpecEgressPolicyEgressTo `json:"egressTo,omitempty" yaml:"egressTo,omitempty"`
}
