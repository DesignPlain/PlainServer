package types

type Container_AttachedClusterFleet struct {
	/*
	   The ID of the project in which the resource belongs.
	   If it is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`

	/*
	   (Output)
	   The name of the managed Hub Membership resource associated to this
	   cluster. Membership names are formatted as
	   projects/<project-number>/locations/global/membership/<cluster-id>.
	*/
	Membership string `json:"membership,omitempty" yaml:"membership,omitempty"`
}
