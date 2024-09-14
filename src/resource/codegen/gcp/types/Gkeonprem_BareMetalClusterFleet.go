package types

type Gkeonprem_BareMetalClusterFleet struct {
	/*
	   (Output)
	   The name of the managed Hub Membership resource associated to this cluster.
	   Membership names are formatted as
	   `projects/<project-number>/locations/<location>/memberships/<cluster-id>`.
	*/
	Membership string `json:"membership,omitempty" yaml:"membership,omitempty"`
}
