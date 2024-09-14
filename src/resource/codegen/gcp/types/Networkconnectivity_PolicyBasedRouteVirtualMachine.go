package types

type Networkconnectivity_PolicyBasedRouteVirtualMachine struct {
	// A list of VM instance tags that this policy-based route applies to. VM instances that have ANY of tags specified here will install this PBR.
	Tags []string `json:"tags,omitempty" yaml:"tags,omitempty"`
}
