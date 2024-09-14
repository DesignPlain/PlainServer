package compute

type SharedVPCServiceProject struct {
	// The deletion policy for the shared VPC service. Setting ABANDON allows the resource to be abandoned rather than deleted. Possible values are: "ABANDON".
	DeletionPolicy string `json:"deletionPolicy,omitempty" yaml:"deletionPolicy,omitempty"`

	// The ID of a host project to associate.
	HostProject string `json:"hostProject,omitempty" yaml:"hostProject,omitempty"`

	// The ID of the project that will serve as a Shared VPC service project.
	ServiceProject string `json:"serviceProject,omitempty" yaml:"serviceProject,omitempty"`
}
