package types

type Accesscontextmanager_ServicePerimetersServicePerimeterStatusEgressPolicy struct {
	/*
	   Defines conditions on the source of a request causing this `EgressPolicy` to apply.
	   Structure is documented below.
	*/
	EgressFrom Accesscontextmanager_ServicePerimetersServicePerimeterStatusEgressPolicyEgressFrom `json:"egressFrom,omitempty" yaml:"egressFrom,omitempty"`

	/*
	   Defines the conditions on the `ApiOperation` and destination resources that
	   cause this `EgressPolicy` to apply.
	   Structure is documented below.
	*/
	EgressTo Accesscontextmanager_ServicePerimetersServicePerimeterStatusEgressPolicyEgressTo `json:"egressTo,omitempty" yaml:"egressTo,omitempty"`
}
