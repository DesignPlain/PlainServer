package compute

type InstanceGroupMembership struct {
	// An instance being added to the InstanceGroup
	Instance string `json:"instance,omitempty" yaml:"instance,omitempty"`

	/*
	   Represents an Instance Group resource name that the instance belongs to.


	   - - -
	*/
	InstanceGroup string `json:"instanceGroup,omitempty" yaml:"instanceGroup,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	// A reference to the zone where the instance group resides.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`
}
