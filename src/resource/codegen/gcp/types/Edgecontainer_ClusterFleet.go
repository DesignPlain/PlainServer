package types

type Edgecontainer_ClusterFleet struct {
	/*
	   (Output)
	   The name of the managed Hub Membership resource associated to this cluster.
	   Membership names are formatted as
	   `projects/<project-number>/locations/global/membership/<cluster-id>`.
	*/
	Membership string `json:"membership,omitempty" yaml:"membership,omitempty"`

	/*
	   The name of the Fleet host project where this cluster will be registered.
	   Project names are formatted as
	   `projects/<project-number>`.
	*/
	Project string `json:"project,omitempty" yaml:"project,omitempty"`
}
