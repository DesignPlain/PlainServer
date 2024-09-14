package compute

type DiskResourcePolicyAttachment struct {
	// A reference to the zone where the disk resides.
	Zone string `json:"zone,omitempty" yaml:"zone,omitempty"`

	/*
	   The name of the disk in which the resource policies are attached to.


	   - - -
	*/
	Disk string `json:"disk,omitempty" yaml:"disk,omitempty"`

	/*
	   The resource policy to be attached to the disk for scheduling snapshot
	   creation. Do not specify the self link.
	*/
	Name string `json:"name,omitempty" yaml:"name,omitempty"`

	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
