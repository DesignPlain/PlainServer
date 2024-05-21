package types

type Container_getClusterFleet struct {
	// Full resource name of the registered fleet membership of the cluster.
	Membership string `json:"membership,omitempty" yaml:"membership,omitempty"`

	// Short name of the fleet membership, for example "member-1".
	MembershipId string `json:"membershipId,omitempty" yaml:"membershipId,omitempty"`

	// Location of the fleet membership, for example "us-central1".
	MembershipLocation string `json:"membershipLocation,omitempty" yaml:"membershipLocation,omitempty"`

	// Whether the cluster has been registered via the fleet API.
	PreRegistered bool `json:"preRegistered,omitempty" yaml:"preRegistered,omitempty"`

	/*
	   The project in which the resource belongs. If it
	   is not provided, the provider project is used.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
